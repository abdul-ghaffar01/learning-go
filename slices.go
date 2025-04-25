package main

import "fmt"

func main (){
	// Creating slice
	var firstSlice = []int{}
	fmt.Println(firstSlice)

	// Initializing the slice
	secondSlice := []int{23,45,6,75,867,45}
	fmt.Println(secondSlice)

	// Length of the slice
	fmt.Println(len(secondSlice))

	// Capacity of the slice
	fmt.Println(cap(secondSlice))

	// Creating a slice from an array
	arr := []int{3,43,234,4,5,6,45,43,23,4,45}
	myslice := arr[2:5]
	fmt.Println(myslice)

}