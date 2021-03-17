package Entity

import (
	"protos"
	"sync"
)

//部队
type Troop struct {
	sync.RWMutex `xorm:"-"`
	Sid          int64 `xorm:"pk autoincr"`
	Uid          int64 `xorm:"index"`
	Create       int64 `xorm:"created"`
	Tp           int32 //行动类型
	Stat         int32 //行动状态

	Sx int32 //起始坐标x,y
	Sy int32
	Tx int32 //目标坐标x,y
	Ty int32

	St int64 //起始时间加速会校正保持行程比 已走路程/全程==(now-st)/(et-st)
	Et int64 //预计到达时间

	Lsid int64 `xorm:"index(FlagShip)"` //集结领队sid
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

//转proto对象
func (this *Troop) ToProto() *protos.Troop {
	return &protos.Troop{
		Sid:  this.Sid,
		Uid:  this.Uid,
		Tp:   this.Tp,
		Stat: this.Stat,
		Sx:   this.Sx,
		Sy:   this.Sy,
		Tx:   this.Tx,
		Ty:   this.Ty,
		St:   this.St,
		Et:   this.Et,
		Lsid: this.Lsid,
	}
}

//转proto前端主键对象
func (this *Troop) ToProtoPK() *protos.TroopPK {
	return &protos.TroopPK{
		Sid: this.Sid,
	}
}

//加到更新
func (this *Troop) AppendTo(updates *protos.Updates) {
	list := updates.Troop
	if list == nil {
		list = []*protos.Troop{}
	}
	updates.Troop = append(list, this.ToProto())
}

//加到删除
func (this *Troop) AppendToPK(removes *protos.Removes) {
	list := removes.Troop
	if list == nil {
		list = []*protos.TroopPK{}
	}
	removes.Troop = append(list, this.ToProtoPK())
}
