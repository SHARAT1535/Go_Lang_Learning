package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["java"] = "java is programming lang"
	languages["python"] = "python is pro lang"
	languages["c++"] = "c++is pro lang"
	//fmt.Printf("list of all programing lang are:\n", languages)
	//fmt.Println("java shorts for:", languages["java"])
	//	/delete(languages, "java")
	grades, exists := languages["java"]
	if !exists {
		fmt.Println("java exits and its value is:", grades)
	} else {
		fmt.Println("java exits and its value is:", grades)
	}
	//fmt.Println("grades areṇ:",grades)

	for index, val := range languages {
		fmt.Printf("key is %s and value is %s\n", index, val)
	}

	person := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	for index, keyval := range person {
		fmt.Printf("key is %s and value is %d\n", index, keyval)
	}
}
