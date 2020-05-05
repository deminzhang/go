package Item

import (
	"common/net"
	"common/sql"
	"protos"
	"slg/rpc"

	"github.com/golang/protobuf/proto"
)

//TODO 优化
func resUp(uid int64, cid int32, newn int64) {
	a := []*protos.Res{&protos.Res{
		Cid: proto.Int32(cid),
		Num: proto.Int64(newn),
	}}
	updates := &protos.Updates{}
	updates.Res = a
	Net.CallUid(uid, Rpc.Response_S, &protos.Response_S{ProtoId: proto.Int32(0),
		Updates: updates,
	})
}

//单加
func AddRes(uid int64, cid int32, num int64, src string) {
	if num <= 0 {
		return
	}
	affected, _, _ := Sql.Exec("update u_res set num=num+? where uid=? and cid=?;", num, uid, cid)
	if affected == 0 {
		Sql.Exec("replace into u_res(uid,cid,num) values(?,?,num+?)", uid, cid, num)
	}
	ret := Sql.Query2Map1("select num from u_res where uid=? and cid=?;", uid, cid)
	newn := int64(ret["num"].(int64))
	resUp(uid, cid, newn)

}

//单消
func DelRes(uid int64, cid int32, num int64, src string) {
	if num <= 0 {
		return
	}
	affected, _, _ := Sql.Exec("update u_res set num=num-? where uid=? and cid=? and num>=?;", num, uid, cid, num)
	if affected == 0 {
		Net.CallError(uid, 0, 1, "lessRes")
		return
	}
	ret := Sql.Query2Map1("select num from u_res where uid=? and cid=?;", uid, cid)
	newn := int64(ret["num"].(int64))
	resUp(uid, cid, newn)
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
