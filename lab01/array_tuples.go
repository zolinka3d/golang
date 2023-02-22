package main

import (
	"fmt"
)

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [10]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	var s []int = primes[1:4] // slice is just a reference
	fmt.Println(s)
	primes[1] = 1
	fmt.Println(s)

	tablica_struct := []struct {
		i int
		b string
	}{
		{2, "pipi"},
		{3, "popo"},
	}
	fmt.Println(tablica_struct)

	fmt.Println("dlugość", len(primes))
}
