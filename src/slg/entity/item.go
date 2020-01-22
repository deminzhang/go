package Entity

//道具实例
type Item struct {
	Sid  int64
	Uid  int64
	Cid  int32
	Num  int64
	Time int64
}

//返回主键
func (this *Item) GetPrimaryKey() int64 {
	return this.Sid
}

//返回分组键 owner
func (this *Item) GetUserKey() int64 {
	return this.Uid
}
