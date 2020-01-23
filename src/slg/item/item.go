package Item

import (
	"common/net"
	"common/sql"
	"common/util"
	"protos"
	"slg/entity"
	"slg/rpc"

	"github.com/golang/protobuf/proto"
)

//TODO 优化
func itemUp(uid int64, item *Entity.Item) {
	updates := &protos.Updates{}
	item.AppendTo(updates)
	Net.CallUid(uid, Rpc.Response_S, &protos.Response_S{ProtoId: proto.Int32(0),
		Updates: updates,
	})
}

//单加
func Add(uid int64, cid int32, num int64, src string) {
	if num <= 0 {
		return
	}
	now := Util.MilliSecond()
	var item Entity.Item
	has, _ := Sql.ORM().Where("uid = ? and cid = ?", uid, cid).Get(&item)
	if has {
		item.Num += num
		item.Time = now
		Sql.ORM().Update(item)
	} else {
		item = Entity.Item{Uid: uid, Cid: cid, Num: num, Time: now}
		Sql.ORM().Insert(item)
	}
	itemUp(uid, &item)
}

//单消
func Del(uid int64, cid int32, num int64, src string) {
	if num <= 0 {
		return
	}
	var item Entity.Item
	has, _ := Sql.ORM().Where("uid = ? and cid = ?", uid, cid).Get(&item)
	if has && item.Num >= num {
		item.Num -= num
		Sql.ORM().Update(item)
	} else {
		Net.CallUidError(uid, 0, 1, "lessItem")
		return
	}
	itemUp(uid, &item)
}

//混加
func Adds(uid int64, list []*protos.IdNum, src string) {
	for _, it := range list {
		Add(uid, it.GetCid(), it.GetNum(), src)
	}
}

//混消
func Dels(uid int64, list []*protos.IdNum, src string) {
	for _, it := range list {
		Del(uid, it.GetCid(), it.GetNum(), src)
	}
}

//使用 TODO
func Use(uid int64, cid int32, num int64) {
	Del(uid, cid, num, "use")
	Add(uid, cid, num, "use")
	AddRes(uid, cid, num, "use")
}
