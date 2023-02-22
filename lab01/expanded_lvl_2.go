// Poziom 2. Napisz program jak wyżej, ale zmodyfikuj go w taki sposób, aby możliwe było podanie odpowiedzi "koniec", po której pytania o liczbę są przerywane a program natychmiast się kończy pisząc "żegnaj". Program powinien w powitaniu napisać informację, że po wpisaniu "koniec" nastąpi zakończenie działania. Poza tym - wszystko może pozostać takie, jak na poziomie pierwszym.
package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	random_int := rand.Intn(1000)
	fmt.Println(random_int)
	fmt.Println("Żeby zakończyć pracę programu napisz - koniec")

	for {
		var a string
		fmt.Println("Zgadnij liczbę: ")
		fmt.Scan(&a)

		if a == "koniec" {
			fmt.Println("Żegnaj")
			break
		} else {
			aInt, err := strconv.Atoi(a)
			if aInt < random_int {
				fmt.Println("za mała")
			} else if aInt > random_int {
				fmt.Println("za duża")
			}
			if aInt == random_int {
				fmt.Println("gratulację :)")
				break
			}

			if err != nil {
				fmt.Println("Error during conversion")
				return
			}
		}

	}
}
