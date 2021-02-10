package common

import (
	"math"
)

func ComputeGCD(a int, b int) int {
	if a == 0 {
		return b;
	}
	return ComputeGCD(b % a, a)
}

func ToRadians(a float64) float64 {
	ans := a * math.Pi / 180
	return ans
}

func ToDegrees(a float64) float64 {
	ans := a * 180 / math.Pi
	return ans
}