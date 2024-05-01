package utils

import "math"

// const (
// 	Deg2Rad float32 = (math.Pi * 2) / 360.0
// 	Rad2Deg float32 = 360.0 / (math.Pi * 2)
// )

func Cos(value float32) float32 {
	return float32(math.Cos(float64(value)))
}

func Sin(value float32) float32 {
	return float32(math.Sin(float64(value)))
}
