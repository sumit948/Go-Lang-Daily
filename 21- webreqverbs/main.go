package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	PerformingGetReq()
	performingPostCall()
	PerformPostFormRequest()
}

func PerformingGetReq() {
	myurl := "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(nil)
	}

	defer response.Body.Close()

	fmt.Println("Get Status code: ", response.StatusCode)
	fmt.Println("Content length: ", response.ContentLength)

	var responseString strings.Builder
	cantent, _ := ioutil.ReadAll(response.Body)
	byteCount, _ := responseString.Write(cantent)
	fmt.Println("Byte Count:  ", byteCount)
	fmt.Println(responseString.String())
	fmt.Println("get content: ", string(cantent))
}

//post data
func performingPostCall() {
	myurl2 := "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"employeeName":"Jhon Weak2",
			"Department":"RPC",
			"Contact":"21453652",
			"isAvailable":"true"
		}
	`)
	response, err := http.Post(myurl2, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content2, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content2))

}

//---------------------------------------------------------------------------
//performing postForm request

func PerformPostFormRequest() {
	myurl3 := "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("FistName", "Jhone Week")
	data.Add("Country", "India")
	data.Add("Contact", "416032")

	response3, err := http.PostForm(myurl3, data)
	if err != nil {
		panic(err)
	}
	defer response3.Body.Close()

	content3, _ := ioutil.ReadAll(response3.Body)
	fmt.Println(string(content3))
}
