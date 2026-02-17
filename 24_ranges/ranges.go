package main

import "fmt"

func main(){
	nums:= []int {1,2,3,4,5}
	fmt.Println(nums)
	sum:=0
	for _,r :=range nums{
		//fmt.Println("index",i,"value",r)
		sum +=r
	}
	fmt.Println("sum of numbers is ",sum)
	//range over maps 
	prs:=map[string]int{"john":20,"asad":55,"qwe":555,"aaaa":5555}
	for k,v := range prs{
		fmt.Println("key",k,"value",v)
	}
	//fmt.Println(add(1,2))
}