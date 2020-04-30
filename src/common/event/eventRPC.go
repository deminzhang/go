package Event

import (
	"log"
	"reflect"
)

//TODO PRC事件 单个回调 使用反射

//RPC回调方法
var rpcs = make(map[uint16]interface{})

//RPC解码对象
var decoders = make(map[uint16]interface{})

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
