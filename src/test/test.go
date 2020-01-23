package main

import (
	// "common/net"
	// "common/sql"
	"common/util"
	"crypto/md5"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"protos"
	"reflect"
	"strconv"
	"strings"
	"time"

	_ "github.com/golang/protobuf/proto"
)

//配置
type GameConf struct {
	Listen string
	DBHost string
}

func init() {
	fmt.Println(">>test.main.init")
}

func main() {
	Util.Info()
	test()
}

func test() {
	fmt.Println(">>test=========================={")
	defer fmt.Println("}")
	log.Print(123, "a" < "b", "xccx"+"xxx", reflect.TypeOf(123))
	//syslog.Debug(123,"xccx")

	fmt.Println(" utf? = 汉字")
	fmt.Println(" math.pi =", math.Pi)
	sign := md5.Sum([]byte(""))
	fmt.Printf(" md5('')=%x\n", sign)

	t1 := &protos.Response_S{}
	t2 := *t1
	t3 := &t2
	fmt.Println(">>>XXX", t3)

	resp, _ := http.Get("http://baidu.com/")
	fmt.Println("baidu:", resp)

	rand.Seed(time.Now().UnixNano())
	fmt.Println("RR", rand.Intn(10))
	fmt.Println("XXXXX", strconv.Itoa(12334), 10&3)

	fmt.Println("rand.Perm", rand.Perm(10))
	fmt.Println("Util.Min", Util.Min(4, 3, 5, 6))
	fmt.Println("Util.Min", Util.Min(3, 5, 6))

	// data := Sql.Query2Map("select * from version;")
	// fmt.Println(data, data[0], data[0]["ver"])

	// data = Sql.Query2Map("select count(*) as count from u_item where cid in(?);", 1)
	// fmt.Println("::", data, data[0]["count"])

	//Sql.Exec("set @@auto_increment_offset=?;", 998)

	//fmt.Println(true?1:2)
	// leastInterval([]byte("AAABBB"), 2)
	aa := []byte{'a', 'b'}
	bb := aa
	bb[1] = 'c'
	fmt.Println(aa, bb)

	ss := "dog cat cat dog"
	sa := strings.Split(ss, " ")
	fmt.Println(sa, len(make([]int, 5)))

	var o IEvent
	o = new(TestMgr)
	o.onTest(123)

	var conf GameConf
	Util.ReadToml("test.toml", &conf)
	fmt.Println("gameConfTest:", conf.Listen)
}

type IEvent interface {
	onTest(t int32)
}
type TestMgr struct {
	abc int32
}

func (test TestMgr) onTest(t int32) {
	fmt.Println("TestMgr.onTest", test.abc, t)
}
