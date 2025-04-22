package main

import "fmt"

func main(){

	/*
	Rules:
		A variable name must start with a letter or an underscore character (_)
		A variable name cannot start with a digit
		A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
		Variable names are case-sensitive (age, Age and AGE are three different variables)
		There is no limit on the length of the variable name
		A variable name cannot contain spaces
		The variable name cannot be any Go keywords
	*/

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

	// Declaring multiple variables on single line
	// Declaring and initializing 
	var a, b, c = 1,2,3

	// only declaration 
	var d, e int

	fmt.Println(a, b, c)
	fmt.Println(d, e)

	// constants

	const PI float64 = 3.14
	const G = 6.67
	fmt.Println(PI, G)
}