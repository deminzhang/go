package util

import (
	"common/proto/comm"
	"math"
)

const (
	float_precision = 0.00001
)

func IsEqualVec3(v1 *comm.Vec3, v2 *comm.Vec3) bool {
	if v1 == v2 {
		return true
	}

	if math.Abs(float64(v1.X-v2.X)) <= float_precision &&
		math.Abs(float64(v1.Y-v2.Y)) <= float_precision &&
		math.Abs(float64(v1.Z-v2.Z)) <= float_precision {
		return true
	}
	return false
}

func IsEqualVec3f(v1 *comm.Vec3F, v2 *comm.Vec3F) bool {
	if v1 == v2 {
		return true
	}

	if math.Abs(float64(v1.X-v2.X)) <= float_precision &&
		math.Abs(float64(v1.Y-v2.Y)) <= float_precision &&
		math.Abs(float64(v1.Z-v2.Z)) <= float_precision {
		return true
	}
	return false
}

func CopyVec3f(v *comm.Vec3F) *comm.Vec3F {
	v1 := &comm.Vec3F{
		X: v.X,
		Y: v.Y,
		Z: v.Z,
	}
	return v1
}
