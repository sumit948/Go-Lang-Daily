package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://www.youtube.com/watch?v=cl7_ouTMFh0&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=26"

func main() {
	fmt.Println("Welcome to urls")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	fmt.Println(result.Host) //youtube
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	fmt.Println(result.Path)

	queryparam := result.Query()
	fmt.Printf("Url query param type: %T", queryparam)

}
