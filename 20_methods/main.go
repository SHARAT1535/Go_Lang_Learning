package main

import "fmt"

type person struct {
	Firstname string
	Lastname  string
	Age       int
}

func (u *person) Printdetails() {
	// u.Firstname = "RAJ"
	// u.Lastname = "ADDRESS"
	// u.Age = 56
	fmt.Println("printed from methods",u)
	u.Firstname = "ramesh"
	u.Lastname = "singh"
	fmt.Println("in methods valus is  ",u)
	// fmt.Println(u.Lastname)
	// fmt.Println(u.Age)
}

func main() {
	p1 := person{
		Firstname: "RAJ",
		Lastname:  "ADDRESS",
		Age:       56,
	}
	//fmt.Println(p1)
	var a=&p1
	p1.Printdetails()
	fmt.Println("val printed from main",a)
	

	//i += 1
}
