package main

import (
	"fmt"
	//"strings"
)

type box[T any] struct{
	value T
}

func (b box[T]) getvalue() T{
	return b.value
}



// func add[T int | float64 | string](a T, b T) T {
// 	return a+b
// }
// func sub[T int | float64](a T,b T)T{
// 	return a - b

// }
func main(){
	// fmt.Println(add(1,2))
	// fmt.Println(add(12.222,111222222.6))
	// fmt.Println(add("1 ","2"))
	// fmt.Println(sub(15,21))
	b1:=box[int]{value: 655}
	b2:=box[string]{value: "hello world"}
	fmt.Println(b1.getvalue())
	fmt.Println(b2.getvalue())
	//fmt.Println(add[int](1,2))


}