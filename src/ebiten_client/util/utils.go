package util

import (
	"bytes"
	"golang.org/x/exp/constraints"
	"os"
	"runtime"
	"strconv"
	"strings"
)

// Args2Map os.Args除[0]外,以=分切成kv对,无=的v为"true"
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

func IsWithinRect[T constraints.Ordered](x, y, areaX, areaY, areaWidth, areaHeight T) bool {
	return x >= areaX &&
		x <= areaX+areaWidth &&
		y >= areaY &&
		y <= areaY+areaHeight
}

// 在go语言中，谷歌开发者不建议大家获取协程ID，主要是为了GC更好的工作，滥用协程ID会导致GC不能及时回收内存。
// Deprecated: 仅用于查BUG,勿用于业务 可在 ThisWorld.Env == defs.EnvDev 下使用
func GetGoroutineId() (gid uint64) {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(err)
	}
	return n
}
