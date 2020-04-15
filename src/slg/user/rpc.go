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
		} else {
			name := "user_" + passport
			user = Entity.User{
				Passport: passport,
				Name:     name,
				Level:    1,
				Version:  0,
				Login:    now,
			}
			x.Insert(&user)
			uid = user.Uid
			log.Println("Event.OnUserNew", uid)

			//==OnUserNew(uid)
			Event.Call(Const.OnUserNew, uid)

		}
		uid = user.Uid
		//updateLoginTime

		log.Println("Event.OnUserLogin", uid)
		Event.Call(Const.OnUserLogin, uid)

		updates := &protos.Updates{}
		updates.User = user.ToProto()
		Event.Call(Const.OnUserGetData, uid, updates)

		ss.CallOut(protoId+1, &protos.Response_S{ProtoId: proto.Int32(protoId),
			Updates: updates,
		})
	})

	Net.RegRPC(Const.Rename_C, func(ss Net.Session, protoId int32, uid int64, data []byte) {
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
	Net.RegRPC(Const.ReIcon_C, func(ss Net.Session, protoId int32, uid int64, data []byte) {
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
	Net.RegRPC(Const.ReIconB_C, func(ss Net.Session, protoId int32, uid int64, data []byte) {
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
	Net.RegRPC(Const.UserView_C, func(ss Net.Session, protoId int32, uid int64, data []byte) {
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
