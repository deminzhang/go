package Test

import (
	"common/net"
	"log"
	"protos"
	"slg/const"
	"slg/item"
	"strconv"

	"github.com/golang/protobuf/proto"
)

//RPC
func init() {
	Net.RegRPC(Const.Cmd_C, func(ss Net.Session, pid int32, data []byte, uid int64) {
		ps := protos.Cmd_C{}
		err := proto.Unmarshal(data, &ps)
		if err != nil {
			log.Println("Decode error: ", err, data)
			ss.Close()
			return
		}
		args := ps.GetArgs()
		log.Println("Cmd_C: ", args)
		switch args[0] {
		case "item":
			cid, err := strconv.Atoi(args[1])
			if err != nil {
				return
			}
			num, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				return
			}
			Item.Add(uid, int32(cid), num, "test")
			break
		case "res":
			cid, err := strconv.Atoi(args[1])
			if err != nil {
				return
			}
			num, err := strconv.ParseInt(args[2], 10, 64)
			if err != nil {
				return
			}
			Item.AddRes(uid, int32(cid), num, "test")
			break
		default:
			break
		}

	})
}
