package main

import (
	//"bufio"
	"fmt"
	// "os"
	// "strconv"
	// "strings"
)

func main() {

	for {
		fmt.Println("Calculator in Go Lang")
		num1 := 0
		fmt.Println("enter number operand")
		fmt.Scan(&num1)

		fmt.Println("enter operator")
		var op string
		fmt.Scan(&op)
		if op != "+" && op != "-" && op != "*" && op != "/" {
			fmt.Println("invalid operator")
			return
		}

		num2 := 0
		fmt.Println("enter number operand")
		fmt.Scan(&num2)

		switch op {
		case "+":
			fmt.Println(num1 + num2)
		case "-":
			fmt.Println(num1 - num2)
		case "*":
			fmt.Println(num1 * num2)
		case "/":
			fmt.Println(num1 / num2)
		}
	}
}