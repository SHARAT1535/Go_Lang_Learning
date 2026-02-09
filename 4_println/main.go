package main

import "fmt"

func main() {
	fmt.Println("test")
	age := 10
	name := "RAJ"
	height := 1.2 //default float64
	fmt.Println(age, name, height)
	fmt.Println("age", age, "height:", height, "name:", name)
	fmt.Printf("age is %d\n", age)
	fmt.Printf("heaght is %.6f\n", height)
	fmt.Printf("Type of name variable is %T\n", name)
	fmt.Printf("Type of age variable is %T\n", age)
	fmt.Printf("types of height age variable is %T", height)
}
