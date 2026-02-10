package main

import "fmt"


func main() {

	//properties
	// dynamic size slice can grow or shirink dynamically during runtime 
	// reference to underlying array
	//syntax the syntax of slice is same as array 
	//use function make to create slice
	//appending element the append function is used to add elements to  a slice  


	
	//	numbers := []int{1,2,3,4}
	// fmt.Println("numbers are ",numbers)
	// fmt.Printf("numbers data type : %T\n",numbers)
	// fmt.Println("length",len(numbers))

	//name := []string{}
	//	fmt.Println("name",name)
	//	fmt.Printf("name data type : %T\n",name)
	//	fmt.Println("length",len(name))
	//	fmt.Println("capacity is",cap(name))
	//name = append(name, "raj")
	//fmt.Printf("name data type : %T\n", name)
	//fmt.Println("names in name slice are", name)
	//	fmt.Println("length",len(name))
	//	fmt.Println("capacity is",cap(name))
	num := make([]int, 1, 5) // 3 is length of slice ,5 capacity of slice 
	num = append(num, 1, 2, 3, 4,6,7,8,9,0,10,11)
	fmt.Println("num", num)
	fmt.Printf("num data type : %T\n", num)
	fmt.Println("length", len(num))
	fmt.Println("capacity is", cap(num))

}
