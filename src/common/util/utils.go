package util

import (
	"bytes"
	"encoding/binary"
	"errors"
	"golang.org/x/exp/constraints"
	"math/rand"
	"net"
	"runtime"
	"strconv"
	"strings"
	"unsafe"
)

func GetIntranetIp() string {
	ifaces, _ := net.Interfaces()
	for _, itf := range ifaces {
		if strings.Index(itf.Name, "lo") == 0 { // loopback
			continue
		}
		addrs, _ := itf.Addrs()

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip.IsLinkLocalUnicast() || ip.IsUnspecified() || ip.IsLoopback() || ip.To4() == nil {
				continue
			}
			return ip.String()
		}
	}
	return ""
}

func IpV4ToInt(ip string) uint32 {
	var ret uint32
	binary.Read(bytes.NewBuffer(net.ParseIP(ip).To4()), binary.BigEndian, &ret)
	return ret
}

func NormalizeHost(host string) string {
	pos := strings.Index(host, ":")
	if pos >= 0 {
		host = host[:pos]
	}
	return strings.ToLower(host)
}

func GetDomainFromHost(host string) string {
	pos := strings.Index(host, ":")
	if pos >= 0 {
		host = host[:pos]
	}
	parts := strings.Split(host, ".")
	sz := len(parts)
	if sz <= 2 {
		return host
	}

	partsLen := 2
	part2 := parts[sz-2]
	switch part2 {
	case "com", "org", "net", "edu", "gov":
		partsLen = 3
	default:
		part1 := parts[sz-1]
		if part1 == "uk" || part1 == "jp" {
			switch part2 {
			case "co", "ac", "me":
				partsLen = 3
			}
		}
	}
	return strings.Join(parts[sz-partsLen:], ".")
}

func RandBetween(min, max int) (int, error) {
	if min < 0 || max < 0 || max == min {
		return 0, errors.New("rand param error")
	}

	if min > max {
		min, max = max, min
	}
	return rand.Intn(max-min) + min, nil
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

// string转ytes
func UnsafeStr2Bytes(s string) []byte {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}

// []byte转string
func UnsafeBytes2Str(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

func Slice2Map[T constraints.Ordered](list []T) map[T]struct{} {
	m := make(map[T]struct{}, len(list))
	for _, id := range list {
		m[id] = struct{}{}
	}
	return m
}

func Int2String[T constraints.Integer](t T) string {
	return strconv.Itoa(int(t))
}

func StringBuilderMisc(fn func() (string, bool), seq string) string {
	builder := strings.Builder{}
	i := 0
	for {
		s, c := fn()
		if i != 0 {
			builder.WriteString(seq)
		}
		builder.WriteString(s)
		i++
		if !c {
			break
		}
	}
	return builder.String()
}

func StringBuilderSliceJoin[T any](slice []T, fn func(t T) string, sep string) string {
	if len(slice) == 0 {
		return ""
	}

	builder := strings.Builder{}
	s1 := fn(slice[0])
	builder.Grow(len(s1) * len(slice))
	builder.WriteString(s1)

	for _, t := range slice[1:] {
		builder.WriteString(sep)
		builder.WriteString(fn(t))
	}
	return builder.String()
}

func StringBuilderMapJoin[K constraints.Ordered, V any](m map[K]V, fn func(k K, v V) string, sep string) string {
	if len(m) == 0 {
		return ""
	}

	builder := strings.Builder{}
	i := 0
	for k, v := range m {
		if i != 0 && len(sep) != 0 {
			builder.WriteString(sep)
		}

		s := fn(k, v)
		if i == 0 {
			builder.Grow(len(s) * len(m) * 3 / 4)
		}
		builder.WriteString(s)
		i++
	}
	return builder.String()
}
