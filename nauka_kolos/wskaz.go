// package main

// import "fmt"

// func main() {
// 	// var a *int = new(int)
// 	// fmt.Println("Value of a:", *a)  // Wyświetli: "Value of a: 0"
// 	// fmt.Println("Address of a:", a) // Wyświetli adres pamięci zmiennej a

// 	// *a = 42
// 	// fmt.Println("Value of a after modification:", *a) // Wyświetli: "Value of a after modification: 42"

// 	/////////

// 	// var zmienna int = 42
// 	// fmt.Println("Value of zmienna:", zmienna)    // Wyświetli: "Value of zmienna: 42"
// 	// fmt.Println("Address of zmienna:", &zmienna) // Wyświetli adres pamięci zmiennej zmienna
// 	// var wskaznik *int = nil
// 	// wskaznik = &zmienna
// 	// fmt.Println("Value of wskaznik:", wskaznik)   // Wyświetli adres pamięci zmiennej zmienna
// 	// fmt.Println("Value of *wskaznik:", *wskaznik) // Wyświetli: "Value of *wskaznik: 42"
// 	// *wskaznik = 777
// 	// fmt.Println("Value of zmienna after modification:", zmienna) // Wyświetli: "Value of zmienna after modification: 777"

// }
package main

import (
	"fmt"
)

type Sizer interface {
	Size() int
}

type MyString string
type MyInt int

func (s MyString) Size() int {
	return len(s)
}

func (i MyInt) Size() int {
	return int(i)
}

func PrintSize(s Sizer) {
	fmt.Println("Size:", s.Size())
}

func main() {
	ms := MyString("Hello, Golang!")
	mi := MyInt(42)

	PrintSize(ms) // Wyświetli: "Size: 14"
	PrintSize(mi) // Wyświetli: "Size: 42"
}
