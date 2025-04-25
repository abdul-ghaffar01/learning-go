package main

import "fmt"

func main(){
	// Syntax of creating an array
	// var array_name = [length]datatype{values}
	
	// Creating an array
	var arr = [3]int{1,2,3}
	fmt.Println(arr)

	// Length is inferred
	var arr1 = [...]int{4,5,6,67,7,8}
	fmt.Println(arr1)

	// Using := 
	arr2 := [3]int{34,6,7}
	fmt.Println(arr2)

	// length inferred
	arr3 := [...]int{87,324,6,422,32,34,5}
	fmt.Println(arr3)

	// Creating an string array
	var cars = [...]string{"Mehran", "audi", "Fortuner", "Something else"}
	fmt.Println(cars)

	// Accessing the elements of an array using indexes
	fmt.Println(arr[0])

	// Changing the values
	arr[0] = 78
	fmt.Println(arr)

	// Initialization of array Note: if not fully initialized takes the default value of type
	// Not initialized
	arr4 := [5]int{}
	fmt.Println(arr4)

	// Partialy initialized
	arr5 := [6]int{23,45}
	fmt.Println(arr5)

	// Fully initialized
	arr6 := [3]int{45,77,57}
	fmt.Println(arr6)

	// Initializing only specific indexes
	arr7 := [8]int{2:89, 3:78, 6:834}
	fmt.Println(arr7)

	// Finding the length of the array
	fmt.Println("The length of arr7 is", len(arr7))
}