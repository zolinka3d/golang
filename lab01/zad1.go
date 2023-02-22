// Napisz prosty program, który pyta ile masz lat, a następnie wypisuje liczbę miesięcy i dni które masz już za sobą.

package main

import "fmt"

func main() {
	var i int

	fmt.Print("Type your age: ")
	fmt.Scan(&i)
	fmt.Println("Your age is:", i)
	fmt.Println("You have lived:", i*12, "months")
}
