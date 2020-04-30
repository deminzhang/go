package User

import (
	"common/event"
	"common/net"
	"common/sql"
	"common/util"
	"fmt"
	"log"
	"protos"
	"slg/const"
	"slg/entity"
	"slg/world"

	"github.com/golang/protobuf/proto"
)

//rpc
func init() {
	Net.RegRPC(Const.Login_C, func(ss Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.Login_C{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		fmt.Println("<<<Login", data, ps.GetOpenId(), ps.GetUid())

		//var rows sql.Rows
		if ps.GetUid() == 0 {
			//平台登陆params.GetOpenId()
		} else {
			//开发测试登陆
		}
		//sid := 999
		passport := ps.GetOpenId()
		now := Util.MilliSecond()

		x := Sql.ORM()

		var user Entity.User
		has, _ := x.Where("Passport = ?", passport).Get(&user)
		if has {
			ss.SetUid(user.Uid)
			user.LoginTime = now
			x.Update(&user)
		} else {
			name := "user_" + passport
			user = Entity.User{
				Passport:  passport,
				Name:      name,
				Level:     1,
				Version:   0,
				LoginTime: now,
				CityX:     -1,
				CityY:     -1,
			}
			t := World.GetEmptyTile(0, 0, 0)
			if t == nil {
				log.Println("Login.landFail")
				return
			}
			t.Lock()
			user.CityX = t.X
			user.CityY = t.Y
			x.Insert(&user)
			t.Tp = 3
			t.Level = 1
			t.Uid = user.Uid
			x.Insert(t)
			t.Unlock()
			uid = user.Uid
			log.Println("Event.OnUserNew..", uid)

			//==OnUserNew(uid)
			Event.Call(Const.OnUserNew, uid)

		}
		uid = user.Uid
		//updateLoginTime

		log.Println("Event.OnUserLogin..", uid)
		Event.Call(Const.OnUserLogin, uid)

		updates := &protos.Updates{}
		updates.User = user.ToProto()
		Event.Call(Const.OnUserGetData, uid, updates)

		ss.CallOut(protoId+1, &protos.Response_S{ProtoId: proto.Int32(protoId),
			Updates: updates,
		})
	})

	Net.RegRPC(Const.GetRoleInfo_C, func(ss Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.GetRoleInfo_C{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		fmt.Println("<<<GetRoleInfo_C")
		updates := ss.Update()
		Event.Call(Const.OnUserGetData, uid, updates)
		ss.CallOut(Const.GetRoleInfo_S, &protos.GetRoleInfo_S{
			Uid: proto.Int64(uid),
		})
	})
	Net.RegRPC(Const.Rename_C, func(ss Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.Rename_C{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		fmt.Println("<<<Rename", data, ps.GetName())
		x := Sql.ORM()
		var user Entity.User
		has, _ := x.Where("Uid = ?", uid).Get(&user)
		if has {
			user.Name = ps.GetName()
			x.Update(&user)
			user.AppendTo(ss.Update())
		}
	})
	Net.RegRPC(Const.ReIcon_C, func(ss Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.ReIcon_C{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		fmt.Println("<<<ReIcon", data, ps.GetIcon())
		x := Sql.ORM()
		var user Entity.User
		has, _ := x.Where("Uid = ?", uid).Get(&user)
		if has {
			user.Head = ps.GetIcon()
			x.Update(&user)
			user.AppendTo(ss.Update())
		}
	})
	Net.RegRPC(Const.ReIconB_C, func(ss Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.ReIconB_C{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		fmt.Println("<<<ReIconB_C", data, ps.GetIconB())
		x := Sql.ORM()
		var user Entity.User
		has, _ := x.Where("Uid = ?", uid).Get(&user)
		if has {
			user.HeadB = ps.GetIconB()
			x.Update(&user)
			user.AppendTo(ss.Update())
		}
	})
	Net.RegRPC(Const.UserView_C, func(ss Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.UserView_C{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		x := Sql.ORM()
		var user Entity.User
		has, _ := x.Where("Uid = ?", uid).Get(&user)
		if has {
			user.AppendTo(ss.Update())
		}
	})

}
