package main

import "fmt"

func simplefunction() {
	fmt.Println("this is function 1")
}

func add(a, b int) int { //return type int
	//a,b int is input and (int) is return type
	return a + b
}

func subs(a int, b int) int { //return type int
	//a,b int is input and (int) is return type
	return a - b
}

func mult(a int, b int) (result int) { //return type int
	//a,b int is input and (int) is return type
	result = a * b
	return
} // in this function if you dont mention what to return at last it will automatically
// takes result as retunr because we have said 1at while declaring function

func main() {
	fmt.Println("functions in go lang")
	simplefunction()
	a := add(3, 5)
	b := subs(4, 2)
	c := mult(8 ,8)
	fmt.Println("ans is ", a)
	fmt.Print("sub is ", b)
	fmt.Println("mult is :", c)
}
