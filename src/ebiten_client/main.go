package main

import (
	"client0/logic"
	"client0/util"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

//>./client0.exe host= user= passwd= api=

func main() {
	args := util.Args2Map()
	user := "test001@test.com"
	passwd := "123456"
	host := "localhost" //localhost

	//用参数
	if val, ok := args["host"]; ok {
		host = val
	}
	if val, ok := args["user"]; ok {
		user = val
	}
	if val, ok := args["passwd"]; ok {
		passwd = val
	}
	//------------------------------------
	game := logic.World.ShowLogin(host, user, passwd)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
