// Poziom 4. Program z poziomu 3 umożliwia wiele gier. Zmodyfikuj go tak, aby umożliwiał zapamiętanie w ilu próbach użytkownik odgadł liczbę. Po zakończeniu gry zapytaj użytkownika o imię, i zapisz je do struktury danych. Gdyby gra była kontynuowana kolejny raz, ponawiaj takie pytania, aby zebrać dane o wynikach różnych użytkowników. Na zakończenie programu wypisz podsumowanie: kto w ilu próbach odgadywał liczby, posortowane rosnąco według liczby prób.

package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("Żeby zakończyć pracę programu napisz - koniec")
	type wynik struct {
		points int
		name   string
	}
	var tab [10]wynik
out:
	for i := 0; i < 10; i++ {
		var acc int
		random_int := rand.Intn(1000)
		fmt.Println(random_int)
		fmt.Println("Zaczynamy zabawę :)")

		for {
			acc++
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
					var name string
					fmt.Println("gratulację :) jakie jest twoje imie?")
					fmt.Scan(&name)
					tab[i] = wynik{points: acc, name: name}
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
	scores := []wynik{}
	//fmt.Println(tab)
	for i := range tab {
		if tab[i].points > 0 {
			scores = append(scores, tab[i])
		}
	}

	fmt.Println("Oto tabela wyników")
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].points < scores[j].points
	})
	fmt.Println(scores)
}
