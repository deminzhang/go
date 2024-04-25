package main

import (
	"fmt"
	"unicode/utf8"
	//"github.com/Luxurioust/excelize"
	// "github.com/aarzilli/golua/lua"
	//"github.com/golang/protobuf/proto"
)

func main() {
	ss := "你好"
	ll := utf8.RuneCountInString(ss)
	fmt.Println("TZ", ll)

	return
}
