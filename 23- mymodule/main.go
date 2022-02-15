package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to module")
	greater()
	r := mux.NewRouter()
	r.HandleFunc("/", servHome).Methods("Get")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greater() {
	fmt.Println("Hey There is MOD")
}

func servHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>THis is from newly created Module</h1>"))
}
