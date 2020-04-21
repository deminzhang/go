package Server

/*
	"common/event"
	"common/sql"
	"log"
	"protos"
	"slg/const"
	"slg/entity"*/

var server_id int

func Init(serverId int) {
	server_id = serverId
}

func GetServerId() int {
	return server_id
}
