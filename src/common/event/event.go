package Event

import (
	"fmt"
	"log"
	"reflect"
)

//静态事件,限init时Reg 以字串为事件名

//事件响应组容器
var events = make(map[string][]interface{})

//事件参数类型 以首次注册时为准
var types = make(map[string]reflect.Type)

func init() {
	//Test测试示例
	fmt.Println(">>event.sample")
	//单响应
	Reg("_EventTest", func() {
		fmt.Println("_EventTest")
	})
	Call("_EventTest")

	//多响应
	Reg("_EventTest2", func(s string, i int) {
		fmt.Println(s, i)
	})
	Reg("_EventTest2", func(s string, i int) {
		fmt.Println(s, i+1)
	})
	Call("_EventTest2", "_EventTest2", 123)
}

//注册响应(事件名,回调) 以首次注册时函数类型为准
func Reg(name string, foo interface{}) {
	t0 := types[name]
	t := reflect.TypeOf(foo)
	s := fmt.Sprintf("%s", t)
	if len(s) < 5 || s[:5] != "func(" {
		log.Fatalf("Event.Reg %s #2 must be a func*, got %s", name, s)
	}
	if t0 == nil {
		types[name] = t
	} else {
		if t0 != t {
			log.Fatalln("Event.Reg type not equal:", name, t0, t)
		}
	}
	list := events[name]
	list = append(list, foo)
	events[name] = list
}

//触发事件(事件名,参数集)
func Call(name string, a ...interface{}) {
	list := events[name]
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
