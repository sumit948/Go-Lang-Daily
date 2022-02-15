package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study in go lang")

	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 01:5pm Monday")) //giving format like days

	//create date

	createDate := time.Date(2020, time.August, 10, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(presentTime.Format("01/02/2006 1:15pm Monday"))
}
