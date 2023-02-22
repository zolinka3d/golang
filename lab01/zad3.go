// Napisz program, który wykona obliczenie pierwiastków trójmianu kwadratowego. Dla przypadku gdy delta jest ujemna program może wypisać komunikat, że nie ma pierwiastków (lub w trudniejszej wersji - może policzyć pierwiastki zespolone).

package main

import "fmt"

func main() {
	var a float64
	var b float64
	var c float64
	fmt.Println("Podaj parametry a*x^2 + b*x + c = 0")
	fmt.Print("a = ")
	fmt.Scan(&a)
	fmt.Print("b = ")
	fmt.Scan(&b)
	fmt.Print("c = ")
	fmt.Scan(&c)
	fmt.Println(a, "x^2 +", b, "x +", c, "= 0")
	var delta float64 = float64(b*b - 4*a*c)
	fmt.Println("Delta =", delta)
	if delta >= 0 {
		fmt.Println("x1 =", (-b+delta)/(2*a))
		fmt.Println("x2 =", (-b-delta)/(2*a))
	} else {
		fmt.Println("Nie ma pierwiastków")
	}
}
