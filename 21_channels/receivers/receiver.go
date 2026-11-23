package main

import "fmt"

// receive
func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult
}

// goroutine synchronizer
func task(done chan bool) {
	defer func() {
		done <- true
	}()

	fmt.Println("processing...")
}

func main() {

	result := make(chan int)

	go sum(result, 21, 12)
	res := <-result
	fmt.Println(res)

	done := make(chan bool)
	go task(done)
	<-done

}
