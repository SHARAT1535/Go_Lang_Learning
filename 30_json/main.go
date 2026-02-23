package main

import (
	"encoding/json"
	
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Isadult bool   `json:"isadult"`
}

func main() {
	//fmt.Println("json learning")
	person := Person{Name: "Johin", Age: 34, Isadult: true}
	fmt.Println("\n", person)

	//convert person into json encoding
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error encoding json:", err)
	}
	fmt.Println("json data is ", string(jsonData))

	//decoding (unmarshal) json data into person struct
	var decodeperson Person
	//var err error
	err = json.Unmarshal(jsonData, &decodeperson)
	if err != nil {
		fmt.Println("error unmarshing ", err)
	}
	fmt.Println("decoded person is ", decodeperson)
}
