package main

import "fmt"

/*
Functions: A piece of code used to perform some specific task

Syntax:

func identifier(parameters) returnType {
	Code
}

Functions can be:
	assigned to a varialbe
	passed as argument
	returned as a value

*/

// Simple function
func sum(a int, b int) int {
	return a + b
}

// Same type of parameters can be grouped
func sum1(a, b, c int) int {
	return a + b + c
}

// Without any return
func PrintPattern(rows, cols int) {
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if i == j || i+j == cols-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}

// Multiple return values
func divide(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, fmt.Errorf("Can't divide with 0")
	}
	return a / b, a % b, nil
}

// Named return values
func subtract(a, b int) (result int ){
	result = a-b
	return
}

func multiply(a, b int) int {
	return a * b;
}

// Variadic Function
func sumAll(nums ...int) int {
	total := 0
	for _, num:= range nums{
		total += num;
	}
	return total;
}

// Closure
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count;
	}
}