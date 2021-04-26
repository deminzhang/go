package Event_test

import (
	"common/event"
	"fmt"
	"log"
	"reflect"
	"testing"

	"protos"

	"github.com/golang/protobuf/proto"
)

func TestEvent(t *testing.T) {
	ExampleEvent()
	ExampleEventR()
}

func ExampleEvent() {
	//Test测试示例
	fmt.Println(">>event0.sample")
	//定义
	_EventTest := -1
	_EventTest2 := -2
	//单响应
	Event.Reg(_EventTest, func() {
		fmt.Println("_EventTest1")
	})
	Event.Call(_EventTest)

	//多响应
	Event.Reg(_EventTest2, func(s string, i int) {
		fmt.Println(s, i)
	})
	Event.When(nil, 100)
	Event.Reg(_EventTest2, func(s string, i int) {
		fmt.Println(s, i+1)
	})
	Event.Call(_EventTest2, "_EventTest2A", 100)
	Event.Call(_EventTest2, "_EventTest2B", 200)
}

//TODO PRC事件 单个回调 使用反射

//RPC回调方法
var rpcs = make(map[int]interface{})

//RPC解码对象
var decoders = make(map[int]interface{})

//注册RPC(协议号,解码对象,回调)
func RegRPC(protoID int, decoder interface{}, call interface{}) {
	if rpcs[protoID] != nil {
		log.Fatalf("RegRPC duplicated %d", protoID)
	}
	rpcs[protoID] = call
	decoders[protoID] = decoder
}

//调用RPC(协议号,网源数据)
func CallRPC(protoID int, data []byte) {
	call := rpcs[protoID]
	if call == nil {
		log.Printf("Event UnReg RPC %d\n", protoID)
		return
	}
	decoder := decoders[protoID]
	// f0 := reflect.ValueOf(proto.Unmarshal)
	// in0 := []reflect.Value{
	// 	reflect.ValueOf(data),
	// 	reflect.ValueOf(xx),
	// }
	// f0.Call(in0)

	f := reflect.ValueOf(call)
	// in := make([]reflect.Value, 3)
	// in[0] = reflect.ValueOf(protoID)
	// in[1] = reflect.ValueOf(data)
	// in[2] = reflect.ValueOf(decoder)
	in := []reflect.Value{
		reflect.ValueOf(protoID),
		//reflect.ValueOf(data),
		reflect.ValueOf(decoder),
	}
	f.Call(in)
	f.Call(in)
}

func ExampleEventR() {
	type RPC interface {
		Reset()
		String() string
		ProtoMessage()
	}
	pb := &protos.Error_S{Code: proto.Int32(123)}
	buf, err := proto.Marshal(pb)
	if err != nil {
		log.Fatal("Marshal error: ", err)
	}
	RegRPC(1, protos.Error_S{}, func(pid int, dec protos.Error_S) {
		log.Println("CallRPC", pid, dec)
		dec.Code = proto.Int32(1234)
	})

	// func() {
	// 	if err := proto.Unmarshal(data, &xx); err != nil {
	// 		log.Println("Unmarshal error:", err)
	// 		return
	// 	}
	// }
	CallRPC(1, buf)
}
