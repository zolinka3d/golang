package main

import (
	"fmt"
)

func main() {

	// var exampleArray [5]int
	exampleSlice := []int{1, 0, 4, 5, 6}
	// fmt.Println(exampleSlice)
	// fmt.Println(exampleArray)
	exampleSlice = append(exampleSlice, 100)
	// fmt.Println(exampleSlice)
	// kopiowanie slice
	exampleSliceCopy := make([]int, len(exampleSlice))
	copy(exampleSliceCopy, exampleSlice)
	// fmt.Println(exampleSliceCopy)

	// tempZmienna := exampleSlice

	// exampleSlice = append(exampleSlice, 300)
	// fmt.Println(exampleSlice)
	// tempZmienna = append(tempZmienna, 200)
	// fmt.Println(tempZmienna)
	// fmt.Println(exampleSlice)
	// exampleSlice = append(exampleSlice, 400)
	// tempZmienna = append(tempZmienna, 500)
	// fmt.Print(tempZmienna)
	// fmt.Print(exampleSlice)

	// last := exampleSlice[len(exampleSlice)-1]
	// fmt.Println(last)
	// fmt.Printf("exampleSlice: %v\n", exampleSlice)
	// fmt.Printf("Binary: %b\\%b", 4, 5)

	//bez powtórzeń
	exampleSlice2 := []int{1, 2, 3, 3, 4, 5, 3}
	var exampleSlice2Copy []int
	var slice []int
	fmt.Println(exampleSlice2Copy)
	for i := 0; i < len(exampleSlice2); i++ {
		acc := 0
		for j := 0; j < len(exampleSlice2Copy); j++ {
			if exampleSlice2[i] != exampleSlice2Copy[j] {
				acc++
			}
		}
		if acc == len(exampleSlice2Copy) {
			exampleSlice2Copy = append(exampleSlice2Copy, exampleSlice2[i])
			slice = append(slice, exampleSlice2[i])
		}
		acc = 0
	}
	fmt.Println(exampleSlice2Copy)
	fmt.Println(slice)

	// bez powtórzeń obok siebie
	// exampleSlice3 := []int{1, 2, 3, 3, 4, 5, 3}
	// var exampleSlice33 []int
	// var slice3 []int
	// for i := 0; i < len(exampleSlice); i++ {

	// }
}
