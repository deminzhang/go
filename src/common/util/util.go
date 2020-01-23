package Util

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

//unix时间ms级
func MilliSecond() int64 {
	return time.Now().UnixNano() / 1e6
}

//运行环境,参数,INFO等
func Info() {
	fmt.Println(">>osinfo========================{")
	defer fmt.Println("}")
	fmt.Println(" version =", runtime.Version())
	fmt.Println(" osType =", runtime.GOOS)
	fmt.Println(" pid =", os.Getpid())
	fmt.Println(" ppid =", os.Getppid())
	if runtime.GOOS == "linux" {
		fmt.Println(" gid =", os.Getgid())
		fmt.Println(" egid =", os.Getegid())
		fmt.Println(" uid =", os.Getuid())
		fmt.Println(" euid =", os.Geteuid())
	}
	host, _ := os.Hostname()
	fmt.Println(" hostname =", host)

	t := time.Now()
	fmt.Println("", t)
	fmt.Println(" unixtime =", t.Unix())
	fmt.Println(" Millisecond =", MilliSecond())
	fmt.Println(" args =", os.Args)
	wd, _ := os.Getwd()
	fmt.Println(" pwd =", wd)
}
