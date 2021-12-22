package main

import (
	"fmt"
	"math"
)

func main() {
	pow1, pow2 := powerSeries(2)
	fmt.Println(pow1, pow2)
}

func powerSeries(num float64) (float64, float64) {
	return math.Pow(num, 2), math.Pow(num, 3)
}
