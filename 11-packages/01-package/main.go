package main

import (
	"fmt"
	"introduction-to-go/11-packages/01-package/math"
)

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}