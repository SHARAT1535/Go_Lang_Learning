package main

import "fmt"

//const pi = 3.14// we can't change the value
const (
	plan =iota //default value 0 if and only if iota is used
	inprogress //default value 1 if and only if iota is used
	done //default value 2 if and only if iota is used

)

func main(){
//r̥	pi = 3264 //rediclation not possible throws error
	fmt.Println(plan)//1
	fmt.Println(inprogress)//2
	fmt.Println(done)	//3
//if and only if iota is used
//iota is used to auto increment the value of const
//if iota is not used, then we have to explicitly set the value of const
}