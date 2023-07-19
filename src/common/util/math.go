package util

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](a0 T, a ...T) T {
	r := a0
	for _, t := range a {
		if r < t {
			r = t
		}
	}
	return r
}

func Min[T constraints.Ordered](a0 T, a ...T) T {
	r := a0
	for _, t := range a {
		if r > t {
			r = t
		}
	}
	return r
}

func Abs[T constraints.Signed | constraints.Float](t T) T {
	if t < 0 {
		return -t
	}
	return t
}

// Clamp clamps the value of x to within min and max.
func Clamp[T constraints.Ordered](x, min, max T) T {
	if x < min {
		return min
	}
	if x > max {
		return max
	}
	return x
}

func If[T any](cond bool, trueVal T, falseVal T) T {
	if cond {
		return trueVal
	} else {
		return falseVal
	}
}
