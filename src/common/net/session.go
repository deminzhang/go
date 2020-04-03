package Net

import (
	"log"
	"protos"
	"sync"

	"github.com/golang/protobuf/proto"
)

const (
	SESSION_MAP_GROUP_NUM = 64
)

//ISession用接口方式可以按需多态,Session将简化只做为基础Session
type Session interface {
	Close()
	Send(int32, []byte)
	CallOut(int32, proto.Message)
	CallIn(int32, []byte)
	PError(int32, int32, string)
	Error(error) bool
	Assert(bool, error) bool

	GetProtoId() int32
	SetUid(int64)
	GetUid() int64
	DecodeFail([]byte, proto.Message) bool

	Update() *protos.Updates
	Remove() *protos.Removes
	Props() *protos.Updates
}

type SessionMap struct {
	sync.RWMutex
	list map[int64]Session
}

type SessionMaps struct {
	lists [SESSION_MAP_GROUP_NUM]*SessionMap
}

func (g *SessionMaps) Set(k int64, v Session) {
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
func (g *SessionMaps) Get(k int64) Session {
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
		G_uid2session.lists[i] = &SessionMap{list: make(map[int64]Session)}
	}
	//TEST
	var c Session
	//c = new(Session)
	c = &SessionS{uid: 1234}

	log.Println("TEST ISession", c.GetUid())

}
