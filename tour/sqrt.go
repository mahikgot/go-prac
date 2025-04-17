package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
	diff := math.MaxFloat64
	res := 0.0
	for z, last := 1.0, 1.0; diff >= 1; res = z {
		last = z
		z -= (z*z - x) / (2 * z)
		diff = math.Abs(last - z)
	}
	return res
}

func main() {
	fmt.Println(sqrt(16))
}
