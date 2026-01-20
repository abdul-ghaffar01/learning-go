package main

import "fmt"

func main() {
	name := "Abdul Ghaffar"
	age := 22.4

	// Print with new line
	fmt.Println("Working just fine")

	// Print with format specifiers
	fmt.Print("Hello world\n");

	// Pass arguments in the print
	fmt.Print("Hello", "World", 32, 45, "\n");  // adds space automatically if arguments are numbers.
	fmt.Print(32, 34, 534, 54, "\n");
	fmt.Print("Check", "The", "Space\n")
	fmt.Print(name, age, "\n")
	n, _ := fmt.Print(name, age, "\n")  // returns number of bytes/characters
	fmt.Print("Printed number of bytes: ", n, "\n")
	
	type Person struct{
		name string
		age float64
	};
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


	
}
