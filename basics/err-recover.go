package main

import "fmt"

func A() {
	panic("Boom")
}

func B() {
	A()
}

func C() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered: ", r)
		}
	}()
	B()
}

func ErrorRecovery() {
	C()
	fmt.Println("Working just fine")
}
