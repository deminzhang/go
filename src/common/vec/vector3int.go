package vec

import (
	"common/proto/comm"
	"common/table"
	"common/util"
)

type POSKEY int64

type Vector3Int struct {
	X int32
	Y int32
	Z int32
}

func (this *Vector3Int) Set(x, y, z int32) {
	this.X, this.Y, this.Z = x, y, z
}

func (this *Vector3Int) Clone() Vector3Int {
	return Vector3Int{X: this.X, Y: this.Y, Z: this.Z}
}

func (this *Vector3Int) Add(v Vector3Int) {
	this.X += v.X
	this.Y += v.Y
	this.Z += v.Z
}

func (this *Vector3Int) Sub(v Vector3Int) {
	this.X -= v.X
	this.Y -= v.Y
	this.Z -= v.Z
}

func (this *Vector3Int) Equal(v Vector3Int) bool {
	//return this.X == v.X && this.Y == v.Y && this.Z == v.Z
	return *this == v
}

func (this *Vector3Int) ManhattanDistance(v Vector3Int) int32 {
	return util.Abs(this.X-v.X) + util.Abs(this.Y-v.Y) + util.Abs(this.Z-v.Z)
}

func (this *Vector3Int) Key() POSKEY {
	return POSKEY(this.Z)<<32 | POSKEY(this.Y)<<16 | POSKEY(this.X)
}

func (this *Vector3Int) ShortKey() uint32 {
	return uint32(this.Z)<<16 | uint32(this.X)
}

func (this *Vector3Int) ToVector2Int() Vector2Int { return Vector2Int{X: this.X, Y: this.Z} }

func (this *Vector3Int) ToProtoVec3() *comm.Vec3 {
	return &comm.Vec3{X: int64(this.X), Y: int64(this.Y), Z: int64(this.Z)}
}

func (this *Vector3Int) SetProtoVec3(v *comm.Vec3) {
	this.Set(int32(v.X), int32(v.Y), int32(v.Z))
}

func (this *Vector3Int) RotateX(angle int) Vector3Int {
	s := table.SinByAngle(angle)
	c := table.CosByAngle(angle)

	x := float32(this.X)*c - float32(this.Z)*s
	z := float32(this.X)*s + float32(this.Z)*c

	return Vector3Int{X: int32(x), Y: this.Y, Z: int32(z)}
}

// ZeroV3Int 返回：零向量(0,0,0)
func ZeroV3Int() Vector3Int {
	return Vector3Int{X: 0, Y: 0, Z: 0}
}
