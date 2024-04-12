package logic

import (
	utilc "client0/util"
	"common/defs"
	"common/proto/comm"
	"common/tlog"
	"common/util"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"net"
	"net/url"
	"strings"
	"time"
)

func (this *game) startNet(server string) {
	ws, _ := this.connectServer(server)
	util.AssertTrue(ws != nil, "connect error")
	self := this.self
	this.conn = ws
	this.rpcLogin()
	this.AddMe(self)
	this.bStart = true
	this.rpcPlayerInfo()
	go this.routeWsMsg()
	//self.OutMsg(constant.NtpReq, &client.NtpReq{ClientTick: time.Now().UnixMilli()})
}

func (this *game) connectServer(host string) (*websocket.Conn, error) {
	//host = strings.Replace(host, "/pre", "", 1)
	if !strings.HasPrefix(host, "http") {
		host = "http://" + host
	}
	u1, err := url.ParseRequestURI(host)
	u := url.URL{Scheme: "ws", Host: u1.Host, Path: u1.Path + "/api/startws"}

	websocket.DefaultDialer.NetDial = func(network, addr string) (net.Conn, error) {
		tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
		if err != nil {
			return nil, err
		}
		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			return nil, err
		}
		conn.SetNoDelay(true)
		return conn, err
	}
	ws, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		tlog.Fatal(err)
		return nil, err
	}
	return ws, nil
}

func (this *game) OutMsg(op uint16, msg proto.Message) error {
	//logger.Debug("player send packet op:%d, msg:%+v", op, msg)
	this.netSeqId++

	out := make([]byte, 6)
	out[0] = byte(this.netSeqId >> 24)
	out[1] = byte(this.netSeqId >> 16)
	out[2] = byte(this.netSeqId >> 8)
	out[3] = byte(this.netSeqId)
	out[4] = byte(op >> 8)
	out[5] = byte(op)

	conn := this.conn
	if conn == nil {
		return errors.New("websocket has closed")
	}
	if msg == nil {
		return conn.WriteMessage(websocket.BinaryMessage, out)
	} else {
		bytes, err := proto.Marshal(msg)
		if err != nil {
			return err
		}
		out = append(out, bytes[:]...)
		return conn.WriteMessage(websocket.BinaryMessage, out)
	}
}

func (this *game) ping(now int64) {
	if now >= this.lastPingTick+10*1000 {
		this.OutMsg(defs.OpcodePing, nil)
		this.lastPingTick = now
	}
}

func (this *game) handleRcvPacket() {
	q := utilc.NewQueue[*ProtoMsg]()
	this.receiveQueue.ExportQueue(q)

	for {
		elem := q.Pop()
		if elem == nil {
			break
		}
		packet := elem.Value //.(*ProtoMsg)
		fn := handleServerMsgMap[packet.op]
		if fn != nil {
			fn(this.self, packet.msg)
		}
	}
}

func (this *game) getProtoMsg(msg []byte) (int16, []byte) {
	if len(msg) < 5 {
		return -1, nil
	}

	var op int16
	op = int16(msg[1])<<8 + int16(msg[2])

	var errCode int16
	errCode = int16(msg[3]<<8) + int16(msg[4])
	if errCode != 0 {
		return -errCode, append(make([]byte, 0), msg[5:]...)
	}
	var protoMsg []byte
	protoMsg = append(protoMsg, msg[5:]...)
	return op, protoMsg
}

func (this *game) receiveMsg() ([]byte, error) {
	//Text(json), Binary
	//if _, data, err = conn.ReadMessage(); err != nil {
	var data []byte
	var err error
	conn := this.conn
	if conn == nil {
		return nil, errors.New("read message on close conn error")
	}
	if _, data, err = conn.ReadMessage(); err != nil {
		//报错关闭websocket
		tlog.Info("ReadMessageErr", err)
		this.conn = nil
		if this.exitTick == 0 {
			this.exitTick = time.Now().UnixMilli() + 2*1000
		}
		//this.conn.Close()
		return nil, err
	}
	return data, nil
}

func (this *game) recvProtoMsg() (*ProtoMsg, error) {
	msg, err := this.receiveMsg()
	if err != nil {
		return nil, err
	}
	op, protoMsg := utilc.GetProtoMsg(msg)
	if op < 0 {
		err = errors.New(util.UnsafeBytes2Str(protoMsg))
		if uiChat != nil {
			UIChatLog("Err:" + err.Error())
			tlog.Info("ServerErr: %v\n", err.Error())
		}
		return nil, err
	}
	return &ProtoMsg{
		op:  uint16(op),
		msg: protoMsg,
	}, nil
}

func (this *game) routeWsMsg() {
	for {
		msg, err := this.recvProtoMsg()
		if err == nil {
			//this.in <- msg
			this.receiveQueue.Push(msg)
		} else {
			//tlog.Infof("routeWsMsg  err: %v\n", err)
			//break
			if this.conn == nil {
				return
			}
		}
	}
}

func (this *game) rpcLogin() bool {
	login := comm.UserLoginReq{
		Host:          defs.HomeSteadName,
		Authorization: this.self.author,
		AppInfo: &comm.AppInfo{
			DistinctId: "go_test_client" + this.self.author,
			DeviceId:   "go_test_client",
			Platform:   "win",
			Version:    0,
		},
	}
	err := this.OutMsg(defs.OpcodeUserLogin, &login)
	util.AssertTrue(err == nil, err)

	protoMsg, err := this.recvProtoMsg()
	util.AssertTrue(err == nil, err)

	util.AssertTrue(protoMsg.op == defs.OpcodeUserLoginResp, "return op error")
	var loginResp comm.UserLoginResp
	err = proto.Unmarshal(protoMsg.msg, &loginResp)
	util.AssertTrue(err == nil, err)

	this.self.id = loginResp.PlayerId
	return true
}

func (this *game) rpcNtp() {
	var req comm.NtpReq
	req.ClientTick = time.Now().UnixMilli()
	err := this.OutMsg(defs.OpcodeNtp, &req)
	fmt.Println(err)
}

func (this *game) rpcPlayerInfo() bool {
	err := this.OutMsg(defs.OpcodePlayerLogin, &comm.PlayerLoginReq{
		PlayerId: this.self.id,
		AppInfo: &comm.AppInfo{
			DistinctId:  "go_test_client" + this.self.author,
			DeviceId:    "go_test_client",
			Platform:    "win",
			Version:     0,
			VprsCurrent: "0",
			Country:     "cn",
			Timezone:    8,
		},
	})
	util.AssertTrue(err == nil, err)
	return true
}

func (this *game) rpcJoinBattle() bool {
	err := this.OutMsg(defs.OpcodePlayerJoinBattle, &comm.PlayerJoinBattleReq{
		//MapTabId: 1,
		//JoinId:     1234,
		ClientTick: time.Now().UnixMilli(),
	})
	util.AssertTrue(err == nil, err)
	return true
}

func (this *game) rpcRebuilding() bool {
	err := this.OutMsg(defs.OpcodePlayerRebuilding, &comm.PlayerRebuildingReq{
		ClientTick: time.Now().UnixMilli(),
	})
	util.AssertTrue(err == nil, err)
	return true
}
