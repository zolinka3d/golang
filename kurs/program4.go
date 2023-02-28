package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var A int64 = 1
	var B int32
	C := 3

	fmt.Printf("Zmienna A jest typu %T\n", A)
	fmt.Printf("Zmienna B jest typu %T\n", B)
	fmt.Printf("Zmienna C jest typu %T\n", C)

	const (
		one = iota
		two
		three
		four
	)
	type MyTypeInt int64
	type MyTypeAlias = MyTypeInt

	var runNum rune // rozszerzone ascii
	fmt.Printf("Default runNum value %d, type %T\n\n", runNum, runNum)

	//stringi
	MyStri := "DevOpsiarz"
	//stringi w Go to nie do końca stringi (?)

	//Jeśli chcemy wydrukować coś
	for i := 0; i < len(MyStri); i++ {
		fmt.Printf(" => %x => %s \n", MyStri[i], string(MyStri[i]))
	}

	Emoji := "🤨"
	fmt.Println("Length of Emoji: ", len(Emoji))
	fmt.Println("Length of Emoji: ", utf8.RuneCountInString(Emoji))

	//stringi są niezmienne, nie ma czegos takiego jak: Emoji[2]='x'
}
