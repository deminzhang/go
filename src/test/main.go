package main

import (
	"common/util"
	"fmt"
)

func init() {
	fmt.Println(">>test.main.init")
}

func main() {
	Util.Info()
	test()
}
