package Entity_test

import (
	// "protos"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/go-xorm/xorm"
)

//https://www.kancloud.cn/kancloud/xorm-manual-zh-cn/56004

//单主键实例
type TestItem struct {
	Sid       int64  `xorm:"pk autoincr"` //自增主键
	Name      string `xorm:"unique"`      //唯一,禁重名
	Uid       int64  `xorm:"index"`       //拥用者分组索引
	Cid       int32  `xorm:"index"`       //配置ID分组索引
	Type      int32  `xorm:"index(type)"` //分组索引
	NewTime   int64  `xorm:"created"`     //Insert时间
	Time      int64  `xorm:"updated"`     //Insert,Update时间
	Version   int32  `xorm:"version"`     //乐观锁
	Deleted   bool   `xorm:"deleted"`     //删除标志,留库但查不到
	Transient bool   `xorm:"-"`           //不会存库
	ForWrite  bool   `xorm:"->"`          //只写不读库
	ForRead   bool   `xorm:"<-"`          //只读不读库
	TabName   string `xorm:"-"`           //指定存表名
}

//http://www.itspire.cn/article/116.html
//实现TableName接口，以达到结构体和表对应，如果不实现该接口，并未设置全局表名禁用复数，gorm会自动扩展表名为articles（结构体+s）
func (it TestItem) TableName() string {
	return it.TabName
}

//复合主键(不适合联动非关系型数据库如redis)
type SampleMK struct {
	X int32 `xorm:"pk"` //实例复合主键1
	Y int32 `xorm:"pk"` //实例复合主键2
}

// type IEntity interface {
// 	ToProto() IEntity
// 	ToProtoPK() IEntity
// 	AppendTo(updates *protos.Updates)
// 	AppendToPK(removes *protos.Removes)
// }

var _orm *xorm.Engine

func init() {
	var err error
	// _orm, err = xorm.NewEngine("mysql", "root:1614@tcp(localhost:3306)/slg999?charset=utf8")
	_orm, err = xorm.NewEngine("postgres", "postgres://postgres:1614@localhost/slg999?sslmode=disable")
	if err != nil {
		panic(err)
	}
	//_orm.ShowSQL(true) //则会在控制台打印出生成的SQL语句；
}

func TestEvent(t *testing.T) {
	//初始化/同步/更新SQL结构
	_orm.Sync2(new(TestItem), TestItem{TabName: "testitem1"}, TestItem{TabName: "testitem2"})

	it := TestItem{Uid: 1, Cid: 1, TabName: "testitem1"}
	_orm.Insert(it)
	it = TestItem{Uid: 2, Cid: 2, TabName: "testitem2"}
	_orm.Insert(it)

	//读取
	var item TestItem
	has, _ := _orm.Where("uid = ? and cid = ?", 1, 1).Get(&item)
	if has {
		item.Time = 1
		_orm.Update(item)
	} else {
		item = TestItem{Uid: 1, Cid: 1, Time: 0}
		_orm.Insert(item)
	}
	//批量读
	items := make([]TestItem, 0)
	err := _orm.Where("Uid = ?", 1).Find(&items)
	if err != nil {
		log.Println(err)
	}
	for _, o := range items {
		item.Time = 2
		_orm.Update(o)
	}

}
