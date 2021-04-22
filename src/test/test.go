package main

import (
	"net"
	// "common/net"
	// "common/sql"
	"common/util"
	"crypto/md5"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"protos"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/Luxurioust/excelize"
	"github.com/golang/protobuf/proto"
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
	fmt.Println("Util.UnixNano", time.Now().UnixNano()/1e6)
	fmt.Println("Util.Unix", time.Now().Unix())

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

	var conf GameConf
	Util.ReadToml("test.toml", &conf)
	fmt.Println("gameConfTest:", conf.Listen)

	xlsx, err := excelize.OpenFile("./test.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Get value from cell by given sheet index and axis.

	for sheetIdx := 1; ; sheetIdx++ {
		sheetName := xlsx.GetSheetName(sheetIdx)
		if sheetName == "" {
			break
		}
		if sheetName[0:1] != "#" {
			continue
		}
		fmt.Println(sheetName + ":")
		rows, _ := xlsx.GetRows(sheetName)
		colNum := 65535
		endLine := false
		for line, row := range rows {
			for col, cell := range row {
				if colNum < col {
					break
				}
				if col == 1 {
					if cell == "" || cell == "///END" {
						endLine = true
						break
					}
				}
				if line == 0 { //desc
					if cell == "" || cell == "///END" {
						if colNum > col-1 {
							colNum = col - 1
							break
						}
					} else {
						fmt.Print(cell, col, "\t")
					}
				} else if line == 2 { //key
					fmt.Print(cell, "\t")
				} else if line == 3 { //type
					fmt.Print(cell, "\t")
				} else if line == 4 { //cs
					fmt.Print(cell, "\t")
				} else {
					fmt.Print(cell, "\t")
				}
			}
			if endLine {
				break
			}
			fmt.Println()
		}
		fmt.Print(colNum, len(rows), "\n")
	}
	f := func(a int, b ...int64) {
		log.Println(a, len(b))
	}
	f(1)
	f(1, 2)
	f(1, 2, 3)

	/////////////////////////////
	var rpcF = make(map[int]interface{})
	var rpcD = make(map[int]proto.Message)
	dd := protos.Item{}
	rpcD[1] = &dd

	ii := protos.Item{
		Sid: 456,
	}

	buf, err := proto.Marshal(&ii)
	if err != nil {
		log.Fatal("Marshal error: ", err)
	}

	ff := func(pb *protos.Item) {
		log.Println(pb.GetSid())
	}
	rpcF[1] = ff
	fff := func(pb proto.Message) {
		ps := pb.(*protos.Item)
		log.Println(ps.GetSid())
	}
	ffff := func(pb interface{}) {
		ps := pb.(*protos.Item)
		log.Println(ps.GetSid())
	}

	ddd := dd //getdd()
	if err = proto.Unmarshal(buf, &ddd); err != nil {
		log.Fatal("Unmarshal error:", err)
	}
	// var dx interface{}
	// dx = &ddd
	// dt := reflect.ValueOf(dx).Elem().Type()
	// ff(dx.(dt))
	fff(&ddd)
	ffff(&ddd)
	log.Println(dd.GetSid())

	事务()

	addr := &net.UnixAddr{Name: "test1.sock", Net: "unix"}
	os.Remove(addr.Name)
	// lis, err := net.ListenUnix("unix", addr)
	// if err != nil {
	// 	fmt.Println("ListenUnix", err)
	// 	return
	// }
	// http.HandleFunc("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	// 	w.WriteHeader(222) // 这里返回一个特殊的 code, 好做验证
	// }))
	// svr := &http.Server{Handler: http.DefaultServeMux}
	// err = svr.Serve(lis)
	// if err != nil {
	// 	fmt.Println("Serve err:", err)
	// }

	// ddd = DDDD{A: 1}
	// fmt.Println("ddd :", ddd)

	buf, err = proto.Marshal(&protos.Item{Sid: 256, Cid: 65535, Num: 8})
	log.Println(buf)
	buf, err = proto.Marshal(&protos.Item{Sid: 2, Cid: 3})
	log.Println(buf)
	m := make(map[int64]int32)
	m[3] = 7
	m[6] = 5
	buf, err = proto.Marshal(&protos.TestType{
		// Fint32:       130,
		// Fint64:       -8,
		// Fuint32:      123,
		// Fuint64:      123456,
		// Fsint32:      -123,
		// Fsint64:      -123,
		// Ffixed32:     123,
		// Ffixed64:     123,
		// Fdouble:      -123.123,
		// Ffloat:       123.123,
		// Fbool:        true,
		// Fsfixed32:    1234,
		// Fsfixed64:    1234,
		Fmap: m,
		// Fint32:       130,
		// Fint64:       -8,
		// Fuint32:      123,
		// Fuint64:      123456,
		// Fsint32:      -123,
		// Fsint64:      -123,
		// Ffixed32:     123,
		// Ffixed64:     123,
		// Fdouble:      -123.123,
		// Ffloat:       123.123,
		// Fbool:        true,
		// Fsfixed32:    1234,
		// Fsfixed64:    1234,
		// Fstring:      "abcde",
		// Fbytes:       []byte{'a', 'b', 'c'},
		// Frepeatbool:  []bool{true, false, true},
		// Frepeatbool2: []bool{true, false, true},
		// Frepeatint: []int32{255, 65536},
		// Frepeatint2:  []int32{255, 65536},
		// Fstring2: []string{"abc", "abcd"},
		// Fenum:        protos.TestEnum_SUNDAY,
		// Fmessage: &protos.TestChild{
		// 	Fsint64: 123,
		// },
		// Frepc: []*protos.TestChild{&protos.TestChild{
		// 	Fsint64: 123,
		// }, &protos.TestChild{
		// 	Fsint64: 234,
		// }},
	})

	log.Println("TestType:", buf)

}

type DDDD struct {
	A int
}

func 事务() {
	defer func() {
		err := recover()
		if err == nil {
			//commit()
		} else { //产生了panic异常
			fmt.Println(err)
			//rollback()
		}
	}()
	//begin()
	panic("xxx")
}
