package main

import (
	"fmt"
	"time"
	//"github.com/Luxurioust/excelize"
	// "github.com/aarzilli/golua/lua"
	//"github.com/golang/protobuf/proto"
)

func main() {

	loc, err := time.LoadLocation("Asia/Shanghai")
	if loc == nil {
		fmt.Println(err)
	}
	fmt.Println("TZ", loc.String())

	return
}
