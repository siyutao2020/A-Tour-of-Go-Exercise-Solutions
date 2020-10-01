package main

import (
	"fmt"
)

const thres = 1e-16

func Sqrt(x float64) float64 {
	z := 1.0
	for (z*z-x)*(z*z-x) > thres {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
