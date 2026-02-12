package main

import "fmt"

//pointer holds the address of another variable in memory

func modifybal(a *int) {
	*a = *a + 10
}

func main() {

	var a int = 10
	ptr := &a

	fmt.Println("VALUE OF A IS ", a)
	fmt.Println("val of a thorugh pointer is ", *ptr)

	a = 20
	modifybal(&a)
	fmt.Println("VALUE OF A IS ", a)
}
