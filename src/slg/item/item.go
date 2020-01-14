package Item

import (
	"common/event"
	"common/net"
	"common/sql"
	"common/util"
	"fmt"
	"log"
	"protocol"
	"slg/rpc"

	"github.com/golang/protobuf/proto"
)

//TODO 优化
func itemUp(uid int64, cid int32, newn int64, show int64) {
	a := []*protos.Item{&protos.Item{
		Cid:  proto.Int32(cid),
		Num:  proto.Int64(newn),
		Show: proto.Int64(show),
	}}
	updates := &protos.Updates{}
	updates.Item = a
	Net.CallUid(uid, Rpc.Response_S, &protos.Response_S{ProtoId: proto.Int32(0),
		Updates: updates,
	})
}

//单加
func Add(uid int64, cid int32, num int64, src string) {
	if num <= 0 {
		return
	}
	now := Util.MSec()

	affected, _, _ := Sql.Exec("update u_item set num=num+?,`show`=? where uid=? and cid=?;", num, now, uid, cid)
	if affected == 0 {
		Sql.Exec("replace into u_item(uid,cid,num,`show`) values(?,?,num+?,?)", uid, cid, num, now)
	}
	ret := Sql.Query2Map1("select num from u_item where uid=? and cid=?;", uid, cid)
	newn := ret["num"].(int64)
	itemUp(uid, cid, newn, now)
}

//单消
func Del(uid int64, cid int32, num int64, src string) {
	if num <= 0 {
		return
	}
	affected, _, _ := Sql.Exec("update u_item set num=num-? where uid=? and cid=? and num>=?;", num, uid, cid, num)
	if affected == 0 {
		///lessItem
		Net.CallUidError(uid, 0, 1, "lessItem")
		return
	}
	ret := Sql.Query2Map1("select num,`show` from u_item where uid=? and cid=?;", uid, cid)
	newn := ret["num"].(int64)
	show := ret["show"].(int64)
	itemUp(uid, cid, newn, show)
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

//使用
func Use(uid int64, cid int32, num int64) {
	Del(uid, cid, num, "use")
	Add(uid, cid, num, "use")
	AddRes(uid, cid, num, "use")
}

//event--------------------------------------
func init() {
	Event.Reg("OnUserNew", func(uid int64) {
		now := Util.MSec()
		for cid := 1; cid < 10; cid++ {
			Sql.Query("replace into u_item(uid,cid,num,`show`) values(?,?,?,?)", uid, cid, 1, now)
		}
	})
	Event.Reg("OnUserInit", func(uid int64, updates *protos.Updates) {
		rows, err := Sql.Query("select cid,num,`show` from u_item where uid=?", uid)
		if err != nil {
			log.Println("Item.OnUserInit error: ", err)
			return
		}
		a := []*protos.Item{}
		for rows.Next() {
			var cid int32
			var num, show int64
			rows.Scan(&cid, &num, &show)
			a = append(a, &protos.Item{
				Cid:  proto.Int32(cid),
				Num:  proto.Int64(num),
				Show: proto.Int64(show),
			})
		}
		updates.Item = a
	})
}

//-----------------------------------------
//RPC
func init() {

	Net.RegRPC(Rpc.ItemUse_C, func(ss *Net.Session, protoId int32, uid int64, data []byte) {
		fmt.Println(">>>ItemUse_C", data)

		ps := protos.ItemUse_C{}
		err := proto.Unmarshal(data, &ps)
		if err != nil {
			log.Println("View_C.Decode error: ", err, data)
			ss.Close()
			return
		}
		fmt.Println(">>>ItemUse_C", ps.GetCid(), ps.GetNum())

		Use(uid, ps.GetCid(), int64(ps.GetNum()))

	})
	Net.RegRPC(Rpc.ItemDel_C, func(ss *Net.Session, protoId int32, uid int64, data []byte) {
		fmt.Println(">>>ItemUse_C", data)

		ps := protos.ItemDel_C{}
		err := proto.Unmarshal(data, &ps)
		if err != nil {
			log.Println("View_C.Decode error: ", err, data)
			ss.Close()
			return
		}
		fmt.Println(">>>ItemUse_C", ps.GetCid(), ps.GetNum())

		Del(uid, ps.GetCid(), int64(ps.GetNum()), "del")

	})

	// Net.RegRPC(Rpc.ShopBuy_C, func(ss *Net.Session, protoId int32, uid int64, data []byte) {
	// 	fmt.Println(">>>ItemUse_C", data)

	// 	ps := protos.ShopBuy_C{}
	// 	err := proto.Unmarshal(data, &ps)
	// 	if err != nil {
	// 		log.Println("View_C.Decode error: ", err, data)
	// 		ss.Close()
	// 		return
	// 	}
	// 	fmt.Println(">>>ShopBuy_C", ps.GetShop(), ps.GetId(), ps.GetNum())
	// })

}
