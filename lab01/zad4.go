//Napisz program, który zagra w zgadywanie z użytkownikiem. Należy wylosować liczbę którą następnie użytkownik ma odgadnąć. Po wprowadzeniu liczby program ma wypisać informację czy liczba jest prawidłowa. Jeżeli nie, to czy jest większa czy mniejsza od wylosowanej. Gra kończy się gdy użytkownik zgadnie liczbę lub przerwie program.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	random_int := rand.Intn(1000)
	fmt.Println(random_int)

	for {
		var a int
		fmt.Println("Zgadnij liczbę: ")
		fmt.Scan(&a)

		if a < random_int {
			fmt.Println("liczba zgadywana jest większa")
		} else if a > random_int {
			fmt.Println("liczba zgadywana jest mniejsza")
		}
		if a == random_int {
			fmt.Println("brawo zgadłeś :)")
			break
		}
	}
}
