package Net

import (
	"database/sql"
	"errors"
	"log"
	"net"
	"protos"

	"github.com/golang/protobuf/proto"
)

//---------------------------------------------------------
//基础SessionClient
type SessionC struct {
	net.Conn
	uid     int64 //userID
	protoId int32 //protoId
	errCode int32 //
	err     error //
}

//外调
func (ss *SessionC) Close() {
	ss.Conn.Close()
}

//外调
func (ss *SessionC) Send(pid int32, data []byte) {
	Send(ss.Conn, pid, data)
}

//外调
func (ss *SessionC) CallOut(pid int32, msg proto.Message) {
	conn := ss.Conn
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	Send(conn, int32(pid), data)
}

//内调
func (ss *SessionC) CallIn(pid int32, data []byte) {
	//fmt.Println(">>Session.CallIn", pid, data)
	ss.protoId = pid
	uid := ss.GetUid()
	if uid == 0 && pid%2 == 1 && pid > 101 { //未登陆 服务端 业务协议
		ss.PostError(pid, 1, "needLogin")
		return
	}
	defer ss.afterCall()
	rpc := rpcs[pid]
	if rpc == nil {
		log.Println(">>CallIn.Default", pid, uid)
		rpc = rpcs[Response_S]
		if rpc == nil {
			log.Fatal(">>NoRegRpc", Response_S)
		}
	}
	var sss Session
	sss = ss
	rpc(sss, pid, uid, data)
}

//外发错误
func (ss *SessionC) PostError(pid int32, code int32, msg string) {
	ss.errCode = code
	ss.err = errors.New(msg)
	ss.afterCall()
}

//非事务操作
func (ss *SessionC) Query(str string, args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}

//非事务操作
func (ss *SessionC) Exec(str string, args ...interface{}) (int64, int64, error) {
	return 0, 0, nil
}

//事务(禁混非事务操作)
func (ss *SessionC) GetTx() *sql.Tx {
	return nil
}

//在线设置
func (ss *SessionC) SetUid(uid int64) {
	if uid == 0 {
		G_uid2session.Set(ss.GetUid(), nil)
	} else {
		G_uid2session.Set(uid, ss)
	}
	ss.uid = uid

}
func (ss *SessionC) GetUid() int64 {
	return ss.uid
}
func (ss *SessionC) GetProtoId() int32 {
	return ss.protoId
}

func (ss *SessionC) onError() {
	log.Println(ss.err)
	if ss.errCode != 0 {
		ss.CallOut(Error_S, &protos.Error_S{
			ProtoId: proto.Int32(ss.GetProtoId()),
			Code:    proto.Int32(ss.errCode),
			Msg:     proto.String(ss.err.Error()),
		})
	}
}

func (ss *SessionC) commit() {

}

func (ss *SessionC) afterCall() {
	ss.errCode = 0
	ss.err = nil
	ss.protoId = 0
}

//错误自动断开
func (ss *SessionC) Error(err error) bool {
	if err == nil {
		return false
	}
	ss.err = err
	ss.onError()
	ss.Close()
	return true
}

//错误自动断开
func (ss *SessionC) Assert(b bool, err error) bool {
	if !b {
		ss.Error(err)
	}
	return b
}

//全量更新
func (ss *SessionC) Update() *protos.Updates {
	return nil
}

//删除
func (ss *SessionC) Remove() *protos.Removes {
	return nil

}

//增量更新
func (ss *SessionC) Props() *protos.Updates {
	return nil
}

//解码错误自动断开
func (ss *SessionC) DecodeFail(buf []byte, pb proto.Message) bool {
	err := proto.Unmarshal(buf, pb)
	return ss.Error(err)
}
