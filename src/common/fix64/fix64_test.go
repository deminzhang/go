// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fix64_test

import (
	. "common/fix64"
	"fmt"
	"golang.org/x/exp/constraints"
	"math"
	"math/rand"
	"testing"
)

//func FastCosM(val Fix64) Fix64 {
//	var temp Fix64
//	if val > 0 {
//		temp = ((-FixPi).Sub(FixHalfPi))
//	} else {
//		temp = FixHalfPi
//	}
//	return Sin1(val.AddOrUpdate(temp))
//}

//var sinLut [LutSize]int64
//
//func GenerateSinLut() {
//	for i := 0; i < int(LutSize); i++ {
//		var angle = float64((float64(i) * float64(math.Pi) * 0.5 / float64(LutSize-1)))
//		var sin = math.Sin(angle)
//		sinLut[i] = int64(NewFix64(sin))
//	}
//}
//
//func Sin1(val Fix64) Fix64 {
//	var clampedL, flipHorizontal, flipVertical = ClampSinValue(val)
//	// var rawIndex = (int64)(clampedL >> 3)
//	var rawIndex = int64(clampedL)
//	if rawIndex >= LutSize {
//		rawIndex = LutSize - 1
//	}
//	var nearestValue = sinLut[rawIndex]
//	if flipHorizontal {
//		nearestValue = sinLut[LutSize-1-rawIndex]
//	}
//	// fmt.Println(nearestValue, LutSize-1-rawIndex, rawIndex)
//	if flipVertical {
//		return Fix64(-nearestValue)
//	} else {
//		return Fix64(nearestValue)
//	}
//}

func ClampSinValue(val Fix64) (Fix64, bool, bool) {
	var clamped2Pi = val % FixPi2
	if val < 0 {
		clamped2Pi += FixPi2
	}
	var clampedPi = clamped2Pi
	for {
		if clampedPi >= FixPi {
			clampedPi -= FixPi
		} else {
			break
		}
	}
	var clampedPiOver2 = clampedPi
	if clampedPiOver2 >= FixHalfPi {
		clampedPiOver2 -= FixHalfPi
	}
	return clampedPiOver2, (clampedPi >= FixHalfPi), (clamped2Pi >= FixPi)
}

var testCases = []struct {
	x      float64
	sFix64 string
	floor  int64
	round  int64
	ceil   int64
}{{
	x:      0,
	sFix64: "0:0000000000",
	floor:  0,
	round:  0,
	ceil:   0,
}, {
	x:      1,
	sFix64: "1:0000000000",
	floor:  1,
	round:  1,
	ceil:   1,
}, {
	x:      1.25,
	sFix64: "1:1073741824",
	floor:  1,
	round:  1,
	ceil:   2,
}, {
	x:      2.5,
	sFix64: "2:2147483648",
	floor:  2,
	round:  3,
	ceil:   3,
}, {
	x:      63 / 64.0,
	sFix64: "0:4227858432",
	floor:  0,
	round:  1,
	ceil:   1,
}, {
	x:      -0.5,
	sFix64: "-0:2147483648",
	floor:  -1,
	round:  +0,
	ceil:   +0,
}, {
	x:      -4.125,
	sFix64: "-4:0536870912",
	floor:  -5,
	round:  -4,
	ceil:   -4,
}, {
	x:      -7.75,
	sFix64: "-7:3221225472",
	floor:  -8,
	round:  -8,
	ceil:   -7,
}}

func TestFix64(t *testing.T) {
	const one = Fix64(1 << 32)
	for _, tc := range testCases {
		x := Fix64(tc.x * (1 << 32))
		if got, want := x.String(), tc.sFix64; got != want {
			t.Errorf("tc.x=%v: String: got %q, want %q", tc.x, got, want)
		}
		if got, want := x.Floor(), tc.floor; got != want {
			t.Errorf("tc.x=%v: Floor: got %v, want %v", tc.x, got, want)
		}
		if got, want := x.Round(), tc.round; got != want {
			t.Errorf("tc.x=%v: Round: got %v, want %v", tc.x, got, want)
		}
		if got, want := x.Ceil(), tc.ceil; got != want {
			t.Errorf("tc.x=%v: Ceil: got %v, want %v", tc.x, got, want)
		}
		if got, want := x.Mul(one), x; got != want {
			t.Errorf("tc.x=%v: Mul by one: got %v, want %v", tc.x, got, want)
		}
	}
}

