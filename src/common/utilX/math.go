package utilX

import (
	"math"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Min(a, b float64, more ...float64) float64 {
	min := math.Min(a, b)
	for _, n := range more {
		min = math.Min(min, n)
	}
	return min
}

func Max(a, b float64, more ...float64) float64 {
	min := math.Max(a, b)
	for _, n := range more {
		min = math.Max(min, n)
	}
	return min
}
