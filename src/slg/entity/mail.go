package Entity

import (
	"protos"

	"github.com/golang/protobuf/proto"
)

//邮件
type Mail struct {
	Sid      int64   `xorm:"pk autoincr"`
	Uid      int64   `xorm:"index"`
	Type     int32   //发送分类
	Cid      int32   //多语言配置ID
	FromUid  int64   //发送者UID
	FromName string  //发送者名
	Title    string  //标题[]
	Context  string  //正文[]
	Time     int64   //发送时间
	Read     bool    //已读
	Take     bool    //已取附件
	Favor    bool    //收藏
	Item     []int32 //附件道具id1,num1,id2,num,...
	// Res
}

//转proto对象
func (this *Mail) ToProto() *protos.Mail {
	return &protos.Mail{
		Sid:      proto.Int64(this.Sid),
		Tp:       proto.Int32(this.Type),
		Cid:      proto.Int32(this.Cid),
		FromUid:  proto.Int64(this.FromUid),
		FromName: proto.String(this.FromName),
		Title:    proto.String(this.Title),
		Content:  proto.String(this.Context),
		Time:     proto.Int64(this.Time),
		Read:     proto.Bool(this.Read),
		Take:     proto.Bool(this.Take),
		Favor:    proto.Bool(this.Favor),
	}
}

//转proto前端主键对象
func (this *Mail) ToProtoPK() *protos.MailPK {
	return &protos.MailPK{
		Sid: proto.Int64(this.Sid),
	}
}

//加到更新
func (this *Mail) AppendTo(updates *protos.Updates) {
	list := updates.Mail
	if list == nil {
		list = []*protos.Mail{}
	}
	updates.Mail = append(list, this.ToProto())
}

//加到删除
func (this *Mail) AppendToPK(removes *protos.Removes) {
	list := removes.Mail
	if list == nil {
		list = []*protos.MailPK{}
	}
	removes.Mail = append(list, this.ToProtoPK())
}
