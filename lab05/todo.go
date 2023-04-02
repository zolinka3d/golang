package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fileContent := readTodo("todo.txt")
	// fmt.Print(fileContent)
	printLines(fileContent)
}

func readTodo(nameOfFile string) string {
	dat, err := os.ReadFile(nameOfFile)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	if len(dat) == 0 {
		log.Fatal("nothing in the file")
	}

	return string(dat)
}

func printLines(fileContent string) {
	arrayOfLines := strings.Split(fileContent, "\n")

	for _, el := range arrayOfLines {
		if len(el) < 3 {
			continue
		} else if el[0] == '#' {
			fmt.Println("\u001b[30m", el, "\u001b[0m")
		} else if el[0] == '[' && el[2] == ']' {
			switch el[1] {
			case 'x':
				fmt.Println("\u001b[32m", el, "\u001b[0m")
			case '-':
				fmt.Println("\u001b[31m", el, "\u001b[0m")
			case '+':
				fmt.Println("\u001b[33m", el, "\u001b[0m")
			}
		}

	}
}
