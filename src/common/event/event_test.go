package event

import (
	"fmt"
	"testing"
)

func TestEventExample(t *testing.T) {
	//Test测试示例
	fmt.Println(">>event0.sample")
	//定义
	_EventTest := -1
	_EventTest2 := -2
	//单响应
	Reg(_EventTest, func() {
		fmt.Println("_EventTest1")
	})
	Call(_EventTest)

	//多响应
	Reg(_EventTest2, func(s string, i int) {
		fmt.Println(s, i)
	})
	When(nil, 100)
	Reg(_EventTest2, func(s string, i int) {
		fmt.Println(s, i+1)
	})
	Call(_EventTest2, "_EventTest2A", 100)
	Call(_EventTest2, "_EventTest2B", 200)

}
