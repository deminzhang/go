package Net

import (
	"common/sql"
	"database/sql"
	"errors"
	"log"
	"net"
	"protos"

	"github.com/golang/protobuf/proto"
)

//---------------------------------------------------------
//基础SessionServer
type SessionS struct {
	net.Conn
	uid     int64   //userID
	protoId int32   //protoId
	tx      *sql.Tx //sql事务
	errCode int32   //
	err     error   //
	//ORM		//
	//Changes
	updates *protos.Updates
	removes *protos.Removes
	props   *protos.Updates
	//logic
	kvmap map[string]int
}

//外调
func (ss *SessionS) Close() {
	ss.Conn.Close()
}

//外调
func (ss *SessionS) Send(pid int32, data []byte) {
	Send(ss.Conn, pid, data)
}

//外调
func (ss *SessionS) CallOut(pid int32, msg proto.Message) {
	conn := ss.Conn
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	Send(conn, pid, data)
}

//内调
func (ss *SessionS) CallIn(pid int32, data []byte) {
	//fmt.Println(">>Session.CallIn", pid, data)
	ss.protoId = pid
	uid := ss.GetUid()
	if uid == 0 && pid%2 == 1 && pid > 101 { //未登陆 服务端 业务协议
		ss.SendError(pid, 1, "needLogin")
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
	rpc(sss, pid, data, uid)
}

//外发错误
func (ss *SessionS) SendError(pid int32, code int32, msg string) {
	ss.errCode = code
	ss.err = errors.New(msg)
	ss.afterCall()
}

//非事务操作
func (ss *SessionS) Query(str string, args ...interface{}) (*sql.Rows, error) {
	if ss.tx != nil {
		log.Println("WARN: Tx is running")
	}
	return Sql.Query(str, args...)
}

//非事务操作
func (ss *SessionS) Exec(str string, args ...interface{}) (int64, int64, error) {
	if ss.tx != nil {
		log.Println("WARN: Tx is running")
	}
	return Sql.Exec(str, args...)
}

//事务(禁混非事务操作)
func (ss *SessionS) GetTx() *sql.Tx {
	tx := ss.tx
	if tx == nil {
		var err error
		tx, err = Sql.Begin()
		if err != nil {
			log.Fatal("Sql error: ", err)
			return nil
		}
		ss.tx = tx
	}
	return tx
}

//在线设置
func (ss *SessionS) SetUid(uid int64) {
	if uid == 0 {
		G_uid2session.Set(ss.GetUid(), nil)
	} else {
		G_uid2session.Set(uid, ss)
	}
	ss.uid = uid

}
func (ss *SessionS) GetUid() int64 {
	return ss.uid
}
func (ss *SessionS) GetProtoId() int32 {
	return ss.protoId
}

func (ss *SessionS) onError() {
	log.Println(ss.err)
	if ss.tx != nil {
		ss.tx.Rollback()
	}
	if ss.errCode != 0 {
		ss.CallOut(Error_S, &protos.Error_S{
			ProtoId: proto.Int32(ss.GetProtoId()),
			Code:    proto.Int32(ss.errCode),
			Msg:     proto.String(ss.err.Error()),
		})
	}
}

func (ss *SessionS) commit() {
	if ss.tx != nil {
		ss.tx.Commit()
	}
	if ss.updates != nil || ss.removes != nil || ss.props != nil {
		ss.CallOut(ss.GetProtoId()+1, &protos.Response_S{ProtoId: proto.Int32(ss.GetProtoId()),
			Updates: ss.updates,
			Removes: ss.removes,
			Props:   ss.props,
		})
	}
}

func (ss *SessionS) afterCall() {
	if ss.err != nil {
		if ss.err != nil {
			ss.onError()
		} else {
			ss.commit()
		}
	}
	ss.errCode = 0
	ss.err = nil
	ss.tx = nil
	ss.protoId = 0
	ss.removes = nil
	ss.updates = nil
	ss.props = nil
}

//错误自动断开
func (ss *SessionS) Error(err error) bool {
	if err == nil {
		return true
	}
	ss.err = err
	ss.onError()
	ss.Close()
	return false
}

//错误自动断开
func (ss *SessionS) Assert(b bool, err error) bool {
	if !b {
		ss.Error(err)
	}
	return b
}

//全量更新
func (ss *SessionS) ProtoUpdate() *protos.Updates {
	if ss.updates == nil {
		ss.updates = &protos.Updates{}
	}
	return ss.updates
}

//删除
func (ss *SessionS) ProtoRemove() *protos.Removes {
	if ss.removes == nil {
		ss.removes = &protos.Removes{}
	}
	return ss.removes

}

//增量更新
func (ss *SessionS) ProtoProps() *protos.Updates {
	if ss.props == nil {
		ss.props = &protos.Updates{}
	}
	return ss.props
}

//自定义数值
func (ss *SessionS) Get(key string) int {
	return ss.kvmap[key]
}
func (ss *SessionS) Set(key string, val int) {
	ss.kvmap[key] = val
}

//解码错误自动断开
func (ss *SessionS) Decode(buf []byte, pb proto.Message) bool {
	err := proto.Unmarshal(buf, pb)
	return ss.Error(err)
}
