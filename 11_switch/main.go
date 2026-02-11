package main

import "fmt"

func main() {
	var day string = "Monday"
	switch {
	case day == "monday" || day == "Monday":
		fmt.Println("Today is Monday")
	case day == "tuesday" || day == "Tuesday":
		fmt.Println("Today is Tuesday")
	default:
		fmt.Println("Today is not Monday or Tuesday")

	}
}
