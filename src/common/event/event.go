package Event

//静态事件,限init时Reg 以常量为事件名
import (
	"fmt"
	"log"
	"reflect"
)

//
type event struct {
	when []interface{} //条件TODO
	call interface{}   //回调
}

//事件响应组容器
var events = make(map[int][]event)

//事件参数类型 以首次注册时为准
var types = make(map[int]reflect.Type)

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
	When(nil, 100)
	Reg(_EventTest2, func(s string, i int) {
		fmt.Println(s, i+1)
	})
	Call(_EventTest2, "_EventTest2A", 100)
	Call(_EventTest2, "_EventTest2B", 200)
}

var newWhen []interface{}

//TODO 响应条件设置
func When(cond ...interface{}) {
	// fmt.Println("_EventTest")
	newWhen = cond
}

//注册响应(事件名,回调) 以首次注册时函数类型为准
func Reg(id int, foo interface{}) {
	t0 := types[id]
	t := reflect.TypeOf(foo)
	s := fmt.Sprintf("%s", t)
	if len(s) < 5 || s[:5] != "func(" {
		log.Fatalf("Event.Reg %s #2 must be a func*, got %s", id, s)
	}
	if t0 == nil {
		types[id] = t
	} else {
		if t0 != t {
			log.Fatalln("Event.Reg type not equal:", id, t0, t)
		}
	}
	list := events[id]
	list = append(list, event{
		when: newWhen,
		call: foo,
	})
	events[id] = list

	newWhen = nil
}

//触发事件_串行(事件名,参数集)
func Call(id int, a ...interface{}) {
	list := events[id]
	for _, e := range list {
		f := reflect.ValueOf(e.call)
		if len(a) != f.Type().NumIn() {
			log.Fatalln("Event.Call params number error:", id)
			return
		}
		in := make([]reflect.Value, len(a))
		for k, v := range a {
			in[k] = reflect.ValueOf(v)
			//checkWhen
			if k < len(e.when) && e.when[k] != nil {
				if e.when[k] != v {
					return
				}
			}
		}
		f.Call(in)
	}
}

//触发事件_并行(事件名,参数集)确定互不冲突可用
func GoCall(id int, a ...interface{}) {
	list := events[id]
	for _, e := range list {
		f := reflect.ValueOf(e.call)
		if len(a) != f.Type().NumIn() {
			log.Fatalln("Event.Call params number error:", id)
			return
		}
		in := make([]reflect.Value, len(a))
		for k, v := range a {
			in[k] = reflect.ValueOf(v)
			//checkWhen
			if k < len(e.when) && e.when[k] != nil {
				if e.when[k] != v {
					return
				}
			}
		}
		go f.Call(in)
	}
}
