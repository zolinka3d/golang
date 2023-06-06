package forest

import (
	"fmt"
	"math/rand"
)

func FireOneForest(percent int, x int, y int) float64 {

	// fmt.Println("Hello, playground")
	pole := pole(x, y)
	fullForestRandom(pole, percent)
	// printForest(pole)
	//🌳

	coords := drawTree(len(pole), len(pole[0]))
	// fmt.Println(coords)
	fireTree(pole, coords.x, coords.y)
	// printForest(pole)
	// fmt.Println("-----------------")
	fireForest(pole)
	// printForest(pole)
	return calculateFiredForest(pole)
}

func pole(x int, y int) [][]rune {

	pole := make([][]rune, x)
	for i := range pole {
		pole[i] = make([]rune, y)
	}

	return pole
}

type Coords struct {
	x int
	y int
}

func drawTree(x int, y int) Coords {
	return Coords{rand.Intn(x), rand.Intn(y)}
}

func printForest(forest [][]rune) {
	for _, el := range forest {
		for _, el2 := range el {
			if el2 == 0 {
				fmt.Print("  ")
			}
			fmt.Printf("%c", el2)

		}
		fmt.Println()
	}

}

func fullForestRandom(forest [][]rune, percent int) {

	for i := range forest {
		for j := range forest[i] {
			if rand.Intn(100) < percent {
				forest[i][j] = '🌳'
			}
		}
	}
}

func fireTree(forest [][]rune, x int, y int) {
	if forest[x][y] == '🌳' {
		forest[x][y] = '🔥'
	}
}

func fireForest(forest [][]rune) bool {
	fired := false

	for i := range forest {
		for j := range forest[i] {
			if forest[i][j] == '🔥' {
				if i > 0 && forest[i-1][j] == '🌳' {
					forest[i-1][j] = '🔥'
					fired = true
				}
				if i < len(forest)-1 && forest[i+1][j] == '🌳' {
					forest[i+1][j] = '🔥'
					fired = true
				}
				if j > 0 && forest[i][j-1] == '🌳' {
					forest[i][j-1] = '🔥'
					fired = true
				}
				if j < len(forest[i])-1 && forest[i][j+1] == '🌳' {
					forest[i][j+1] = '🔥'
					fired = true
				}

			}
		}
	}
	if fired {
		return fireForest(forest)
	}

	return fired

}

func calculateFiredForest(forest [][]rune) float64 {
	fired := 0.0
	survived := 0.0

	for i := range forest {
		for j := range forest[i] {
			if forest[i][j] == '🔥' {
				fired++
			} else if forest[i][j] == '🌳' {
				survived++
			}
		}
	}
	return fired / (fired + survived) * 100
}
