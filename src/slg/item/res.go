package Item

import (
	"common/net"
	"common/sql"
	"protocol"
	"slg/rpc"

	"github.com/golang/protobuf/proto"
)

const (
	RES_FOOD  = 1  //粮食
	RES_WOOD  = 2  //木头
	RES_IRON  = 3  //铁
	RES_GOLD  = 4  //金子
	RES_VIP   = 5  //vip点
	RES_EXP   = 6  //领主经验
	RES_7     = 7  //体力
	RES_GCOIN = 8  //联盟货币
	RES_9     = 9  //竞技场荣誉值
	RES_GEMB  = 10 //赠送砖石
	RES_GEM   = 11 //充值钻石
	RES_GEMS  = 12 //通用钻石

)

func init() {

}

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
		Net.CallUidError(uid, 0, 1, "lessRes")
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
