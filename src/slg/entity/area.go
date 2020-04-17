package Entity

//世界地图分区
type Area struct {
	Sid    int32   `xorm:"pk"`
	Col    int32   `xorm:"index(pos)"`
	Row    int32   `xorm:"index(pos)"`
	Inited bool    `xorm:"-"`
	Tiles  []*Tile `xorm:"-"`
}
