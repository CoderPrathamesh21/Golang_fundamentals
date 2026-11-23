package main

import (
	"fmt"
	"time"
)

func emailSender(emailChannel chan string, done chan bool) {

	defer func() {
		done <- true
	}()

	for email := range emailChannel {
		fmt.Println("Sending emails to", email)
		time.Sleep(time.Second)
	}

}

func main() {

	emailChannel := make(chan string, 100)
	done := make(chan bool)

	go emailSender(emailChannel, done)

	for i := range 10 {
		emailChannel <- fmt.Sprintf("%d@gmail.com", i)
	}

	fmt.Println("Sent!!")
	close(emailChannel)
	<-done
	// emailChannel <- "1@example.com"
	// emailChannel <- "2@example.com"

	// fmt.Println(<-emailChannel)
	// fmt.Println(<-emailChannel)
}
