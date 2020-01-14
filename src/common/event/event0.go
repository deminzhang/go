package Event

import (
	"fmt"
	"log"
	"reflect"
)

//静态事件,限init时Reg 以常量为事件名

//事件响应组容器
var events2 = make(map[int][]interface{})

//事件参数类型 以首次注册时为准
var types2 = make(map[int]reflect.Type)

func init() {
	//Test测试示例
	fmt.Println(">>event0.sample")
	//定义
	_EventTest := -1
	_EventTest2 := -2
	//单响应
	Reg0(_EventTest, func() {
		fmt.Println("_EventTest")
	})
	Call0(_EventTest)

	//多响应
	Reg0(_EventTest2, func(s string, i int) {
		fmt.Println(s, i)
	})
	Reg0(_EventTest2, func(s string, i int) {
		fmt.Println(s, i+1)
	})
	Call0(_EventTest2, "_EventTest2", 123)
}

//注册响应(事件名,回调) 以首次注册时函数类型为准
func Reg0(name int, foo interface{}) {
	t0 := types2[name]
	t := reflect.TypeOf(foo)
	s := fmt.Sprintf("%s", t)
	if len(s) < 5 || s[:5] != "func(" {
		log.Fatalf("Event.Reg %s #2 must be a func*, got %s", name, s)
	}
	if t0 == nil {
		types2[name] = t
	} else {
		if t0 != t {
			log.Fatalln("Event.Reg type not equal:", name, t0, t)
		}
	}
	list := events2[name]
	list = append(list, foo)
	events2[name] = list
}

//触发事件(事件名,参数集)
func Call0(name int, a ...interface{}) {
	list := events2[name]
	for _, foo := range list {
		f := reflect.ValueOf(foo)
		if len(a) != f.Type().NumIn() {
			log.Fatalln("Event.Call params number error:", name)
			return
		}
		in := make([]reflect.Value, len(a))
		for k, v := range a {
			in[k] = reflect.ValueOf(v)
		}
		f.Call(in)
	}
}
