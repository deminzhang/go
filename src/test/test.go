package main

import (
	// "common/net"
	// "common/sql"
	"crypto/md5"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"protocol"
	"reflect"
	_ "slg/rpc"
	"slg/world"
	"strconv"
	"strings"
	"time"

	_ "github.com/golang/protobuf/proto"
)

func init() {
	fmt.Println(">>test.test.init")
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
	fmt.Println("XXXXX", strconv.Itoa(12334))

	fmt.Println("rand.Perm", rand.Perm(10))

	// data := Sql.Query2Map("select * from version;")
	// fmt.Println(data, data[0], data[0]["ver"])

	// data = Sql.Query2Map("select count(*) as count from u_item where cid in(?);", 1)
	// fmt.Println("::", data, data[0]["count"])

	//Sql.Exec("set @@auto_increment_offset=?;", 998)

	fmt.Println(World.TILEX_NUM_X)
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
