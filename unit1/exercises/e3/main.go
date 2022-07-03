package main

import (
	"fmt"
)

/* func Sqrt(x float64) (float64, error) {
	return 0, nil
}*/

type ErrNegativeSqrt float64

var e ErrNegativeSqrt

func Sqrt(x float64) (float64, error) {
	var z float64
	z = x
	if z < 0 {
		return 0, e
	}
	for i := 1; i < 20; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: -2")
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
