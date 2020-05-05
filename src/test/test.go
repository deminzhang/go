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
	"os"
	"protos"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/Luxurioust/excelize"
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

	var o IEvent
	o = new(TestMgr)
	o.onTest(123)
	var o1 IEvent
	o1 = new(TestMgr1)
	o1.onTest(123)

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

type TestMgr1 struct {
	abc int32
}

func (test TestMgr1) onTest(t int32) {
	fmt.Println("TestMgr.onTest", test.abc, t)
}
