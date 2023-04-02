package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// można inicjalizować zmienne na dwa sposoby
	// np można podawać typ
	var a1 int = 5
	// albo można zostawić deklarowanie typu compilatorowi
	a2 := 5
	fmt.Print(a1, a2)

	// można inicjalizować tablice - wtedy trzeba określić ich dokładny rozmiar
	// tutaj będą same zera jeśli nic nie prypisaliśmy
	var array [10]int
	// a tu już nie
	var array1 = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// i przypisanie elementu do tablicy
	array[0] = 10
	fmt.Print(array, array1)

	//automatyczne wylicznenie długości tablicy
	array := [...]int{1,2,3,4,5,6,7}

	// alternatywą do tablic są slicy,
	//nie trzeba podawać wielkości przy deklaracji
	exampleSlice := []int{1, 2, 3}
	//można dodawać do niego wartości
	exampleSlice = append(exampleSlice, 11)
	// można je kopiować ale trzeba na to bardzo uważać, bo kopiowanie sprawia, że
	// slicy traktują się jak aliasy. Ostatnia operacja najważniejsza
	exampleSliceCopy := make([]int, len(exampleSlice))
	copy(exampleSliceCopy, exampleSlice)

	// sprawdznie zapasu długości
	fmt.Print(cap(exampleSlice))

	//a to jak chcemy tylko kawałek tablicy odczytać, jest to też wycinek
	fmt.Print(array[0:5])

	// co dla mnei specyficzne w tym języku jest to brak nawiasów przy ifach i pętlach
	if array == array1 {
		fmt.Print("true")
	}

	// mamy różne sposoby przechodzenia po tablicach
	for i := 0; i < len(array); i++ {
		fmt.Print(array[i])
	}
	for i := range array {
		fmt.Print(i)
	}
	for i, v := range array {
		fmt.Print(i, v)
	}
	i := 0
	for range array {
		fmt.Print(array[i])
		i++
	}
	for _, v := range array {
		fmt.Print(v)
	}

	// ciekawsze typy języka
	var a uint8 = 255   //(0,...,255)
	b := complex(10, 1) // liczby zespolone
	var c rune = 'R'
	fmt.Print(a, b, c)
	// w zależnośći od architektury int, uint, intptr 16,32,64 bity

	// Dla przejrzystości można zapisać długie liczby z kreską
	// nic to nie zmienia w liczbie
	longie := 1_000_000_000
	fmt.Print(longie)

	//poznaliśmy też pare ciekawych operatorów
	// & - przypisanie pamięcie
	// * - wskaźnik
	// ^ - XOR
	// <- - wczytanie z kanału do kanału
	// << >> - przesunięcia bitowe

	// Konwencją nazewniczą camelCase

	// Nazwy interfaców piszemy dodając -er
	// Writer, Reader, Sender

	// przykłady konwersji
	liczba := 3.14
	tekst := fmt.Sprintf("%8.3f", liczba)
	fmt.Println(tekst) // output to "   3.140"

	tekst2 := "0.314e+01"
	liczba2, error := strconv.ParseFloat(tekst2, 64)
	fmt.Print(liczba2)
	//trzeba łapać te błędy!
	if error != nil {
		log.Fatal(error)
	}

	// Mamy coś takiego jak Mapy, trochę jak słowniki, tworzenie księgi z przypisanymi do siebie wartościami :)
	var lista map[int]string
	lista[5] = "pięc"
	lista[3] = "trzy"

	// Bardzo przydatne są struct - tworzenie swojej klasy
	// definiowanie struktury
	type Student struct {
		firstName string
		lastName  string
		age       int
	}
	// tworzenie instacji
	studentka1 := Student{firstName: "Zosia", lastName: "Tokarczyk", age: 21}
	fmt.Print(studentka1)

	// można dodać do strutury metodę
	func (s Student) getAge() int {
		return s.age
	}

	// można też szybciej zdefiniować structurę po przecinku, jeśli typy się zgadzają
	type Teacher struct {
		firstName, lastName, hobby string
	}

	// można zagnieżdżać w sobie struktury
	type Job struct {
		position string
		ageExperience int
	}
	type Man struct {
		firstName string,
		job Job
	}

	// Ja używałam struktur w zadaniu na zgadywanie liczb, żeby utworzyć tabelę wyników
	// Potrzebowałam wtedy uzyć czegoś w stylu mapy/krotki i znalazłam ten sposób
	type wynik struct {
		points int
		name   string
	}

	// odczyt pliku
	tresc, err := os.ReaderFile("plik.txt")
	if ee ! nil {
		panic(err)
	}

	// odczyt http 
	resp, err1 := http.Get("...")
	defer resp.Body.Close()
	body, err2 := ioutil.ReadeAll(resp.Body)
}

// Ważne przy definiowaniach funkcji ich typowanie, jeśli nie ma żadnego funckja to void i nic nie zwraca
//wynik funkcji musi się zgadzać z typem
func add (a int, b int) {
	return a + b
}


// Golang jest bardzo przejrzystym językiem. Nie ma za dużo skomplikowanych konstrukcji.
// W większości używamy zwykłego fora. Jest silnie typowany, a jednak bardzo przyjemny
// Mam kontrowersyjną opinię, że to powinien być język którego byśmy się uczyli na drugim semestrze zamiast Scali
