package event

import (
	"fmt"
	"log"
	"reflect"
)

//Event 反射版静态事件
//因用反射,限init时Def,Reg,不作UnReg
//以常量为事件名
type Event struct {
	id        int
	tp        reflect.Type
	caller    reflect.Value //总回调
	listeners []listener
	argn      int //参数数
}
type listener struct {
	filter []interface{} //条件 按参数顺序,nil为无条件
}

//事件
var events = make(map[int]*Event)

func def(id int, foo interface{}) {
	e := events[id]
	if e != nil {
		log.Fatalln("Event.Reg Duplicate Define:", id)
	}

	t := reflect.TypeOf(foo)
	s := fmt.Sprintf("%s", t)
	if len(s) < 5 || s[:5] != "func(" {
		log.Fatalf("Event.Reg %d #2 must be a func*, got %s", id, s)
	}
	call := reflect.ValueOf(foo)
	events[id] = &Event{
		id:     id,
		tp:     t,
		caller: call,
		argn:   call.Type().NumIn(),
	}
}

var newWhen []interface{} //TODO 这个在并发情况下不安全,当然保证不动态Reg可没事

//Filter响应条件设置 仅对下一个Reg有效
func When(cond ...interface{}) {
	// fmt.Println("_EventTest")
	newWhen = cond
}

//Listener
//注册响应(事件名,回调) 以首次注册时函数类型为准 尽量init时全注册好 动态注册可能造成call时的不确定性
func Reg(id int, cb interface{}) {
	e := events[id]
	if e == nil {
		def(e.id, cb)
	}
	//	e.Reg(cb)
	//}
	//func (e *Event) Reg(cb interface{}) {
	t := reflect.TypeOf(cb)
	s := fmt.Sprintf("%s", t)
	if len(s) < 5 || s[:5] != "func(" {
		log.Fatalf("Event.Reg %d #2 must be a func*, got %s", e.id, s)
	}
	if e.tp != t {
		log.Fatalln("Event.Reg type not equal:", e.id)
	}
	e.listeners = append(e.listeners, listener{
		filter: newWhen,
	})
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
	e := events[id]
	//	e.Call(a...)
	//}
	//func (e *Event) Call(a ...interface{}) {
	list := e.listeners
	caller := e.caller
	in, argn := makeArgv(a...)
	if argn != e.argn {
		log.Fatalln("Event.Call params number error:", e.id, argn, caller.Type().NumIn())
		return
	}
	for _, e := range list {
		if unMatch(e.filter, a) {
			continue
		}
		caller.Call(in)
	}
}

//触发事件_异步并行(事件名,参数集)确定互不冲突可用
func GoCall(id int, a ...interface{}) {
	e := events[id]
	//	e.GoCall(a...)
	//}
	//func (e *Event) GoCall(a ...interface{}) {
	list := e.listeners
	caller := e.caller
	in, argn := makeArgv(a...)
	if argn != e.argn {
		log.Fatalln("Event.Call params number error:", id, argn, caller.Type().NumIn())
		return
	}
	for _, e := range list {
		if unMatch(e.filter, a) {
			continue
		}
		go caller.Call(in)
	}
}

//触发事件_同步并行(事件名,参数集)确定互不冲突可用
func GoCallWaitAll(id int, a ...interface{}) {
	e := events[id]
	//	e.GoCallWaitAll(a...)
	//}
	//func (e *Event) GoCallWaitAll(a ...interface{}) {
	list := e.listeners
	caller := e.caller
	num := len(list)
	if num == 0 {
		return
	}
	in, argn := makeArgv(a...)
	if argn != e.argn {
		log.Fatalln("Event.Call params number error:", id, argn, caller.Type().NumIn())
		return
	}
	for _, e := range list {
		if unMatch(e.filter, a) {
			num--
			continue
		}
		go func() {
			defer func() {
				num--
			}()
			caller.Call(in)
		}()
	}
	for num > 0 {
	}
}
