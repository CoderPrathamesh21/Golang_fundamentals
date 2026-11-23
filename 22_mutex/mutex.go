package main

import (
	"fmt"
	"sync"
)

type post struct {
	mu    sync.Mutex
	views int
}

func (p *post) increment(wg *sync.WaitGroup) {
	defer func() {
		p.mu.Unlock()
		wg.Done()
	}()

	p.mu.Lock()
	p.views += 1

}

func main() {

	var wg sync.WaitGroup
	myPost := post{views: 0}

	for range 100 {
		wg.Add(1)
		myPost.increment(&wg)
	}

	// myPost.increment()
	// myPost.increment()

	wg.Wait()
	fmt.Println(myPost.views)
}
