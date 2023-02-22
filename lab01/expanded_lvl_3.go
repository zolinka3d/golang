// Poziom 3. Napisz program taki jak na poziomie 2, ale dodaj do niego dodatkowe pytanie na końcu, po odgadnięciu liczby, które brzmi "Czy gramy jeszcze raz? [T/N]". Gdy użytkownik wybierze odpowiedź (domyślnie tak lub nie), program rozpoczyna jeszcze raz grę z INNĄ wylosowaną liczbą, albo się kończy.
package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	fmt.Println("Żeby zakończyć pracę programu napisz - koniec")
out:
	for {
		random_int := rand.Intn(1000)
		fmt.Println(random_int)
		fmt.Println("Zaczynamy zabawę :)")

		for {
			var a string
			fmt.Println("Zgadnij liczbę: ")
			fmt.Scan(&a)

			if a == "koniec" {
				break out
			} else {
				aInt, err := strconv.Atoi(a)
				if aInt < random_int {
					fmt.Println("za mała")
				} else if aInt > random_int {
					fmt.Println("za duża")
				}
				if aInt == random_int {
					fmt.Println("gratulację :)")
					var answer string
					fmt.Println("Gramy jeszcze raz? [T/N]")
					fmt.Scan(&answer)
					if answer == "T" {
						break
					} else {
						break out
					}
				}

				if err != nil {
					fmt.Println("Ni to poprawna liczba, ni to koniec, spadówa")
					return
				}
			}

		}
	}
	fmt.Println("Żegnaj")
}
