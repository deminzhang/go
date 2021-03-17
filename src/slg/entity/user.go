package Entity

import (
	"protos"
)

//角色数据
type User struct {
	Uid        int64  `xorm:"pk autoincr"`
	Passport   string `xorm:"index"`
	Name       string `xorm:"unique"`
	Gender     int32
	Head       int32
	HeadB      int32
	Level      int32 //`xorm:"index"`
	CityX      int32 `xorm:"index(pos)"`
	CityY      int32 `xorm:"index(pos)"`
	CreateTime int64 `xorm:"created"`
	LoginTime  int64 `xorm:"->"`
	AllianceId int64
	Version    int `xorm:"version"` //乐观锁
}

//转proto对象
func (this *User) ToProto() *protos.User {
	return &protos.User{
		Uid:        this.Uid,
		Name:       this.Name,
		Gender:     this.Gender,
		Icon:       this.Head,
		IconB:      this.HeadB,
		Level:      this.Level,
		CityX:      this.CityX,
		CityY:      this.CityY,
		AllianceId: this.AllianceId,
	}
}

//转proto前端主键对象
func (this *User) ToProtoPK() *protos.UserPK {
	return &protos.UserPK{
		Uid: this.Uid,
	}
}

//加到更新
func (this *User) AppendTo(updates *protos.Updates) {
	list := updates.User
	if list == nil {
		list = []*protos.User{}
	}
	updates.User = append(list, this.ToProto())
}

//加到删除
// func (this *User) AppendToPK(removes *protos.Removes) {
// 	list := removes.User
// 	if list == nil {
// 		list = []*protos.UserPK{}
// 	}
// 	removes.User = append(list, this.ToProtoPK())
// }
