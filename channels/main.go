package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
	}()
	//fmt.Println(messages)

	msg := <-messages
	fmt.Println(msg)
}
