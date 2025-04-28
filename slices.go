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


	// Modifying Slices

	newSlice := make([]int, 10)
	fmt.Println(newSlice)

	// Changing the values
	newSlice[0] = 32;
	newSlice[3] = 53;
	fmt.Println(newSlice)

	// Append
	newSlice = append(newSlice, 345,645,4643)
	fmt.Println(newSlice)

	// Appending one slice into another one
	newSlice = append(newSlice, myslice...)
	fmt.Println(newSlice)

	// Copy a slice
	// Creating a slice with 3 elements of newSlice
	anotherSlice := newSlice[0:3]
	fmt.Println(anotherSlice)
	
	// Creating an empty slice to copy the elements
	anotherOne := make([]int, 3)
	
	// copying the elements of anotherSlice into anotherone
	copy(anotherOne, anotherSlice)
	fmt.Println("The elements that are copied", anotherOne)
}