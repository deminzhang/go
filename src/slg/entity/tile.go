package Entity

//地板格
type Tile struct {
	X        int32 `xorm:"pk"`
	Y        int32 `xorm:"pk"`
	Tp       int32 `xorm:"index(type)"`
	Tp2      int32 `xorm:"index(type)"`
	Level    int32 `xorm:"index(type)"`
	Uid      int64 `xorm:"index(uid)"`
	Alliance int64 `xorm:"index(aid)"`
	Troops   []int64
	Version  int32 `xorm:"version"` //乐观锁
}

//返回主键
func (this *Tile) GetPK() (int32, int32) {
	return this.X, this.Y
}
