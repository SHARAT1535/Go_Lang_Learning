package main

import (
	"errors"
	"fmt"
	"os"
)

type agerror struct {
	age int
}

func (e *agerror) Error() string {
	return fmt.Sprintf("age %d id below 18", e.age)

}
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divdide by zero")
	}
	return a / b, nil
}
func checkage(Age int) error {
	if Age < 18 {
		return &agerror{age: Age}
	}
	return nil
}
func openfile(filename string) error {
	_, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file:%w", err)
	}
	return nil
}

func panixexample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("revocered from panic", r)
		}
	}()
	panic("somthing went wrong terror")
}

func main() {
	fmt.Println("--------Dividon example")
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("result:", result)
	}

	fmt.Println("--------Age check example")
	err = checkage(26)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("\n--------file open example-------")
	err = openfile("test.txt")
	if err != nil {
		fmt.Println("file does not exist:", err)
	}

	fmt.Println("\n----- Panic & Recover Example -----")
	panixexample()

	fmt.Println("\nProgram continues after recovered")

}
