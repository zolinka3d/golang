// czego się nauczyliśmy?
// udowodnij że nauczyłeś się tablic, wycinków

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// arrayInit()
	// zad3(3)
	// zad4(zad3(3))
	// zad5(matrix(3), matrix(3))
	matrix1 := matrix(3)
	matrix2 := matrix(3)
	fmt.Println(areSliceEqual(zad5(matrix1, matrix2), zad5(matrix2, matrix1)))
}

// mniej niż dwóch sąsiadów - trup
// jeśli jest puste pole, które ma trzech sąsiadów to pojawia się
// stablina koniczynka

func arrayInit() {
	var array1 [20]float64
	var array2 [20]float64
	for i := 0; i < 20; i++ {
		array1[i] = 2.0
	}

	for i := 0; i < 20; i++ {
		array2[i] = 3.0
	}

	var suma [20]float64
	for i := 0; i < 20; i++ {
		suma[i] = array1[i] + array2[i]
	}

	for _, el := range suma {
		fmt.Println(el)
	}
}

func zad2() {
	var wektor1 = [5]float64{1, 2, 3, 4, 5}
	var wektor2 = [5]float64{1, 2, 3, 4, 5}
	var iloczyn [5]float64

	for i := 0; i < 5; i++ {
		iloczyn[i] = wektor1[i] * wektor2[i]
		fmt.Println(iloczyn[i])
	}

}

func zad3(size int) [][]int {

	matrix := [][]int{}
	acc := 0
	for i := 0; i < size; i++ {
		var row []int
		for j := 0; j < size; j++ {
			row = append(row, acc)
			acc++
		}
		matrix = append(matrix, row)
	}
	fmt.Println(matrix)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Print("\n")
	}
	return matrix
}

func zad4(matrix [][]int) {
	inverse_matrix := [][]int{}
	size := len(matrix)
	for i := 0; i < size; i++ {
		row := matrix[i]
		for i2, j := 0, len(row)-1; i2 < j; i2, j = i2+1, j-1 {
			row[i2], row[j] = row[j], row[i2]
		}
		inverse_matrix = append(inverse_matrix, row)

	}
	fmt.Println(" ")
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Print(inverse_matrix[i][j], " ")
		}
		fmt.Print("\n")
	}

}

func matrix(size int) [][]int {
	matrix := [][]int{}
	for i := 0; i < size; i++ {
		var row []int
		for j := 0; j < size; j++ {
			row = append(row, rand.Intn(10))
		}
		matrix = append(matrix, row)
	}
	for i := 0; i < size; i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("")
	return matrix
}

// el[0][0] = matrix1[0][0] * matrix2[0][0] + matrix1[0][1] * matrix[1][0] + matrix1[0][2] * matrix2[2][0]
// el[0][1] = matrix1[0][0] * matrix2[0][1] + matrix1[0][1] * matrix2[1][1] + matrix1[0][2] * matrix2[2][1]
func zad5(matrix1 [][]int, matrix2 [][]int) [][]int {
	size1 := len(matrix1)
	size2 := len(matrix2[0])
	multiplied_matrix := [][]int{}
	// multiplied_matrix := make([][]int, size1)
	// for i := range multiplied_matrix {
	// 	multiplied_matrix[i] = make([]int, size2)
	// }
	for i := 0; i < size1; i++ {
		row := []int{}
		for j := 0; j < size2; j++ {
			el := 0
			for k := 0; k < size1; k++ {
				el += matrix1[i][k] * matrix2[k][j]
			}
			row = append(row, el)
		}
		multiplied_matrix = append(multiplied_matrix, row)

	}
	for i := 0; i < size1; i++ {
		fmt.Println(multiplied_matrix[i])
	}
	return multiplied_matrix
}

func areSliceEqual(matrix1 [][]int, matrix2 [][]int) bool {
	size1 := len(matrix1)
	size11 := len(matrix1[0])
	size2 := len(matrix2)
	size22 := len(matrix2[0])

	if size1 != size2 || size11 != size22 {
		return false
	}
	for i := 0; i < size1; i++ {
		for j := 0; j < size2; j++ {
			if matrix1[i][j] != matrix2[i][j] {
				return false
			}
		}
	}
	return false
}
