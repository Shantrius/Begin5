package main

import (
	"fmt"
	"math"
)

func main() {
	var A float64 = 6
	var V float64 = (math.Pow(A, 3))
	var S float64 = (6 * (math.Pow(A, 2)))
	fmt.Println(V, S)

}
