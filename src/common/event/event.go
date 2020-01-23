package Event

import (
	"fmt"
	"log"
	"reflect"
)

//静态事件,限init时Reg 以常量为事件名

//事件响应组容器
var _event = make(map[int][]interface{})

//事件参数类型 以首次注册时为准
var _types = make(map[int]reflect.Type)

func init() {
	//Test测试示例
	fmt.Println(">>event0.sample")
	//定义
	_EventTest := -1
	_EventTest2 := -2
	//单响应
	Reg(_EventTest, func() {
		fmt.Println("_EventTest")
	})
	Call(_EventTest)

	//多响应
	Reg(_EventTest2, func(s string, i int) {
		fmt.Println(s, i)
	})
	Reg(_EventTest2, func(s string, i int) {
		fmt.Println(s, i+1)
	})
	Call(_EventTest2, "_EventTest2", 123)
}

//注册响应(事件名,回调) 以首次注册时函数类型为准
func Reg(name int, foo interface{}) {
	t0 := _types[name]
	t := reflect.TypeOf(foo)
	s := fmt.Sprintf("%s", t)
	if len(s) < 5 || s[:5] != "func(" {
		log.Fatalf("Event.Reg %s #2 must be a func*, got %s", name, s)
	}
	if t0 == nil {
		_types[name] = t
	} else {
		if t0 != t {
			log.Fatalln("Event.Reg type not equal:", name, t0, t)
		}
	}
	list := _event[name]
	list = append(list, foo)
	_event[name] = list
}

//触发事件(事件名,参数集)
func Call(name int, a ...interface{}) {
	list := _event[name]
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

//触发事件_异步(事件名,参数集)
func GoCall(name int, a ...interface{}) {
	list := _event[name]
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
		go f.Call(in)
	}
}