var mulTestCases = []struct {
	x      float64
	y      float64
	zFix64 float64 // Equals truncate52_12(x)*truncate52_12(y).
	sFix64 string
}{{
	x:      0,
	y:      1.5,
	zFix64: 0,
	sFix64: "0:0000000000",
}, {
	x:      +1.25,
	y:      +4,
	zFix64: +5,
	sFix64: "5:0000000000",
}, {
	x:      +1.25,
	y:      -4,
	zFix64: -5,
	sFix64: "-5:0000000000",
}, {
	x:      -1.25,
	y:      +4,
	zFix64: -5,
	sFix64: "-5:0000000000",
}, {
	x:      -1.25,
	y:      -4,
	zFix64: +5,
	sFix64: "5:0000000000",
}, {
	x:      1.25,
	y:      1.5,
	zFix64: 1.875,
	sFix64: "1:3758096384",
}, {
	x:      1234.5,
	y:      -8888.875,
	zFix64: -10973316.1875,
	sFix64: "-10973316:0805306368",
}, {
	x:      1.515625,       // 1 + 33/64 = 97/64
	y:      1.531250,       // 1 + 34/64 = 98/64
	zFix64: 2.32080078125,  // 2 + 1314/4096 = 9506/4096
	sFix64: "2:1377828864", // 2.32080078125
}, {
	x:      0.500244140625,     // 2049/4096, approximately 32/64
	y:      0.500732421875,     // 2051/4096, approximately 32/64
	zFix64: 0.2504884600639343, // 4202499/16777216
	sFix64: "0:1075839744",     // 0.25048828125, which is closer than 0:1027 (in decimal, 0.250732421875)
}, {
	x:      0.015625,             // 1/64
	y:      0.000244140625,       // 1/4096, approximately 0/64
	zFix64: 0.000003814697265625, // 1/262144
	sFix64: "0:0000016384",       // 0, which is closer than 0:0001 (in decimal, 0.000244140625)
}, {
	// Round the fix64 calculation down.
	x:      1.44140625,         // 1 + 1808/4096 = 5904/4096, approximately 92/64
	y:      1.44140625,         // 1 + 1808/4096 = 5904/4096, approximately 92/64
	zFix64: 2.0776519775390625, // 2 +  318/4096 +  256/16777216 = 34857216/16777216
	sFix64: "2:0333512704",     // 2.07763671875, which is closer than 2:0319 (in decimal, 2.077880859375)
}, {
	// Round the fix64 calculation up.
	x:      1.44140625,         // 1 + 1808/4096 = 5904/4096, approximately 92/64
	y:      1.441650390625,     // 1 + 1809/4096 = 5905/4096, approximately 92/64
	zFix64: 2.0780038833618164, // 2 +  319/4096 + 2064/16777216 = 34863120/16777216
	sFix64: "2:0335024128",     // 2.07812500000, which is closer than 2:0319 (in decimal, 2.077880859375)
}}

func TestFix64Mul(t *testing.T) {
	for _, tc := range mulTestCases {
		x := Fix64(tc.x * (1 << 32))
		y := Fix64(tc.y * (1 << 32))
		if z := float64(x) * float64(y) / (1 << 64); z != tc.zFix64 {
			t.Errorf("tc.x=%v, tc.y=%v: z: got %v, want %v", tc.x, tc.y, z, tc.zFix64)
			continue
		}
		if got, want := x.Mul(y).String(), tc.sFix64; got != want {
			t.Errorf("tc.x=%v: Mul: got %q, want %q", tc.x, got, want)
		}
	}
}

