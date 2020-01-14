package Net

import (
	"common/sql"
	"database/sql"
	"errors"
	"log"
	"net"
	"protocol"
	"sync"

	"github.com/golang/protobuf/proto"
)

const (
	SESSION_MAP_GROUP_NUM = 64
)

//RPC事务上下文---------------------------------------------------------
type Session struct {
	net.Conn
	Uid     int64   //userID
	ProtoId int32   //protoId
	tx      *sql.Tx //sql事务
	errCode int32   //
	err     error   //
	//ORM		//
	//Changes
	updates *protos.Updates
	removes *protos.Removes
	props   *protos.Updates
	//logic
	SightC int
	SightR int
}

type SessionMap struct {
	sync.RWMutex
	list map[int64]*Session
}

type SessionMaps struct {
	lists [SESSION_MAP_GROUP_NUM]*SessionMap
}

func (g *SessionMaps) Set(k int64, v *Session) {
	i := k % SESSION_MAP_GROUP_NUM
	m := g.lists[i]
	m.Lock()
	if v == nil {
		m.list[k] = v
	} else {
		delete(m.list, k)
	}
	m.Unlock()
}
func (g *SessionMaps) Get(k int64) *Session {
	i := k % SESSION_MAP_GROUP_NUM
	m := g.lists[i]
	m.Lock()
	defer m.Unlock()
	return m.list[k]
}

var G_uid2session *SessionMaps

func init() {
	G_uid2session = &SessionMaps{}
	for i := 0; i < SESSION_MAP_GROUP_NUM; i++ {
		G_uid2session.lists[i] = &SessionMap{list: make(map[int64]*Session)}
	}
}

//外调
func (ss *Session) CallOut(pid int32, msg proto.Message) {
	conn := ss.Conn
	data, err := proto.Marshal(msg)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	Send(conn, int32(pid), data)
}

//内调
func (ss *Session) CallIn(pid int32, data []byte) {
	//fmt.Println(">>Session.CallIn", pid, data)
	//conn := ss.Conn
	ss.ProtoId = pid
	uid := ss.Uid
	if uid == 0 && pid%2 == 1 && pid > 101 { //未登陆 服务端 业务协议
		ss.PError(pid, 1, "needLogin")
		return
	}
	defer ss.afterCall()
	rpc := rpcs[pid]
	if rpc == nil {
		log.Println(">>CallIn.Default", pid, uid)
		rpc = rpcs[Response_S]
	}
	rpc(ss, pid, uid, data)
}

//外发错误
func (ss *Session) PError(pid int32, code int32, msg string) {
	ss.errCode = code
	ss.err = errors.New(msg)
	ss.afterCall()
}

//非事务操作
func (ss *Session) Query(str string, args ...interface{}) (*sql.Rows, error) {
	if ss.tx != nil {
		log.Println("WARN: Tx is running")
	}
	return Sql.Query(str, args...)
}

//非事务操作
func (ss *Session) Exec(str string, args ...interface{}) (int64, int64, error) {
	if ss.tx != nil {
		log.Println("WARN: Tx is running")
	}
	return Sql.Exec(str, args...)
}

//事务(禁混非事务操作)
func (ss *Session) GetTx() *sql.Tx {
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
func (ss *Session) SetUid(uid int64) {
	if uid == 0 {
		G_uid2session.Set(ss.Uid, nil)
	} else {
		G_uid2session.Set(uid, ss)
	}
	ss.Uid = uid

}
func (ss *Session) GetUid() int64 {
	return ss.Uid
}

func (ss *Session) onError() {
	log.Println(ss.err)
	if ss.tx != nil {
		ss.tx.Rollback()
	}
	if ss.errCode != 0 {
		ss.CallOut(Error_S, &protos.Error_S{
			ProtoId: proto.Int32(ss.ProtoId),
			Code:    proto.Int32(ss.errCode),
			Msg:     proto.String(ss.err.Error()),
		})
	}
}

func (ss *Session) commit() {
	if ss.tx != nil {
		ss.tx.Commit()
	}
	if ss.updates != nil || ss.removes != nil || ss.props != nil {
		ss.CallOut(ss.ProtoId+1, &protos.Response_S{ProtoId: proto.Int32(ss.ProtoId),
			Updates: ss.updates,
			Removes: ss.removes,
			Props:   ss.props,
		})
	}
}

func (ss *Session) afterCall() {
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
	ss.ProtoId = 0
	ss.removes = nil
	ss.updates = nil
	ss.props = nil
}

//错误自动断开
func (ss *Session) Error(err error) bool {
	if err == nil {
		return false
	}
	ss.err = err
	ss.onError()
	ss.Close()
	return true
}

//错误自动断开
func (ss *Session) Assert(b bool, err error) bool {
	if !b {
		ss.Error(err)
	}
	return b
}

//全量更新
func (ss *Session) Update() *protos.Updates {
	if ss.updates == nil {
		ss.updates = &protos.Updates{}
	}
	return ss.updates
}

//删除
func (ss *Session) Remove() *protos.Removes {
	if ss.removes == nil {
		ss.removes = &protos.Removes{}
	}
	return ss.removes

}

//增量更新
func (ss *Session) Props() *protos.Updates {
	if ss.props == nil {
		ss.props = &protos.Updates{}
	}
	return ss.props
}

//解码错误自动断开
func (ss *Session) DecodeFail(buf []byte, pb proto.Message) bool {
	err := proto.Unmarshal(buf, pb)
	return ss.Error(err)
}
