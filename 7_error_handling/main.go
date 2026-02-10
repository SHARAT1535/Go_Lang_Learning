package main

import "fmt"

func divison(a, b float32)(float32, error) {
	if b == 0 {
		return 0,fmt.Errorf("number not divisible by zero")
	}
	return a / b,nil
}

func main() {
	//fmt.Println("error handling in go lang ")
	ans ,err := divison(10, 2)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(ans)
}
