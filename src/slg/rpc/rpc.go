package Rpc

//server common RPC
import (
	"common/net"
	"fmt"
	"protos"
)

func init() {
	Net.RegRPC(Response_S, func(ss Net.Session, pid int32, uid int64, data []byte) {
		ps := protos.Response_S{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		fmt.Println(">>>Response_S", ps)
	})
	Net.RegRPC(Error_S, func(ss Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.Error_S{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		fmt.Println("<<<Error_S", protoId, ps.GetCode(), ps.GetMsg())
	})

}
