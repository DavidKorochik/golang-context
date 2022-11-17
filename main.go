package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func addNumbersToChannel(ctx context.Context, numsCh chan int) {
	go doSomethingWithTheCtx(ctx, numsCh)

	for num := 1; num <= 6; num++ {
		numsCh <- num
	}

}

func doSomethingWithTheCtx(ctx context.Context, numsCh <-chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context has finished executing within the time passed")
			return
		case printNum := <-numsCh:
			fmt.Println(printNum)
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func addKeysValuesToCtx(ctx context.Context) context.Context {
	return context.WithValue(ctx, "user_id", "id12345")
}

func main() {
	numsCh := make(chan int)

	// We set a context with a total of 2 seconds of execution time
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// We cancel the execution of the context in the end of the main function
	defer cancel()

	// Add to the parent context the user_id key and its associated value
	ctx = addKeysValuesToCtx(ctx)

	// Add numbers to the channel
	addNumbersToChannel(ctx, numsCh)

	// Get the number of go routines just for logging purposes
	fmt.Println(runtime.NumGoroutine())

	time.Sleep(2 * time.Second)
}
