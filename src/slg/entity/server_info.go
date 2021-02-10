package Entity

import (
	"common/util"
	"protos"

	"github.com/golang/protobuf/proto"
)

type ServerInfo struct {
	Sid    int32 `xorm:"pk"`
	Minute int64 //
	// Time     int64 //对表时间现取
	VersionH int32 `xorm:"-"` //大版本
	VersionL int32 `xorm:"-"` //小版本
}

//转proto对象
func (this *ServerInfo) ToProto() *protos.Server {
	return &protos.Server{
		Region: proto.Int32(this.Sid),
		Time:   proto.Int64(Util.MilliSecond()),
	}
}
