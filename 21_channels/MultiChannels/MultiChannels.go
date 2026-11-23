package main

import "fmt"

func main() {

	channel1 := make(chan int)
	channel2 := make(chan string)

	go func() {
		channel1 <- 10
	}()

	go func() {
		channel2 <- "Ten"
	}()

	for i := 0; i < 2; i++ {
		select {
		case channel1Val := <-channel1:
			fmt.Println("Received data from channel1", channel1Val)
		case channel2Val := <-channel2:
			fmt.Println("Received data from channel2", channel2Val)
		}
	}
}
