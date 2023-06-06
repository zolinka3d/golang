package main

import (
	"fmt"
	"os"
	"strconv"
)

var RANGE int = 1000
var COLUMN int = 0
var FORMAT string = "%12.4f%12.4f\n"

func main() {
	AnalyzeOPT()

}

func AnalyzeOPT() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Println("opcja:", i, os.Args[i])
		opt := os.Args[i]

		if len(opt) < 2 {
			continue
		}

		switch opt[:2] {
		case "-f":
			if len(opt) > 2 { // -fzxczxczxc
				FORMAT = opt[2:] // obcinam -f, zostaje zxczxcz
				continue
			} else {
				if i+1 < len(os.Args) {
					FORMAT = os.Args[i+1]
					i++
					continue
				} else {
					fmt.Println("[w] opcja -f MUSI mieć argument")
				}
			}
		case "-c":
			if len(opt) > 2 {
				COLUMN = ustawopt(opt[2:], COLUMN, "błąd w opcji -c")
			} else {
				if i+1 < len(os.Args) {
					COLUMN = ustawopt(os.Args[i+1], COLUMN, "błąd w opcji -c")
				} else {
					fmt.Println("[w] opcja -c wymaga argumentu")
				}
			}
		case "-r":
			if len(opt) > 2 {
				RANGE = ustawopt(opt[2:], RANGE, "błąd w opcji -r")
			} else {
				if i+1 < len(os.Args) {
					RANGE = ustawopt(os.Args[i+1], RANGE, "błąd po opcji -r")
				} else {
					fmt.Println("[w] opcja -r wymaga argumentu")
				}
			}
		default:
		}
	}
	fmt.Println("RANGE:", RANGE)
	fmt.Println("COLUMN:", COLUMN)
	fmt.Println("FORMAT:", FORMAT)
}

func ustawopt(opt string, def int, e string) int {
	if v, err := strconv.Atoi(opt); err == nil {
		return v
	} else {
		fmt.Println("[w]", e, err)
		return def
	}
}