func TestFix64MulByOneMinusIota(t *testing.T) {
	const (
		totalBits = 64
		fracBits  = 32

		oneMinusIota  = Fix64(1<<fracBits) - 1
		oneMinusIotaF = float64(oneMinusIota) / (1 << fracBits)
	)

	for _, neg := range []bool{false, true} {
		for i := uint(0); i < totalBits; i++ {
			x := Fix64(1 << i)
			if neg {
				x = -x
			} else if i == totalBits-1 {
				// A signed int64 can't represent 1<<63.
				continue
			}

			// want equals x * oneMinusIota, rounded to nearest.
			want := Fix64(0)
			if -1<<fracBits < x && x < 1<<fracBits {
				// (x * oneMinusIota) isn't exactly representable as an
				// fix64. Calculate the rounded value using float64 math.
				xF := float64(x) / (1 << fracBits)
				wantF := xF * oneMinusIotaF * (1 << fracBits)
				want = Fix64(math.Floor(wantF + 0.5))
			} else {
				// (x * oneMinusIota) is exactly representable.
				want = oneMinusIota << (i - fracBits)
				if neg {
					want = -want
				}
			}

			if got := x.Mul(oneMinusIota); got != want {
				t.Errorf("neg=%t, i=%d, x=%v, Mul: got %v, want %v", neg, i, x, got, want)
			}
		}
	}
}

func TestMuli32(t *testing.T) {
	rng := rand.New(rand.NewSource(2))
	for i := 0; i < 10000; i++ {
		u := int32(rng.Uint32())
		v := int32(rng.Uint32())
		lo, hi := muli32(u, v)
		got := uint64(lo) | uint64(hi)<<32
		want := uint64(int64(u) * int64(v))
		if got != want {
			t.Errorf("u=%#08x, v=%#08x: got %#016x, want %#016x", uint32(u), uint32(v), got, want)
		}
	}
}

func TestMulu32(t *testing.T) {
	rng := rand.New(rand.NewSource(3))
	for i := 0; i < 10000; i++ {
		u := rng.Uint32()
		v := rng.Uint32()
		lo, hi := mulu32(u, v)
		got := uint64(lo) | uint64(hi)<<32
		want := uint64(u) * uint64(v)
		if got != want {
			t.Errorf("u=%#08x, v=%#08x: got %#016x, want %#016x", u, v, got, want)
		}
	}
}

// muli32 multiplies two int32 values, returning the 64-bit signed integer
// result as two uint32 values.
//
// muli32 isn't used directly by this package, but it has the same structure as
// muli64, and muli32 is easier to test since Go has built-in 64-bit integers.
func muli32(u, v int32) (lo, hi uint32) {
	const (
		s    = 16
		mask = 1<<s - 1
	)

	u1 := uint32(u >> s)
	u0 := uint32(u & mask)
	v1 := uint32(v >> s)
	v0 := uint32(v & mask)

	w0 := u0 * v0
	t := u1*v0 + w0>>s
	w1 := t & mask
	w2 := uint32(int32(t) >> s)
	w1 += u0 * v1
	return uint32(u) * uint32(v), u1*v1 + w2 + uint32(int32(w1)>>s)
}

// mulu32 is like muli32, except that it multiplies unsigned instead of signed
// values.
//
// This implementation comes from $GOROOT/src/runtime/softfloat64.go's mullu
// function, which is in turn adapted from Hacker's Delight.
//
// mulu32 (and its corresponding test, TestMulu32) isn't used directly by this
// package. It is provided in this test file as a reference point to compare
// the muli32 (and TestMuli32) implementations against.
func mulu32(u, v uint32) (lo, hi uint32) {
	const (
		s    = 16
		mask = 1<<s - 1
	)

	u0 := u & mask
	u1 := u >> s
	v0 := v & mask
	v1 := v >> s

	w0 := u0 * v0
	t := u1*v0 + w0>>s
	w1 := t & mask
	w2 := t >> s
	w1 += u0 * v1
	return u * v, u1*v1 + w2 + w1>>s
}

var testSqrt = []struct {
	x    float64
	sRst string
	fRst float64
}{
	{x: 2, sRst: "1:1779033704", fRst: 1.414},
	{x: 3, sRst: "1:3144134278", fRst: 1.732},
	{x: 2.5, sRst: "1:2495972270", fRst: 1.581},
	{x: 1.9, sRst: "1:1625236564", fRst: 1.378},
}

func TestSqrt(t *testing.T) {
	for _, ts := range testSqrt {
		x := NewFix64(ts.x)
		if got, want := x.Sqrt().String(), ts.sRst; got != want {
			t.Errorf("tc.x=%v: Sqrt: got %q, want %q", ts.x, got, want)
		}
	}
}

var testSin = []struct {
	x    float64
	sRst string
	fRst float64
}{
	//{x: 0.785, sRst: "-0:4080", fRst: 1.414},
	{x: 0.7386, sRst: "0:2891675434", fRst: 0.67327065245941378082211628671578},
	//{x: 1.058, sRst: "0:3547", fRst: 0.866},
}

