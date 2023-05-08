package main

import (
	"flag"
	"fmt"
	"math"
)

func main() {
	a := flag.Float64("a", 0, "a")
	b := flag.Float64("b", 0, "b")
	c := flag.Float64("c", 0, "c")
	flag.Parse()
	fmt.Print(t(*a, *b, *c))
}

type T struct {
	x, y float64
}

type W struct {
	p T
	c bool
}

func t(a, b, c float64) W {
	d := b*b - 4*a*c
	if d < 0 || a == 0 {
		return W{T{0, 0}, false}
	}
	return W{T{(-b - math.Sqrt(d)) / 2 * a, (-b + math.Sqrt(d)) / 2 * a}, true}

}
