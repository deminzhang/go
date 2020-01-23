package Entity

//世界地图分区
type Area struct {
	cid    int32 `xorm:"pk"`
	col    int32 `xorm:"index(pos)"`
	row    int32 `xorm:"index(pos)"`
	inited bool
	citys  []int
	tiles  []*Tile
}

//返回主键
func (this *Area) GetPK() int32 {
	return this.cid
}
