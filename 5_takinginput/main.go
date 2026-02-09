package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//fmt.Println("Enter your name")
	var name string
	//fmt.Scan(&name)// input as raju hire

	fmt.Println("enter your name  ", name) // when you print this ouput would be only raju i .e it will print till  black space
	reader := bufio.NewReader(os.Stdin)
	name, _ = reader.ReadString('\n')
	fmt.Println("your name is ", name)

}
