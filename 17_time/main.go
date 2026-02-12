package main

import (
	"fmt"
	"time"
)

func main() {
	CurrentTime := time.Now()
	fmt.Println("current time is ", CurrentTime)
	fmt.Printf("current time formate is %T\n", CurrentTime)

	formateedtime := CurrentTime.Format("2026/05/2, 15:04:05, Monday")
	fmt.Println(formateedtime)

	layoutstr := "2025-02-01"
	datestr := "2025-02-05"
	formated_time, _ := time.Parse(layoutstr, datestr)
	fmt.Println(formated_time)

}
