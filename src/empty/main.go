package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"
)

func init() {
	fmt.Println(">>main.init")
	rand.Seed(time.Now().UnixNano())
	info()
}

func main() {
	//test()
	var a [3]int
	a[0] = 1
	a[1] = 21
	a[2] = 3
	fmt.Println(">>main.hello", a)
	test()
}

//Info info
func info() {
	fmt.Println(">>osinfo========================{")
	defer fmt.Println("}")
	fmt.Println(" version =", runtime.Version())
	fmt.Println(" GOOS =", runtime.GOOS)
	fmt.Println(" pid =", os.Getpid())
	fmt.Println(" ppid =", os.Getppid())
	if runtime.GOOS == "linux" {
		fmt.Println(" gid =", os.Getgid())
		fmt.Println(" egid =", os.Getegid())
		fmt.Println(" uid =", os.Getuid())
		fmt.Println(" euid =", os.Geteuid())
	}
	fmt.Println(" NumCPU =", runtime.NumCPU())
	fmt.Println(" NumCgoCall =", runtime.NumCgoCall())
	fmt.Println(" NumGoroutine =", runtime.NumGoroutine())
	host, _ := os.Hostname()
	fmt.Println(" hostname =", host)

	t := time.Now()
	fmt.Println("", t)
	fmt.Println(" unixtime =", t.Unix())
	fmt.Println(" Millisecond =", time.Now().UnixNano()/1e6)
	fmt.Println(" args =", os.Args, len(os.Args))
	wd, _ := os.Getwd()
	fmt.Println(" pwd =", wd)
	args := Args2Map()
	fmt.Println(" Args2Map =", args)
	fmt.Println(" Args2Map =", args["123XXX"] == "")
	fmt.Println(" Args2Map =", args["a"] == "")
}

func Args2Map() map[string]string {
	m := make(map[string]string)
	for i, s := range os.Args {
		if i == 0 {
			continue
		}
		ss := strings.Split(s, "=")
		if len(ss) > 1 {
			m[ss[0]] = ss[1]
		} else {
			m[ss[0]] = "true"
		}
	}
	return m
}

type Vec3F struct {
	X float32
	Y float32
	Z float32
}

func test() {
	// a := []byte{1, 2, 3}
	// b := []byte{1, 2, 3}
	// log.Println(a == b)
	s1 := "123"
	s2 := "123"

	log.Println(s1 == s2)

}
