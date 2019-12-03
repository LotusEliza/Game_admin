package tools

import (
	"math/rand"
)

func RandMinMax(min, max int) int {
	return rand.Intn(max-min) + min
}
