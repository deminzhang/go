package vec_test

import (
	. "common/fix64"
	. "common/vec"
	"fmt"
	"testing"
)

var mulTestCases = []struct {
	x float64
	y float64
	z float64
}{
	{
		x: 0,
		y: 1.5,
		z: 3,
	}, {
		x: +1.25,
		y: +4,
		z: +2.75,
	}, {
		x: +1.25,
		y: -4,
		z: -2.75,
	}, {
		x: -1.25,
		y: +4,
		z: +4.25,
	}, {
		x: -1.25,
		y: -4,
		z: -5,
	}, {
		x: 1.25,
		y: 1.5,
		z: 2,
	},
	{
		x: 1234.5,
		y: -8888.875,
		z: 345,
	},
	// {
	// 	x: 1.515625,
	// 	y: 1.531250,
	// 	z: 1.44140625,
	// },
	{
		x: 0.500244140625,
		y: 0.500732421875,
		z: 0.000244140625,
	}, {
		x: 0.015625,
		y: 0.000244140625,
		z: 1.44140625,
	}, {
		x: 1.44140625,
		y: 1.44140625,
		z: 1.441650390625,
	},
	{
		x: 1.44140625,
		y: 1.441650390625,
		z: 0.500244140625,
	},
	{
		x: 1231445,
		y: 4634356,
		z: 5675323,
	},
}

func TestVector3Fix(t *testing.T) {
	for _, tc := range mulTestCases {
		a := NewFixVector3(NewFix64(tc.x), NewFix64(tc.y), NewFix64(tc.z))
		fmt.Println(tc.x, tc.y, tc.z)
		fmt.Println(a.X, a.Y, a.Z)
		fmt.Println(NewFix64(tc.x), NewFix64(tc.y), NewFix64(tc.z))
		fmt.Println(a.Float())
		fmt.Println("---------------------------------------------------")
	}
	fmt.Println("TestVector3Add")
}

func TestVector3FixEqual(t *testing.T) {
	for _, tc := range mulTestCases {
		a := NewFixVector3(NewFix64(tc.x), NewFix64(tc.y), NewFix64(tc.z))
		b := NewFixVector3(NewFix64(tc.x), NewFix64(tc.y), NewFix64(tc.z))
		if !a.Equal(b) {
			t.Errorf("tc.x=%v, tc.y=%v, tc.z=%v", tc.x, tc.y, tc.z)
		}
	}
}

func TestVector3FixAdd(t *testing.T) {
	for i := 0; i < len(mulTestCases); i += 2 {
		tc := mulTestCases[i]
		td := mulTestCases[i+1]
		a := NewFixVector3(NewFix64(tc.x), NewFix64(tc.y), NewFix64(tc.z))
		b := NewFixVector3(NewFix64(td.x), NewFix64(td.y), NewFix64(td.z))
		a.Add(b)
		x, y, z := tc.x+td.x, tc.y+td.y, tc.z+td.z
		c := NewFixVector3(NewFix64(x), NewFix64(y), NewFix64(z))
		if !a.Equal(c) {
			t.Errorf("tc.x=%v, tc.y=%v, tc.z=%v", tc.x, tc.y, tc.z)
		}
	}
}

func TestVector3FixSub(t *testing.T) {
	for i := 0; i < len(mulTestCases); i += 2 {
		tc := mulTestCases[i]
		td := mulTestCases[i+1]
		a := NewFixVector3(NewFix64(tc.x), NewFix64(tc.y), NewFix64(tc.z))
		b := NewFixVector3(NewFix64(td.x), NewFix64(td.y), NewFix64(td.z))
		a.Sub(b)
		x, y, z := tc.x-td.x, tc.y-td.y, tc.z-td.z
		c := NewFixVector3(NewFix64(x), NewFix64(y), NewFix64(z))
		if !a.Equal(c) {
			t.Errorf("tc.x=%v, tc.y=%v, tc.z=%v", tc.x, tc.y, tc.z)
		}
	}
}

func TestVector3FixMul(t *testing.T) {
	for i := 0; i < len(mulTestCases); i += 2 {
		tc := mulTestCases[i]
		td := mulTestCases[i+1]
		a := NewFixVector3(NewFix64(tc.x), NewFix64(tc.y), NewFix64(tc.z))
		a.Multiply(NewFix64(td.x))
		x, y, z := tc.x*td.x, tc.y*td.x, tc.z*td.x
		c := NewFixVector3(NewFix64(x), NewFix64(y), NewFix64(z))
		if !a.Equal(c) {
			if a.X.Sub(c.X).GetHashCode() > 1 ||
				a.Y.Sub(c.Y).GetHashCode() > 1 ||
				a.Z.Sub(c.Z).GetHashCode() > 1 {
				t.Errorf("tc.x=%v, tc.y=%v, tc.z=%v", tc.x, tc.y, tc.z)
			}
		}
	}
}

func TestVector3FixDiv(t *testing.T) {
	for i := 0; i < len(mulTestCases); i += 2 {
		tc := mulTestCases[i]
		td := mulTestCases[i+1]
		a := NewFixVector3(NewFix64(tc.x), NewFix64(tc.y), NewFix64(tc.z))
		a.Divide(NewFix64(td.x))
		x, y, z := tc.x/td.x, tc.y/td.x, tc.z/td.x
		c := NewFixVector3(NewFix64(x), NewFix64(y), NewFix64(z))
		if !a.Equal(c) {
			if a.X.Sub(c.X).GetHashCode() > 1 ||
				a.Y.Sub(c.Y).GetHashCode() > 1 ||
				a.Z.Sub(c.Z).GetHashCode() > 1 {
				d := a.Clone()
				d.Sub(c)
				t.Errorf("tc.x=%v, tc.y=%v, tc.z=%v", tc.x, tc.y, tc.z)
				fmt.Println(d)
			}
		}
	}
}

func TestVector3FixDot(t *testing.T) {
	for i := 0; i < len(mulTestCases); i += 2 {
		tc := mulTestCases[i]
		td := mulTestCases[i+1]
		a := NewFixVector3(NewFix64(tc.x), NewFix64(tc.y), NewFix64(tc.z))
		b := NewFixVector3(NewFix64(td.x), NewFix64(td.y), NewFix64(td.z))

		c := a.Dot(b)
		d := NewFix64(tc.x*td.x + tc.y*td.y + tc.z*td.z)
		if c != d {
			t.Errorf("tc.x=%v, tc.y=%v, tc.z=%v", tc.x, tc.y, tc.z)
			fmt.Println(d.GetHashCode(), c.GetHashCode())
		}
	}
}

func TestVector3FixScale(t *testing.T) {
	v := NewFixVector3(FixOne, 0, 0)
	v2 := v.ScaledToLength(NewFix64(2))
	fmt.Println(v.X, v.Y, v.Z)
	fmt.Println(v2.X, v2.Y, v2.Z)
	fmt.Println(v2.X, v2.Y, v2.Z)
}

func TestVector3_RotateY(t *testing.T) {
	center := Vector3{}
	v3 := Vector3{X: 1, Z: 2}
	v := v3.RotateYAnticlockwise(center, 90)
	fmt.Println(v.X, v.Z)
	v = v3.RotateYAnticlockwise(center, 270)
	fmt.Println(v.X, v.Z)
	v = v3.RotateYAnticlockwise(center, 180)
	fmt.Println(v.X, v.Z)
}
