package main

import (
	"fmt"
	"time"
)

func printNumber(num int){
	fmt.Println(num)
}

func main() {
	fmt.Println("-------Goroutines-------")

	go printNumber(3)
	go printNumber(4)
	go printNumber(5)

	time.Sleep(time.Second * 1)
	fmt.Println("Main end")

	
	// Channels: Communication between goroutines
	Channels()

}
