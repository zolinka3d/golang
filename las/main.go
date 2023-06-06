package main

import (
	"fmt"
	"las/forest"
)

func main() {

	fmt.Println(averageFiredForest(50, 10, 10))

}

func averageFiredForest(percent int, x int, y int) float64 {

	count := 1.0
	acc := 0.0
	average := 100.1
	averageLast := 100.2

	for abs(average-averageLast) > 0.000001 {
		averageLast = average
		acc = acc + forest.FireOneForest(percent, x, y)
		average = acc / count
		count++
		// fmt.Println(average)
	}
	return average
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
