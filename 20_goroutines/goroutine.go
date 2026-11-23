package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Doing task", id)
}

func main() {

	var waitGroup sync.WaitGroup

	for i := 0; i <= 10; i++ {
		waitGroup.Add(1)
		go task(i, &waitGroup)

		// go func(i int) {
		// 	fmt.Println(i)
		// }(i)
	}

	//time.Sleep(time.Second * 2)
	waitGroup.Wait()
}
