package Event

//静态事件,限init时Reg 以常量为事件名
import (
	"fmt"
	"log"
	"reflect"
)

//
type event struct {
	when []interface{} //条件 按参数顺序,nil为无条件
	call interface{}   //回调
}

//事件响应组容器
var events = make(map[int][]event)

//事件参数类型 以首次注册时为准
var types = make(map[int]reflect.Type)

var newWhen []interface{}

//响应条件设置 仅对下一个Reg有效
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
		log.Fatalf("Event.Reg %d #2 must be a func*, got %s", id, s)
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
		if e.when == nil {
			for k, v := range a {
				in[k] = reflect.ValueOf(v)
			}
		} else {
			for k, v := range a {
				in[k] = reflect.ValueOf(v)
				if k < len(e.when) && e.when[k] != nil && e.when[k] != v {
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
		if e.when == nil {
			for k, v := range a {
				in[k] = reflect.ValueOf(v)
			}
		} else {
			for k, v := range a {
				in[k] = reflect.ValueOf(v)
				if k < len(e.when) && e.when[k] != nil && e.when[k] != v {
					return
				}
			}
		}
		go f.Call(in)
	}
}
