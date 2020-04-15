package Entity

//地板格
type Troop struct {
	Sid    int64 `xorm:"pk autoincr"`
	Uid    int64 `xorm:"index"`
	Create int64 `xorm:"created"`
	tp     int32 //行动类型
	stat   int32 //行动状态

	sx int32 //起始坐标x,y
	sy int32
	tx int32 //目标坐标x,y
	ty int32

	st int64 //起始时间加速会校正保持行程比 已走路程/全程==(now-st)/(et-st)
	et int64 //预计到达时间

	lsid int64 `xorm:"FlagShip"` //集结领队sid
	// ttp  int32 //目标类型
	// tval int64 //目标值
	// sumTime int64  //初始总时间,前端用于显示总进度
	// hero int64[]//英雄id(主将,副将,副将)
	// heroList Hero `xorm:"-"`

	// //    private byte[] unit;//兵种数量
	//     private List<DataUnit> unit = new ArrayList<>();
	//     private byte[] res;//携带资源 TODO 改成 List
	//     private List<DataRes> resList = new ArrayList<>();

}

//返回主键
func (this *Troop) GetPK() int64 {
	return this.Sid
}
