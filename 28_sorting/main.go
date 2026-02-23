package main

import "fmt"

func main(){
	arr:=[]int {1,4,32,6,8,4,2}
	fmt.Println("Before sorting:",arr)
	//bubble sort 
	for i:=0;i<len(arr);i++{
		for j:=0;j<len(arr)-i-1;j++ {
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1]=arr[j+1],arr[j]
			}
		}
	}
	fmt.Println("sorted array is ")
	for i:=0;i<len(arr);i++{
		fmt.Printf("%d \n",arr[i])
	}
}