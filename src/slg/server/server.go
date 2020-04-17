package Server

import (
	"common/event"
	"math/rand"
	"protos"
	"slg/const"

	// "slg/entity"
	"log"
	"time"
	// "github.com/golang/protobuf/proto"
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
	return "slg999"
}

//----------------------------------------------------------
//event
func init() {

	Event.Reg(Const.OnInitDB, func() {
		// x := Sql.ORM()
		// x.Sync2(new(Entity.ServerInfo))
	})
	Event.Reg(Const.OnUserGetData, func(uid int64, updates *protos.Updates) {
		log.Println("ServerInfo.OnUserGetData", uid)

	})
}
