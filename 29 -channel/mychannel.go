package main

import (
	"fmt"
	"sync"
)

func main() {
	mychan := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		//receive data
		fmt.Println(<-mychan)
		wg.Done()
	}(mychan, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		//send data
		mychan <- 5
		mychan <- 7
		close(mychan)
		wg.Done()
	}(mychan, wg)

	wg.Wait()
}
