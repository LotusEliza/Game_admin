package tools

import "math"

func Distance(ax, ay, bx, by int) float64 {
	return math.Sqrt(math.Pow(float64(ax)-float64(bx), 2) + math.Pow(float64(ay)-float64(by), 2))
}
