//Napisz program, który zagra w zgadywanie z użytkownikiem. Należy wylosować liczbę którą następnie użytkownik ma odgadnąć. Po wprowadzeniu liczby program ma wypisać informację czy liczba jest prawidłowa. Jeżeli nie, to czy jest większa czy mniejsza od wylosowanej. Gra kończy się gdy użytkownik zgadnie liczbę lub przerwie program.

package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	var a string
	fmt.Println("Napisz liczbę: ")
	fmt.Scan(&a)

	fmt.Println("Type of variable1:", reflect.TypeOf(a))

	aNew, err := strconv.Atoi(a)

	if err != nil {
		fmt.Println("Error during conversion")
		return
	}
	fmt.Println("Type of variable2:", reflect.TypeOf(aNew))
}
