package main

import "fmt"

func main(){

	// Declaring variables using var keyword 
	var name string
	name = "Abdul Ghaffar"

	// var name1 string = "Ahmad"
	// var name2 = "Asghar"

	// Delaring variables without var keyword
	age := 21	// This method can't be used outside of a funciton body.

	// Default values of primitive data types
	var boolean bool
	var integer int
	var str string

	fmt.Println("My name is", name, "and my age is ", age)
	fmt.Println("Printing default values\n", "Bool:", boolean, "\n", "Integer: ", integer, "\n", "String:", str)
}