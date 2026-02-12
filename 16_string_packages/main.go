package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "hello world"
	fmt.Println(str)

	fmt.Println(len(str))

	fmt.Println(str[1:4]) //ello 1 inclusive 4 exclusive

	fmt.Println(str[1:]) //ello world

	fmt.Println(str[:4]) //hello
	var parts = strings.Fields(str)
	fmt.Println(parts)

	str1 := "  The Go,Programming,Language  "
	p2 := strings.Split(str1, "-")
	fmt.Println(p2)

	programming := ""
	p3 := strings.Count(str1 , programming)
	fmt.Println(p3)

}
