package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}
func (a person) speak(){
	fmt.Println(a.first,"is my name is raj")
	

}

type human struct{
	gender string
	person
}

func main(){
	d:=human{person:person{"raj","s",4},gender: "male"}
	d.speak()
	
}