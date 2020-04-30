package Entity_test

import (
	"protos"
)

//https://www.kancloud.cn/kancloud/xorm-manual-zh-cn/56004

//单主键实例
type _SampleA struct {
	Sid       int64  `xorm:"pk autoincr"` //自增主键
	Name      string `xorm:"unique"`      //唯一,禁重名
	Uid       int64  `xorm:"index"`       //拥用者分组索引
	Cid       int32  `xorm:"index"`       //配置ID分组索引
	Type      int32  `xorm:"index(type)"` //分组索引
	NewTime   int64  `xorm:"created"`     //Insert时间
	Time      int64  `xorm:"updated"`     //Insert,Update时间
	Version   int32  `xorm:"version"`     //乐观锁
	Deleted   bool   `xorm:"deleted"`     //删除标志,留库但查不到
	Transient bool   `xorm:"-"`           //不会存
	ForWrite  bool   `xorm:"->"`          //只写不读库
	ForRead   bool   `xorm:"<-"`          //只读不读库
}

//复合主键(不适合联动非关系型数据库如redis)
type _SampleB struct {
	X int32 `xorm:"pk"` //实例复合主键1
	Y int32 `xorm:"pk"` //实例复合主键2
}

type IEntity interface {
	ToProto() IEntity
	ToProtoPK() IEntity
	AppendTo(updates *protos.Updates)
	AppendToPK(removes *protos.Removes)
}
