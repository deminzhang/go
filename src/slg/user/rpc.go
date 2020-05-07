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

//Net版RPC
func init() {
	Net.RegRpcC(Const.Login_C, func(ss *Net.Conn, pid int, data []byte, uid int64) {
		// Net.RegRpcC(Const.Login_C, func(ss *Net.Conn, pid int, data []byte) {
		ps := protos.Login_C{}
		if !ss.Decode(data, &ps) {
			return
		}
		// if err := Net.Decode(data, &ps); err != nil {
		// 	return
		// }
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
			log.Println("Event.OnUserNew..", user.Uid)

			//==OnUserNew(uid)
			Event.Call(Const.OnUserNew, user.Uid)

		}

		log.Println("Event.OnUserLogin..", user.Uid)
		Event.Call(Const.OnUserLogin, user.Uid)

		ss.CallOut(Const.Login_S, &protos.Login_S{})
	})

	Net.RegRpcC(Const.GetRoleInfo_C, func(ss *Net.Conn, pid int, data []byte, uid int64) {
		ps := protos.GetRoleInfo_C{}
		if !ss.Decode(data, &ps) {
			return
		}
		fmt.Println("<<<GetRoleInfo_C")
		updates := &protos.Updates{}
		Event.Call(Const.OnUserGetData, uid, updates)
		ss.CallOut(Const.Response_S, &protos.Response_S{
			Updates: updates,
		})
		ss.CallOut(Const.GetRoleInfo_S, &protos.GetRoleInfo_S{
			Uid: proto.Int64(uid),
		})
	})
	Net.RegRpcC(Const.Rename_C, func(ss *Net.Conn, pid int, data []byte, uid int64) {
		ps := protos.Rename_C{}
		if !ss.Decode(data, &ps) {
			return
		}
		fmt.Println("<<<Rename", data, ps.GetName())
		x := Sql.ORM()
		var user Entity.User
		has, _ := x.Where("Uid = ?", uid).Get(&user)
		if has {
			user.Name = ps.GetName()
			x.Update(&user)
			updates := &protos.Updates{}
			user.AppendTo(updates)
			ss.CallOut(Const.Response_S, &protos.Response_S{
				Updates: updates,
			})
		}
	})
	Net.RegRpcC(Const.UserView_C, func(ss *Net.Conn, pid int, data []byte, uid int64) {
		ps := protos.UserView_C{}
		if !ss.Decode(data, &ps) {
			return
		}
		x := Sql.ORM()
		var user Entity.User
		has, _ := x.Where("Uid = ?", uid).Get(&user)
		if has {
			updates := &protos.Updates{}
			user.AppendTo(updates)
			ss.CallOut(Const.Response_S, &protos.Response_S{
				Updates: updates,
			})
		}
	})

}
