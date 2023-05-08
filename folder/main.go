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
	messageChannel := make(chan string) // mogę tutaj dodać drugi paramentr czyli ilość elementów jakie mogę dodać do kanału
	// FIFO
	go printMessage(messageChannel)

	messageChannel <- "Hello"
	messageChannel <- "Golang"
	messageChannel <- "World"

	// fmt.Print(<-messageChannel) // nmg tego tutaj zrobić
	close(messageChannel)

	// fmt.Print(<-messageChannel) // tutaj mogę
	time.Sleep(3 * time.Second)

}
