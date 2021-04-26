package Item

import (
	"common/net"
	"common/sql"
	"protos"
	"slg/const"
	"slg/entity"
)

//TODO 优化
func resUp(uid int64, item *Entity.Res) {
	updates := &protos.Updates{}
	item.AppendTo(updates)
	Net.CallUid(uid, Const.Response_S, &protos.Response_S{
		ProtoId: 0,
		Updates: updates,
	})
}

//单加
func AddRes(uid int64, cid int32, num int64, src string) {
	if num <= 0 {
		return
	}
	var item Entity.Res
	has, _ := Sql.ORM().Where("uid = ? and cid = ?", uid, cid).Get(&item)
	if has {
		item.Num += num
		Sql.ORM().Update(item)
	} else {
		item = Entity.Res{Uid: uid, Cid: cid, Num: num}
		Sql.ORM().Insert(item)
	}
	resUp(uid, &item)

}

//单消
func DelRes(uid int64, cid int32, num int64, src string) {
	if num <= 0 {
		return
	}
	var item Entity.Res
	has, _ := Sql.ORM().Where("uid = ? and cid = ?", uid, cid).Get(&item)
	if has && item.Num >= num {
		item.Num -= num
		Sql.ORM().Update(item)
	} else {
		Net.CallError(uid, 0, 1, "lessItem")
		return
	}
	resUp(uid, &item)
}

//混加
func AddRess(uid int64, list []*protos.IdNum, src string) {
	for _, it := range list {
		AddRes(uid, it.GetCid(), it.GetNum(), src)
	}
}

//混消
func DelRess(uid int64, list []*protos.IdNum, src string) {
	for _, it := range list {
		DelRes(uid, it.GetCid(), it.GetNum(), src)
	}
}
