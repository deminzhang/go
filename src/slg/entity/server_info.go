package Entity

import (
	"common/util"
	"protos"

	"github.com/golang/protobuf/proto"
)

type ServerInfo struct {
	Sid    int32 `xorm:"pk"`
	Minute int64
}

//转proto对象
func (this *ServerInfo) ToProto() *protos.Server {
	return &protos.Server{
		Region: proto.Int32(this.Sid),
		Time:   proto.Int64(Util.MilliSecond()),
	}
}

//加到更新
func (this *ServerInfo) AppendTo(updates *protos.Updates) {
	updates.Server = this.ToProto()
}
