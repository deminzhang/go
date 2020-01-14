package Event

import (
	"fmt"
	"log"
	"protocol"
	"reflect"

	"github.com/golang/protobuf/proto"
)

//TODO PRC事件 单个回调 使用反射

//RPC回调方法
var rpcs = make(map[uint16]interface{})

//RPC解码对象
var decoders = make(map[uint16]interface{})

func init() {
	//Test测试示例
	fmt.Println(">>eventRPC.sample")

	RegRPC(101, protos.Login_C{}, func(protoID uint16, data []byte, ps protos.Login_C) {
		//fmt.Println("RPCEventTest", pid, ps.GetUid())
		proto.Unmarshal(data, &ps)
		fmt.Println("RPCEventTest", protoID, ps.GetUid())
	})

	data, _ := proto.Marshal(&protos.Login_C{OpenId: proto.String("test123"), Uid: proto.Int64(123)})
	data2, _ := proto.Marshal(&protos.Login_C{OpenId: proto.String("test123"), Uid: proto.Int64(456)})
	CallRPC(101, data)
	CallRPC(101, data2)
}

//注册RPC(协议号,解码对象,回调)
func RegRPC(protoID uint16, decoder interface{}, call interface{}) {
	if rpcs[protoID] != nil {
		log.Fatalf("RegRPC duplicated %d", protoID)
	}
	rpcs[protoID] = call
	decoders[protoID] = decoder
}

//调用RPC(协议号,网源数据)
func CallRPC(protoID uint16, data []byte) {
	call := rpcs[protoID]
	decoder := decoders[protoID]
	//d := decoder.(protos.Login_C)
	//proto.Unmarshal(data, &d)

	f := reflect.ValueOf(call)
	in := make([]reflect.Value, 3)
	in[0] = reflect.ValueOf(protoID)
	in[1] = reflect.ValueOf(data)
	in[2] = reflect.ValueOf(decoder)
	f.Call(in)
}
