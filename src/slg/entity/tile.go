package Entity

//地板格
type Tile struct {
	X       int32 `xorm:"pk"`
	Y       int32 `xorm:"pk"`
	Type    int32 `xorm:"index"`
	Troops  []int64
	Version int32 `xorm:"version"` //乐观锁
}

//返回主键
func (this *Tile) GetPK() (int32, int32) {
	return this.X, this.Y
}
