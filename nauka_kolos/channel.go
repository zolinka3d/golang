package main

import (
	"fmt"
	"time"
)

func printMessage(channel chan string) {
	for msg := range channel {
		fmt.Println(msg)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	messageChannel := make(chan string)

	go printMessage(messageChannel)

	messageChannel <- "Hello"
	messageChannel <- "Golang"
	messageChannel <- "World"

	close(messageChannel)
	time.Sleep(3 * time.Second)

}
