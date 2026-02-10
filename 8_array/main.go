package main

import "fmt"

func main() {
	//empty arrays (int) have 0 values
	//empty arrays (string) have "" values
	//empty arrays (bool) have false values
	//empty arrays (float) have 0.0 values
	//in array indexing starts with zero


	//initializing arrays are of different types 
	fmt.Println("arrays ")
	var name [5]string	
	name[0] = "raju"
	name[1] = "kumar"
	name[2] = "singh"
	name[4] = "singh2"
	fmt.Printf("array names are %q\n", name) //%q is used to print string with double quotes

	var num = [8]float32{1, 2, 3, 4, 5}
	fmt.Print(num)

	var number = [5]int{1, 2, 3, 4, 5}
	fmt.Println(number[2])

}
