package util

import "fmt"

func AssertTrue(cond bool, param ...interface{}) {
	if !cond {
		if len(param) > 0 {
			if len(param) > 1{
				panic(fmt.Sprintf(param[0].(string), param[1:]...))
			}else{
				panic(fmt.Sprint(param[0]))
			}
		}else{
			panic("AssertTrue")
		}
	}
}


func AssertFalse(cond bool, param ...interface{}) {
	if cond {
		if len(param) > 0 {
			if len(param) > 1{
				panic(fmt.Sprintf(param[0].(string), param[1:]...))
			}else{
				panic(fmt.Sprint(param[0]))
			}
		}else{
			panic("AssertTrue")
		}
	}
}
