package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("welcome to mutex")
	var score = []int{0}

	var wg = &sync.WaitGroup{}
	var mut = &sync.Mutex{}
	wg.Add(3) //instead of seprately adding we can do direct numb of func
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("route one")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	//wg.Add(2)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("route two")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("route three")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println("Yes printing...", score)
}
