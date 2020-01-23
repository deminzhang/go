package Server

import (
	"math/rand"
	"time"

	//"common/sql"
	"common/event"
	"protocol"

	"github.com/golang/protobuf/proto"
)

func init() {
	rand.Seed(time.Now().UnixNano())

	//rows, err := Sql.Query("select * from server_data where sid=?", 999)
	//if not
	//Sql.Exec("replace into server_data(sid,openTime,lastTick,bornArea,bornNum) values(?,?,?,?)", 999, tile.X, tile.Y, tile.Tp, tile.Val )

}

func GetServerId() int {
	return 999
}

func GetDBname() string {
	return "s999_slg"
}

//----------------------------------------------------------
//event
func init() {
	Event.RegA("OnUserInit", func(uid int64, updates *protos.Updates) {
		now := (time.Now().UnixNano() / 1e6)
		updates.Server = &protos.Server{
			Time:    proto.Int64(now),
			Appid:   proto.Int32(2148),
			Region:  proto.Int32(2148999),
			ChatUrl: proto.String("localhost"),
		}
	})
}
