package Event

// import (
// 	"log"
// 	"reflect"
// )

// //TODO PRC事件 单个回调 使用反射

// //RPC回调方法
// var rpcs = make(map[int]interface{})

// //RPC解码对象
// var decoders = make(map[int]interface{})

// //注册RPC(协议号,解码对象,回调)
// func RegRPC(protoID int, decoder interface{}, call interface{}) {
// 	if rpcs[protoID] != nil {
// 		log.Fatalf("RegRPC duplicated %d", protoID)
// 	}
// 	rpcs[protoID] = call
// 	decoders[protoID] = decoder
// }

// //调用RPC(协议号,网源数据)
// func CallRPC(protoID int, data []byte) {
// 	call := rpcs[protoID]
// 	if call == nil {
// 		log.Printf("Event UnReg RPC %d\n", protoID)
// 		return
// 	}
// 	decoder := decoders[protoID]
// 	f := reflect.ValueOf(call)
// 	in := make([]reflect.Value, 3)
// 	in[0] = reflect.ValueOf(protoID)
// 	in[1] = reflect.ValueOf(data)
// 	in[2] = reflect.ValueOf(decoder)
// 	f.Call(in)
// }
