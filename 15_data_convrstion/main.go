package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "100"
	fmt.Println(a)
	fmt.Printf("type of a is %T\n",a)
	b, _ := strconv.Atoi(a)
	fmt.Println(b)
	fmt.Printf("type of b is %T\n",b)

	c :="4.23"
	fmt.Println("c is :,c")
	fmt.Printf("type of c is %T\n",c)
	d,_ := strconv.ParseFloat(c,64)
	fmt.Println("d is :",d)
	fmt.Printf("type of d is %T\n",d)

	e:="true"
	fmt.Println(e)
	fmt.Printf("type of e is %T\n",e)
	f,_ := strconv.ParseBool(e)
	fmt.Println(f)
	fmt.Printf("type of f is %T\n",f)

}
