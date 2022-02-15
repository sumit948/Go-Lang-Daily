package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in go")
	filecontent := "This is my file"

	myfile, err := os.Create("./mygofile.txt")

	if err != nil {
		panic(err)
	}

	filelength, err := io.WriteString(myfile, filecontent)

	if err != nil {
		panic(err)
	}
	fmt.Println(filelength)
	defer myfile.Close()
	readFile("./mygofile.txt")

}

//readfile
func readFile(myfilename string) {
	databyte, err := ioutil.ReadFile(myfilename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text data inside file is: \n", string(databyte))
}
