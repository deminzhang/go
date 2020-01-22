package Entity

//角色数据
type User struct {
	Uid        int64
	Name       string `xorm:"unique"`
	Head       int32
	HeadB      int32
	Level      int32
	CityX      int32
	CityY      int32
	AllianceId int64
}

//返回主键
func (this *User) GetPrimaryKey() int64 {
	return this.Uid
}
