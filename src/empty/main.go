package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
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
	fmt.Println(" args =", os.Args)
	wd, _ := os.Getwd()
	fmt.Println(" pwd =", wd)
}

type Vec3F struct {
	X float32
	Y float32
	Z float32
}

func (v3 *Vec3F) Set(X, Y, Z float32) {
	v3.X = X
	v3.Y = Y
	v3.Z = Z
}

// type Vec3Fxxx Vec3F

func (v3 *Vec3Fxxx) Set(X, Y, Z float32) {
	v3.X = X
	v3.Y = Y
	v3.Z = Z
}

type VV3 [3]float32

func test() {
	// a := []byte{1, 2, 3}
	// b := []byte{1, 2, 3}
	// log.Println(a == b)
	s1 := "123"
	s2 := "123"

	log.Println(s1 == s2)
	var v3 Vec3F
	v3.Set(1, 3, 4)
	log.Println(v3)
	var v33 Vec3Fxxx
	v3 = Vec3F(v33)
	v3.Set(1, 3, 42)
	log.Println(v3)
	var v333 VV3
	vb := []byte(v333)
	log.Println(vb)

}
