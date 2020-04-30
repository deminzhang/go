package Event_test

import (
	"common/event"
	"fmt"
	"testing"
)

func TestEvent(t *testing.T) {
	ExampleEvent()
	// ExampleEventRPC()
}

func ExampleEvent() {
	//Test测试示例
	fmt.Println(">>event0.sample")
	//定义
	_EventTest := -1
	_EventTest2 := -2
	//单响应
	Event.Reg(_EventTest, func() {
		fmt.Println("_EventTest")
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

// func ExampleEventRPC() {
// 	//Test测试示例
// 	fmt.Println(">>eventRPC.sample")

// 	RegRPC(101, protos.Login_C{}, func(protoID uint16, data []byte, ps protos.Login_C) {
// 		//fmt.Println("RPCEventTest", pid, ps.GetUid())
// 		proto.Unmarshal(data, &ps)
// 		fmt.Println("RPCEventTest", protoID, ps.GetUid())
// 	})

// 	data, _ := proto.Marshal(&protos.Login_C{OpenId: proto.String("test123"), Uid: proto.Int64(123)})
// 	data2, _ := proto.Marshal(&protos.Login_C{OpenId: proto.String("test123"), Uid: proto.Int64(456)})
// 	CallRPC(101, data)
// 	CallRPC(101, data2)
// }
