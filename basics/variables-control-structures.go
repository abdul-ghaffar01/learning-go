package main

import "fmt"

func Conditionals() {
	fmt.Println("-----------Conditionals start-----------")
	/*
		Conditional control structures:
		Syntax:

		if condition {
			Code
		}

		Parenthesis () are not allowed
	*/

	age := 10

	if age < 20 {
		fmt.Println("you are a teenager")
	} else {
		fmt.Println("You are not a teenager anymore.")
	}

	score := 79

	if score > 90 {
		fmt.Println("A")
	} else if score > 80 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

	// Go allows a short statement before condition
	if x := 50; x > 5 { // Scope of x is in the if block
		fmt.Println("X is greater than 5")
	}

	// ----------- Switch Statement -----------
	// Simple switch
	day := 5
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Frdiay")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")

	}

	// Switch with multiple conditions in one case
	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("Weekday")
	case 6, 7:
		fmt.Println("Weekend")
	}

	// Switch with initialization
	switch x := 50; x {
	case 50:
		fmt.Println("X is 50")
	}

	// Switch without expression
	switch {
	case score >= 90:
		fmt.Println("A")
	case score >= 80:
		fmt.Println("B")
	case score < 80:
		fmt.Println("C")
	}

	// fallthrough executes the next case without satisfying the condition
	switch x := 1; x {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
	}

}

func Loops() {
	/*
		Loop control structures:
		Syntax:
		for initialization/condition/incrementation{
			code
		}

		It is the only loop in Go.
	*/

	// Simple for loop
	for i := 0; i < 10; i++ {
		fmt.Println("The value of i is: ", i)
	}

	// For loop as while
	i := 0
	for i < 10 {
		fmt.Println("The value of i again is: ", i)
		i++
	}

	// infinite for loop
	// for {  // used in servers and workers
	// 	fmt.Println("Infinity")
	// }

	// Looping through array/slice
	nums := []int{10, 20, 30, 40}
	for index, nums := range nums { // index or value can be ignored using _
		fmt.Println(index, nums)
	}

	// Looping through map
	m := map[string]int{"a": 1, "b": 2}
	for key, value := range m {
		fmt.Println(key, value)
	}

	// Looping through string
	for i, char := range "string" {
		/*
			Range iterates over runes not bytes in a string.
			Rune: It is the byte code for each character as some characters are of size from
			1 - 4 bytes so prune helps decode correctly
			type of rune is int32
		*/
		fmt.Printf("%d -> %c\n", i, char)
	}

	/*
		Break/Continue:
			Break is used to break the loop, Continue is used to skip the iteration.
		
		Break outer:
		It breaks both of the loops if loops are nested.
		
	*/

}

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
	Loops()
}
