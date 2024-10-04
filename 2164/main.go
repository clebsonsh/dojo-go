package main

import (
	"fmt"
	"math"
)

// https://judge.beecrowd.com/pt/problems/view/2164
func main() {
	var n float64 = 0
	fmt.Scanln(&n)

	var output float64 = ((math.Pow((1+math.Sqrt(5))/2, n)) - (math.Pow((1-math.Sqrt(5))/2, n))) / math.Sqrt(5)
	fmt.Printf("%.1f\n", output)
}
