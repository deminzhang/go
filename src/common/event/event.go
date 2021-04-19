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
	// foo  interface{}   //原始方法
	call reflect.Value //回调
	argn int           //参数数
}

//事件响应组容器
var events = make(map[int][]event)

//事件参数类型 以首次注册时为准
var types = make(map[int]reflect.Type)

var newWhen []interface{} //TODO 这个在并发情况下不安全,当然保证不动态Reg可没事

//响应条件设置 仅对下一个Reg有效
func When(cond ...interface{}) {
	// fmt.Println("_EventTest")
	newWhen = cond
}

//Listener
//注册响应(事件名,回调) 以首次注册时函数类型为准 尽量init时全注册好 动态注册可能造成call时的不确定性
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
	call := reflect.ValueOf(foo)
	list := append(events[id], event{
		when: newWhen,
		// foo:  foo,
		call: call,
		argn: call.Type().NumIn(),
	})
	events[id] = list
	newWhen = nil
}

//构建反射调用参数
func makeArgv(a ...interface{}) ([]reflect.Value, int) {
	argn := len(a)
	in := make([]reflect.Value, argn)
	for k, v := range a {
		in[k] = reflect.ValueOf(v)
	}
	return in, argn
}

//筛选条件执行
func unMatch(when []interface{}, a ...interface{}) bool {
	if when != nil {
		for k, v := range a {
			if k < len(when) && when[k] != nil && when[k] != v {
				return true
			}
		}
	}
	return false
}

//Dispatcher
//触发事件_同步串行(事件名,参数集)
func Call(id int, a ...interface{}) {
	list := events[id]
	in, argn := makeArgv(a...)
	for _, e := range list {
		if argn != e.argn {
			log.Fatalln("Event.Call params number error:", id, argn, e.argn)
			return
		}
		if unMatch(e.when, a) {
			continue
		}
		e.call.Call(in)
	}
}

//触发事件_异步并行(事件名,参数集)确定互不冲突可用
func GoCall(id int, a ...interface{}) {
	list := events[id]
	in, argn := makeArgv(a...)
	for _, e := range list {
		if argn != e.argn {
			log.Fatalln("Event.GoCall params number error:", id, argn, e.argn)
			return
		}
		if unMatch(e.when, a) {
			continue
		}
		go e.call.Call(in)
	}
}

//触发事件_同步并行(事件名,参数集)确定互不冲突可用
func GoCallS(id int, a ...interface{}) {
	list := events[id]
	num := len(list)
	if num == 0 {
		return
	}
	in, argn := makeArgv(a...)
	for _, e := range list {
		if argn != e.argn {
			log.Fatalln("Event.GoCallS params number error:", id, argn, e.argn)
			return
		}
		if unMatch(e.when, a) {
			num--
			continue
		}
		go func() {
			defer func() {
				num--
			}()
			e.call.Call(in)
		}()
	}
	for num > 0 {
	}
}
