package main

import (
	"fmt"
	"sync"
)

func createChan(ch chan<- int, nums []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, v := range nums {
		fmt.Println(v)
		ch <- v
	}
	close(ch)
}

func printSquares(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v * v)
	}
}

func Practice_Channels() {
	// ch := make(chan int)	// Unbuffered channel

	// go func(){ch <- 3}()

	// data := <- ch

	// fmt.Println(data)

	// Task: have slice of numbers print each number and then it's square

	// Creating WaitGroup to wait for go routines to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// slice of nums
	nums := []int{3, 6, 2, 7, 8}

	// declaring channel
	ch := make(chan int)

	go createChan(ch, nums, &wg)

	go printSquares(ch, &wg)

	wg.Wait()
	fmt.Println("Task completed")
}

func Practice() {
	// Practice of channels
	Practice_Channels()
}
