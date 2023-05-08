package main

import (
	"fmt"
)

func main() {

	fmt.Println("coś1")
	go func() {
		fmt.Println("2")
	}()
	// time.Sleep(1 * time.Second)
	fmt.Println("coś3")

}