func TestSin(t *testing.T) {

	//fmt.Println(sinTable)
	var Pi = 3.14159265358979323846264338327950288419716939937510582097494459
	//fmt.Println(NewFix64(Pi / 2).GetHashCode())
	//fmt.Println(NewFix64(Pi).GetHashCode())
	//fmt.Println(NewFix64(Pi / 2.0 * 3.0).GetHashCode())
	//fmt.Println(NewFix64(Pi * 2).GetHashCode())
	//for i := 0; i <= 90; i += 1 {
	//	//fmt.Println(NewFix64(float64(i) / 180.0 * Pi).Sin())
	//	fmt.Println(NewFix64(math.Sin(float64(float64(i) / 180.0 * Pi))))
	//}

	//fmt.Println(NewFix64(math.Sin(float64(float64(202.49999999216567) / 180.0 * Pi))))

	for i := FixZero; i < Fix64(26986075409); i += 168662971 {
		if got, want := i.Sin(), NewFix64(math.Sin(i.Float64())); true /*math.Abs(float64(got-want)) > float64(Fix64(1))*/ {
			t.Errorf("tc.x=%v: Sin: got %q, want %q", i.Float64()/Pi*180, got, want)
		}
	}

	//for _, ts := range testSin {
	//	x := NewFix64(ts.x)
	//	if got, want := Sin(x).String(), F(math.Sin(ts.x)).String(); got != want {
	//		t.Errorf("tc.x=%v: Sin: got %q, want %q", ts.x, got, want)
	//	}
	//}
}

var testCos = []struct {
	x    float64
	sRst string
	fRst float64
}{
	{x: 0.5238, sRst: "0:3547", fRst: 0.86592477405103730192926824175542},
}

func TestCos(t *testing.T) {
	var Pi = 3.14159265358979323846264338327950288419716939937510582097494459

	for i := FixZero; i < Fix64(26986075409); i += 1686629713 {
		if got, want := i.Cos(), NewFix64(math.Cos(i.Float64())); (got - want) > (Fix64(1)) {
			t.Errorf("tc.x=%v: Cos: got %q, want %q", i.Float64()/Pi*180, got, want)
		}
	}
}

var testTan = []struct {
	x    float64
	sRst string
	fRst float64
}{
	{x: 0.5238, sRst: "0:2364", fRst: 0.57761859956932406917327260226729},
}

func TestTan(t *testing.T) {
	for _, ts := range testTan {
		x := NewFix64(ts.x)
		if got, want := x.Tan().String(), ts.sRst; got != want {
			t.Errorf("tc.x=%v: Tan: got %q, want %q", ts.x, got, want)
		}
	}
}

func TestCreateSin(t *testing.T) {
	LutSize := 900
	var tscl []int64 = make([]int64, LutSize)

	var Pi = 3.14159265358979323846264338327950288419716939937510582097494459
	for i := 0; i < int(LutSize); i++ {
		var angle = float64((float64(i) * float64(Pi) * 0.5 / float64(LutSize-1)))
		tscl[i] = int64(NewFix64(math.Sin(angle)))
	}

	fmt.Println(tscl)

}

//func TestAA(t *testing.T) {
//
//	fmt.Println(Sin1(NewFix64(0.3234)))
//}

func BenchmarkSin(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		NewFix64(0.3234).Sin()
	}
}

//func BenchmarkSin1(b *testing.B) {
//	b.StopTimer()
//	b.StartTimer()
//	GenerateSinLut()
//	for i := 0; i < b.N; i++ {
//		Sin1(NewFix64(0.3234))
//	}
//}

func NewFF[T constraints.Float](val T) Fix64 {

	return 1
}

func TestCC(t *testing.T) {
	fmt.Println(math.Round(1.5))
	fmt.Println(0)
	fmt.Println(NewFix64(0.0))
	fmt.Println(NewFix64(0.5))
	fmt.Println(NewFix64(-0.5))
	fmt.Println(NewFix64(0.5))
	fmt.Println(NewFix64(0))
	fmt.Println(NewFF(123.0))
}

func TestNewFix64(t *testing.T) {
	fmt.Println(NewFix64(0.67327065245941378082211628671578))

}
