package main

import "fmt"

func getData(name string, ch chan string) {
	// Sending into channel
	ch <- "AGS"
}

// Create a channel of values of slices
func Task1() {
	fmt.Println("Task 1")
	nums := []int{2, 6, 7, 4}

	ch := make(chan int)

	// Sending everything to the channel
	go func() {
		for _, value := range nums {
			ch <- value
		}
		close(ch)
	}()

	// recieving from channel
	for value := range ch {
		fmt.Println(value)
	}
}

// Task 2  ------------------------------------------

func convertToChannel(nums []int) <-chan int {
	out := make(chan int)

	go func() {
		for _, value := range nums {
			fmt.Println("In ctc for value ", value)
			out <- value
		}
		close(out)
	}()
	return out
}

func makeSquares(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for value := range in {
			fmt.Println("In ms for value ", value)
			out <- value * value
		}
		close(out)
	}()
	return out
}

func Task2() {
	nums := []int{3, 6, 2, 5}

	firstStage := convertToChannel(nums)
	secondStage := makeSquares(firstStage)

	// third stage
	for value := range secondStage {
		fmt.Println("In print for value ", value)
		fmt.Println(value)
	}
}

func BufferedChannels() {
	fmt.Println("Buffered channels ---------------------")
	nums := []int{2, 6, 3, 7}
	ch := make(chan int, len(nums))

	// Writing to the channel
	// go func(){
	// 	for _, value := range nums{
	// 		ch <- value
	// 	}
	// 	close(ch)
	// }()

	// Alternative
	for _, value := range nums {
		ch <- value
	}
	close(ch)

	for value := range ch {
		fmt.Println(value)
	}
}

func Channels() {
	fmt.Println("-------Channels--------")

	// Creating a channel
	ch := make(chan string)

	go getData("Abdul Ghaffar", ch) // goroutine

	// Receiving the data from channel in main
	data := <-ch // Because of this main function waits for the goroutine

	fmt.Printf("%v\n", data)

	// Task 1: Convert a slice into a channel
	Task1()

	// Task 2: Do all the work in stages
	// Stage 1: Convert slice into a channel
	// Stage 2: Make the square of each value of channel
	// Stage 3: Print all the squared numbers
	Task2()

	BufferedChannels()

}

/*
Key notes:
- Channels are mechanism of communication between goroutines.
- Range works on closed channels else it waits forever and is a deadlock situation
- Channel must be closed
- Buffered channels has fixed size and doesn't block until the buffer is full it pushes without blocks.
- Unbuffered channels are of size 1 means everything that is written must be read at the same time
else it blocks the execution flow.
*/
