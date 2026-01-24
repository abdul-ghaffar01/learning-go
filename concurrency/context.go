package main

import (
	"context"
	"fmt"
	"time"
)

/*
Context:
It cancels the request after some specific time so the request doesn't hang and returns with timeout error
*/

// Interface of context
type Context interface {
	Deadline(deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key any) any
}

func fetchData(ctx context.Context) []int {

	select {
	case <-ctx.Done():
		return []int{}
	case <-time.After(time.Second*2):
		data := []int{3, 5, 2, 5, 4}
		return data
	}

}
func Contexts() {
	fmt.Println("Context --------------------------")

	// Creating a canceling context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Creating a timeout context
	ctx1, cancel1 := context.WithTimeout(ctx, time.Second*3)
	data := fetchData(ctx1)
	fmt.Println("Data: ", data)
	cancel1()

}
