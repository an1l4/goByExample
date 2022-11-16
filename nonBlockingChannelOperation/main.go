package main

import "fmt"

func main() {
	message := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-message:
		fmt.Println("message received", msg)
	default:
		fmt.Println("message not received")
	}

	msg := "Hi"

	select {
	case message <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("message not sent")
	}

	select {
	case msg := <-message:
		fmt.Println("received message", msg)
	case sgl := <-signals:
		fmt.Println("signal received", sgl)
	default:
		fmt.Println("no activity")
	}

}
