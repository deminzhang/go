package Entity

//邮件
type Mail struct {
	Sid      int64 `xorm:"pk autoincr"`
	Uid      int64 `xorm:"index"`
	Type     int32
	Cid      int32
	FromUid  int64
	FromName string
	Title    string
	Content  string
	Time     int64
	Read     bool
	Take     bool
	Favor    bool
	// Item
	// Res
}
