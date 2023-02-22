package main

import (
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	d1 := []byte("hello\ngo\n3")
	// f, err := os.Create("./cos")
	err := os.WriteFile("./cos", d1, 0644)
	check(err)
	//defer f.Close()
}
