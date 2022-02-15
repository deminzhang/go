package Entity

import (
	"common/utilX"
	"protos"
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
		Region: this.Sid,
		Time:   utilX.MilliSecond(),
	}
}
