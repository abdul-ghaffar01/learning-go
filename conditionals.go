package main
import "fmt"

func main(){
	a := 7
	b := 3

	// If-else
	if a > b {
		fmt.Println("a is greater than b")
	} else {	// Adding new line before else block will result in error
		fmt.Println("b is greater than a")
	}

	// If-elseIf-else
	x := 5
	y := 8
	z := 3

	if x > y && x > z {
		fmt.Println("x is greater")
	} else if y > z {
		fmt.Println("Y is greater")
	} else {
		fmt.Println("Z is greater")
	}

}