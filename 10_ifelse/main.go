package main

import "fmt"

func main() {

	numi := 0
	fmt.Println("enter number")
	fmt.Scan(&numi)
	if numi > 0 && numi <= 10 {
		fmt.Println("number is between 0 and 10")
	}else if numi >= 10 && numi < 20 {
		fmt.Println("number is beteen 10 and 20")
	} else if numi >= 20 && numi < 40 {
		fmt.Println("number is bettween 30 and 40 ")
	} else {
		fmt.Println("number is above 40 ")
	}
}

// func main() {
// 	numi := 0
// 	fmt.Println("Enter number:")
// 	fmt.Scan(&numi)

// 	if numi < 0 {
// 		fmt.Println("number is negative")
// 	} else if numi == 0 {
// 		fmt.Println("number is 0")
// 	} else if numi > 0 && numi < 10 {
// 		fmt.Println("number is between 1 and 9")
// 	} else if numi >= 10 && numi < 20 {
// 		fmt.Println("number is between 10 and 19")
// 	} else if numi >= 20 && numi < 30 {
// 		fmt.Println("number is between 20 and 29")
// 	} else if numi >= 30 && numi < 40 {
// 		fmt.Println("number is between 30 and 39")
// 	} else {
// 		fmt.Println("number is 40 or above")
// 	}
// }

