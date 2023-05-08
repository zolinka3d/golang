package main

import (
	"fmt"
)

func main() {
	slice := MySlice{Int(1), String("2")}
	fmt.Println(slice)
}

type IntOfString interface {
	isIntOfString()
}

type Int int

func (i Int) isIntOfString() {}

type String string

func (s String) isIntOfString() {}

type MySlice []IntOfString
