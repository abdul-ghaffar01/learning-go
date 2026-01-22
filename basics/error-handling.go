package main

import (
	"errors"
	"fmt"
)

// Error is basically a type in go and can be treated like other types nothing fancy.

// An error is anything that implements the interface given below.

// The foundation of error handling is this interface
type error interface{
	Error() string
}


func ReturnError() error {
	return errors.New("This is an error")
}

var ErrorDividingByZero = errors.New("Can't divide with 0")

func divide1(a, b int) (int, error){
	if b == 0 {
		return 0, ErrorDividingByZero
	}
	return a/b, nil
}


func ErrorHandling(){
	fmt.Println("----------Error Handling-----------")
	ans, err := divide1(3, 0)
	if err != nil {
		fmt.Print(err.Error(), "\n")
		return;
	}
	fmt.Println("Division answer is ", ans)
	
}