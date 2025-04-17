package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (errNegativeSqrt ErrNegativeSqrt) Error() string {
	return fmt.Sprint(float64(errNegativeSqrt))
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	diff := math.MaxFloat64
	res := 0.0
	for z, last := 1.0, 1.0; diff >= 1; res = z {
		last = z
		z -= (z*z - x) / (2 * z)
		diff = math.Abs(last - z)
	}
	return res, nil
}

func main() {
	fmt.Println(sqrt(-1))
	fmt.Println(sqrt(16))
}
