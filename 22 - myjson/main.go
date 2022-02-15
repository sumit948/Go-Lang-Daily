package main

import (
	"encoding/json"
	"fmt"
)

type Course struct { //it wil reflect to json names
	Name       string `json:"CourseName"`
	Department string
	Salary     int
	Password   string   `json:"-"`              //in json it will remove pass field
	Tags       []string `json:"tags,omitempty"` //omitempty remove null value
}

func main() {
	fmt.Println("Welocome to json")
	//EncodeJson()
	DecodeJson()
}
func EncodeJson() {
	lcocourse := []Course{
		{"Sumit", "EAC", 2500, "abc123", []string{"Js", "Go", "Java"}},
		{"Jhon", "RAC", 2600, "bcd123", []string{"Js", "Go", "Java"}},
		{"Satya", "BAC", 2700, "khj123", nil},
	}
	//packge data to json data

	// resultJson, err := json.Marshal(lcocourse) //not better not good loocking
	resultJson, err := json.MarshalIndent(lcocourse, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", resultJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"CourseName":"Sumit",
		"Department":"EAC",
		"Salary": 2500,
		"tags": [
					"JS",
					"GO",
					"Java"
				]
	}
	`)

	var getRefCourse Course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json Data is Valid")
		json.Unmarshal(jsonDataFromWeb, &getRefCourse)
		fmt.Printf("%#v\n", getRefCourse)
	} else {
		fmt.Println("Json Data is not valid")
	}

	//som case where you wnat to add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("keys: %v and values: %v type is: %T\n", k, v, v)
	}
}
