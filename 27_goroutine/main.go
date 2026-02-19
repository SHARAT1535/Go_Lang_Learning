package main
//goroutine example: make program to run concurrently using goroutine
import (
	"fmt"
	"time"
)

func func1() {
	fmt.Println("Func1 started")
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("Func1 ended")

}

func func2(){
	fmt.Println("func2 started")
	fmt.Println("func2 ended")
}

func main(){
	fmt.Println("hello world from main function")
	time.Sleep(10000*time.Millisecond)
	 func1()
	 	time.Sleep(3000 * time.Millisecond)

	go func2()
		time.Sleep(10000*time.Millisecond)

}