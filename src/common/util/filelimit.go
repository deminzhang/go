//go:build !windows

package util

import (
	"fmt"
	"syscall"
)

func SetFileLimit() error {
	fileLimitWant := uint64(655360)
	var rLimit syscall.Rlimit

	err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		fmt.Println(err)
		return err
	}
	if rLimit.Cur >= fileLimitWant {
		return nil
	}

	if rLimit.Max < fileLimitWant {
		fmt.Printf("warning: file limit max %d, but want %d\n", rLimit.Max, fileLimitWant)
		fileLimitWant = rLimit.Max
	}

	oldLimit := rLimit.Cur
	rLimit.Cur = fileLimitWant
	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if rLimit.Cur < fileLimitWant {
		err = fmt.Errorf("warning: failed to set limit, current %d, want %d", rLimit.Cur, fileLimitWant)
		fmt.Println(err)
		return err
	}

	fmt.Printf("file limit set %d, old limit %d\n", rLimit.Cur, oldLimit)
	return nil
}
