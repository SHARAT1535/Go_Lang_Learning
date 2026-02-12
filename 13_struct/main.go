package main

import (
	"fmt"
)

// struct should be declared outside clas as we need to access it in other functions also access through out function
type person struct {
	firstname string
	lastname  string
	age       int
}

type contact struct {
	email string
	phone int
}
type address struct {
	city    string
	state   string
	houseno string
}
type employee struct {
	person_details  person
	person_address  address
	person_conatact contact
}

func main() {
	// var person1 person
	// fmt.Println("person1 is ",person1)// output is " "," ",0
	// //initializing struct
	// person1.firstname = "raju"
	// person1.lastname = "singh"
	// person1.age =33

	// fmt.Println("person1 data is ",person1)// output is "raju","singh",33
	// per2  := person{
	// 	firstname:"raju",
	// 	lastname:"singh",
	// 	age:33,
	// }

	// fmt.Println("person 2 age is ",per2.age)
	// fmt.Println("************************************************************")

	var p1 employee
	p1.person_details.firstname = "john"
	p1.person_details.lastname = "don"
	p1.person_details.age = 23

	fmt.Println("details of employee are ", p1) //details of employee are  {{john don 23} {  } {  }}

	p1.person_address = address{
		city:    "dwd",
		state:   "beng",
		houseno: "12",
	}
	fmt.Println("details of employ p1 are", p1)
	fmt.Println("address of emply p1 are", p1.person_address)

	// p1.person_conatact.email: = "test@gmal.com"
	// p1.person_conatact.phone = 123456
	fmt.Println("details of employ p1 are", p1)
	fmt.Println("contact of emply p1 are", p1.person_conatact)
	
	p1.person_conatact.email="test@gmail.com"
	p1.person_conatact.phone=123456
	fmt.Println(p1)

}
