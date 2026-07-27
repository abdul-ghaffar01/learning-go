package main

import (
	"fmt"
)

func main() {
	name := "Abdul Ghaffar"
	age := 22.4

	// Print with new line
	fmt.Println("Working just fine")

	// Print with format specifiers
	fmt.Print("Hello world\n")

	// Pass arguments in the print
	fmt.Print("Hello", "World", 32, 45, "\n") // adds space automatically if arguments are numbers.
	fmt.Print(32, 34, 534, 54, "\n")
	fmt.Print("Check", "The", "Space\n")
	fmt.Print(name, age, "\n")
	n, _ := fmt.Print(name, age, "\n") // returns number of bytes/characters
	fmt.Print("Printed number of bytes: ", n, "\n")

	type Person struct {
		name string
		age  float64
	}
	p := Person{name, age}
	// Can print any type String, int, array, slice, struct
	fmt.Print("\nString", 22, [2]int{2, 3}, true, 3.2783, p, "\n")

	// fomatting verbs
	fmt.Printf("My name is %q, and I am %f years old.\n", name, age)

	/*
		%v for value
		%#v for go-syntax value
		%s for string
		%d for int
		%f for float
		%q for quoted string
		%T for type
		%t for boolean
		%x for hexadecimal
		%p for pointer address
	*/

	// Formating floats
	temp := 3443.2348

	fmt.Printf("Temp with 2 decimal places %.2f\n", temp)
	fmt.Printf("Temp with 2 decimal places and 8 width %08.2f \n", temp)
	fmt.Printf("Temp with 2 decimal places and 8 width %-8.2fcheck\n", temp)

	// Use of arguments postion
	fmt.Printf("%[2]s is %.1[1]f years old\n", age, name)

	// Sprintf returns the string instead of printing it
	str := fmt.Sprintf("%f is my age", age)
	fmt.Println(str)

	// Why to use these functions instead of string concatenation ?
	/*
		String concatenation is handy when we have small number of sub strings but becomes really
		complecated when it comes to huge paramters.
		and when we need to concatenate different types better to use Sprintf because in concatenation
		we will have to write a lot of extra code.
	*/

	Conditionals()
	fmt.Println("--------------------Loops--------------------")
	Loops()
	Variables()

	// Functions
	PrintPattern(5, 5)
	total := sumAll(3, 4, 5, 6, 3, 54, 5, 1)
	fmt.Println(total)


	var fu func(int, int) int
	fu = multiply
	fmt.Println(fu(3,5))

	// Anonymous Functions
	func() {
		fmt.Println("Anonymous function")
	}()

	// Closure 
	c := counter()
	fmt.Println(c())
	fmt.Println(c())
	fmt.Println(c())

	// structs and methods in go
	structsAndMethods()

	// Interface and polymorphisim
	interfacesAndPolymorphism()

	// Generics
	Generics()

	// Error Handling
	ErrorHandling()

	// Json encoding/decoding
	JsonEncodingDecoding()

	// Error recovery
	ErrorRecovery()
}
