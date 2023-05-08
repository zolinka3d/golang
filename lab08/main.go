package main

import (
	"fmt"
	"lab8/rect"
)

func main() {
	r := rect.Rect{10, 12}
	fmt.Println(r.GetArea())
	fmt.Println(r.GetCircuit())

}
