package Entity

//联盟/公会/帮会/组织
type Alliance struct {
	Sid        int64  `xorm:"pk"`
	Name       string `xorm:"unique"`
	Flag       int32
	Notice     string
	Leader     int64
	LeaderName string
	Member     int32
	Version    int32 `xorm:"version"`
}

//返回主键
func (this *Alliance) GetPK() int64 {
	return this.Sid
}
