package main

import "fmt"

// Two seperate functions for sum
func addInt(a, b int) int {
	return a + b
}

func addFloat(a, b float64) float64 {
	return a + b
}

// To overcome redundancy we use generics
func add[T int | float64](a, b T) T {
	return a + b
}

// any can be replaced with interface{}
func Example[T any](v T) {
	fmt.Println(v)
}

// Custom constraints
type Number interface {
	int | float64 | int64
}

func addLimited[T Number](a, b T) T {
	return a + b
}

// Generics with structs
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) push(value T) {
	s.items = append(s.items, value)
}

func (s *Stack[T]) pop() T {
	return s.items[len(s.items)-1]
}

func Generics() {
	fmt.Println("----------Generics-----------")

	fmt.Println(add(2, 6))
	fmt.Println(add(29.67, 6.76))
}
