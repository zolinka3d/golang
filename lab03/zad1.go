package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//max()
	zad3()

}

func algo(a int, acc int) int {
	if a <= 0 {
		return 1
	} else if a == 1 {
		return acc + 1
	} else if a%2 == 0 {
		return algo(a/2, acc+1)
	} else {
		return algo(a*3+1, acc+1)
	}
}

func max() {
	max := 0
	number := 0
	sum := 0
	for i := 10; i < 101; i++ {
		if max < algo(i, 0) {
			max = algo(i, 0)
			number = i
		}
		sum += algo(i, 0)
	}
	fmt.Println("Maximum: ", max, " for number: ", number)
	fmt.Println("Summary of lenght: ", sum)

}

func zad2() {
	for i := 1; i < 10; i++ {
		max := 0
		number := 0
		for j := i * 1000; j < (i+1)*1000; j++ {
			if max < algo(j, 0) {
				max = algo(j, 0)
				number = j
			}
		}
		fmt.Println("For range: ", i*1000, " - ", (i+1)*1000, " the max lenght of loop is: ", max, " and its for number: ", number)
	}
}

func zad3() {
	f, err := os.Create("dane.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for i := 1; i < 100_000; i++ {
		number := algo(i, 0)
		str := strconv.Itoa(i) + " " + strconv.Itoa(number) + "\n"
		_, err2 := f.WriteString(str)
		if err2 != nil {
			panic(err2)
		}
	}

	fmt.Println(f.Name())

}

func zad4() {

}
