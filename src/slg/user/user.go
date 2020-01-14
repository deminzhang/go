package User

import (
	"common/event"
	"common/net"
	"common/sql"
	"common/util"
	"fmt"
	"log"
	"protocol"
	"slg/rpc"
	"time"

	"github.com/golang/protobuf/proto"
)

func init() {
	go tickUser()
}

func tickUser() {
	//fmt.Println(">>ticking")
	for {
		time.Sleep(time.Second * 10)
		//fmt.Println(">>G_userNum=", len(G_conn2user))
		// for conn := range G_conn2user {
		// 	ss.CallOut( 12, &protos.Response_S{ProtoId: proto.Int32(0),
		// 		Updates: &protos.Updates{},
		// 	})
		// }
	}
}

//event
func init() {
	Event.Reg("OnDisconn", func(uid int64) {

	})
}

//rpc
func init() {

	Net.RegRPC(Rpc.Login_C, func(ss *Net.Session, protoId int32, uid int64, data []byte) {
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
		now := Util.MSec()

		newUser := true
		var user protos.User

	LAB_USER:
		rows, err := Sql.Query("select uid, name, gender, icon, iconB, level, cityX, cityY from u_user where passport=?", passport)
		for rows.Next() {
			newUser = false
			var uid int64
			var name []byte
			var gender, icon, iconB, level, cityX, cityY int32
			err = rows.Scan(&uid, &name, &gender, &icon, &iconB, &level, &cityX, &cityY)
			user = protos.User{
				Uid:    proto.Int64(uid),
				Name:   proto.String(string(name)),
				Gender: proto.Int32(gender),
				Icon:   proto.Int32(icon),
				IconB:  proto.Int32(iconB),
				Level:  proto.Int32(level),
				CityX:  proto.Int32(cityX),
				CityY:  proto.Int32(cityY),
			}
			if err != nil {
				log.Println("Sql error: ", err)
				return
			}
			log.Println("Rename.user", user)
			rows.Close()
			ss.SetUid(uid)
			break
		}
		if newUser {
			name := "user_" + passport

			tx, err := Sql.Begin()
			if err != nil {
				log.Println("Sql error: ", err)
			}
			res, _ := tx.Exec(`insert into u_user(uid, passport,sid,name,gender,icon, 
			iconB,level,exp,gold,regTime, regIp,loginTime, loginIp,status,vipLevel,
			 dayTime,renames,combat,cityX,cityY) values(null,?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?, ?,?,?,?,?)`,
				passport, 2148999, name, 1, 0, 0, 1, 0, 0, now, "ip", now, "ip", 0, 0, now, 0, 0, 0, 0)
			uid, _ = res.LastInsertId()
			fmt.Println("User.OnUserNew", uid)
			tx.Commit()
			//==OnUserNew(uid)
			Event.Call("OnUserNew", uid)

			goto LAB_USER
		} else {
			//updateLoginTime
		}
		uid = user.GetUid()

		fmt.Println("User.OnLogin", uid)
		Event.Call("OnLogin", uid)

		fmt.Println("User.OnUserInit", uid)
		updates := &protos.Updates{}
		updates.User = &user
		Event.Call("OnUserInit", uid, updates)

		ss.CallOut(protoId+1, &protos.Response_S{ProtoId: proto.Int32(protoId),
			Updates: updates,
		})
	})

	Net.RegRPC(Rpc.Rename_C, func(ss *Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.Rename_C{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		fmt.Println("<<<Rename", data, ps.GetName())
		Sql.Exec("update u_user set name=? where uid=?", ps.GetName(), uid)
		user := &protos.User{
			Uid:  proto.Int64(uid),
			Name: ps.Name,
		}
		ss.Props().User = user
	})
	Net.RegRPC(Rpc.ReIcon_C, func(ss *Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.ReIcon_C{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		fmt.Println("<<<ReIcon", data, ps.GetIcon())

		Sql.Exec("update u_user set icon=? where uid=?", ps.GetIcon(), uid)
		user := &protos.User{
			Uid:  proto.Int64(uid),
			Icon: ps.Icon,
		}
		ss.Props().User = user
	})
	Net.RegRPC(Rpc.ReIconB_C, func(ss *Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.ReIconB_C{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		fmt.Println("<<<ReIconB_C", data, ps.GetIconB())
		Sql.Exec("update u_user set iconB=? where uid=?", ps.GetIconB(), uid)
		user := &protos.User{
			Uid:   proto.Int64(uid),
			IconB: ps.IconB,
		}
		ss.Props().User = user
	})
	Net.RegRPC(Rpc.UserView_C, func(ss *Net.Session, protoId int32, uid int64, data []byte) {
		ps := protos.UserView_C{}
		if ss.DecodeFail(data, &ps) {
			return
		}
		a := []*protos.User{}

		rows, err := Sql.Query("select uid, name, gender, icon, iconB, level, cityX, cityY from u_user where uid=?", ps.GetUid())
		if err != nil {
			log.Println("Sql error: ", err)
			return
		}
		for rows.Next() {
			var uid int64
			var name []byte
			var gender, icon, iconB, level, cityX, cityY int32
			err = rows.Scan(&uid, &name, &gender, &icon, &iconB, &level, &cityX, &cityY)
			if err != nil {
				log.Println("Sql error: ", err)
				return
			}
			user := &protos.User{
				Uid:    proto.Int64(uid),
				Name:   proto.String(string(name)),
				Gender: proto.Int32(gender),
				Icon:   proto.Int32(icon),
				IconB:  proto.Int32(iconB),
				Level:  proto.Int32(level),
				CityX:  proto.Int32(cityX),
				CityY:  proto.Int32(cityY),
			}
			a = append(a, user)
		}
		ss.Update().Other = a
	})

}
