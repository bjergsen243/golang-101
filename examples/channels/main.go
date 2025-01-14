package main

import "fmt"

func main() {
	messages := make(chan string)

	// send a value into a channel using the channel <- syntax
	go func() {
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println("message is:", msg)
}
