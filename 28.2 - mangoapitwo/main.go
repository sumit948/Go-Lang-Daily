package main

import (
	"fmt"
	"log"
	"net/http"

	"main.go/router"
)

func main() {
	fmt.Println("MONGO API")
	r := router.Router()
	fmt.Println("Server Started...")
	log.Fatal(http.ListenAndServe(":2000", r))
	fmt.Println("Port - 2000")
}
