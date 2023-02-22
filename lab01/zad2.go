//Napisz program, który liczy ile lat miałby użytkownik, gdyby mieszkał na Marsie, na Wenus lub na innych planetach. Ile Ty możesz mieć aktualnie lat na takich planetach? Wiek mieszkańca danej planety to liczba pełnych obrotów planety wokół słońca. Sprawdź w internecie...

package main

import "fmt"

func main() {
	var i float64
	var mars_age_length float64 = 686.9601 / 365
	var venus_age_length float64 = 224.701 / 365
	var mercury_age_length float64 = 87.969 / 365

	fmt.Println("Mars age length:", mars_age_length)
	fmt.Println("Venus age length:", venus_age_length)
	fmt.Println("Maercury age length:", mercury_age_length)

	fmt.Print("Type your age: ")
	fmt.Scan(&i)
	fmt.Println("Your age is:", i*mars_age_length, "on Mars")
	fmt.Println("Your age is:", i*venus_age_length, "on Venus")
	fmt.Println("Your age is:", i*mercury_age_length, "on Mercury")
}
