package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup //pointer
var mut sync.Mutex    //pointer

func main() {
	fmt.Println("Welcome to check status of websites")

	websiteslist := []string{
		"google.com",
		"github.com",
	}
	for _, res := range websiteslist {
		go getstatuscode(res)
		wg.Add(1)
	}
	wg.Wait()
}
func getstatuscode(endpoint string) {
	defer wg.Done()
	result, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("Not getting status code....")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", result.StatusCode, endpoint)
	}
}
