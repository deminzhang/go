package Entity

//世界地图分区
type Area struct {
	Cid    int32   `xorm:"pk"`
	Col    int32   `xorm:"index(pos)"`
	Row    int32   `xorm:"index(pos)"`
	Inited bool    `xorm:"-"`
	Tiles  []*Tile `xorm:"-"`
}

//返回主键
func (this *Area) GetPK() int32 {
	return this.Cid
}
