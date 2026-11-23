package main

import (
	"fmt"
	"math/rand"
	"time"
)

// send
func processNumber(numberChannel chan int) {

	for num := range numberChannel {
		fmt.Println("Processing number", num)
		time.Sleep(time.Second * 1)
	}
}

func main() {

	// messageChannel := make(chan string)

	// messageChannel <- "ping" //blocking operation
	// messsage := <-messageChannel

	// fmt.Println(messsage)

	numberChannel := make(chan int)

	go processNumber(numberChannel)
	//numberChannel <- 5

	for {
		numberChannel <- rand.Intn(100)
	}
	//time.Sleep(time.Second * 2)

}
