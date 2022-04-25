package event_test

import (
	Event "common/event"
	"fmt"
	"testing"
)

func TestEventExample(t *testing.T) {
	//Test测试示例
	fmt.Println(">>event0.sample")
	//定义
	const (
		_EventTest  = -1
		_EventTest2 = -2
	)
	//单响应
	Event.Reg(_EventTest, func() {
		fmt.Println("_EventTest1")
	})
	Event.Call(_EventTest)

	//多响应
	Event.Reg(_EventTest2, func(s string, i int) {
		fmt.Println(s, i)
	})
	//
	Event.When(nil, 100)
	Event.Reg(_EventTest2, func(s string, i int) {
		fmt.Println(s, i+1)
	})
	Event.Call(_EventTest2, "_EventTest2A", 100)
	Event.Call(_EventTest2, "_EventTest2B", 200)

}
