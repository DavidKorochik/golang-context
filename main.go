package main

import (
	"context"
	"fmt"
	"time"
)

func executeContextLoop(ctx context.Context) {
	reqId := ctx.Value("request-id").(string)
	fmt.Println(reqId)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Done with the function")
		default:
			fmt.Println("Doing something with the function")
		}
	}
}

func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request-id", "12345")
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	ctx = enrichContext(ctx)
	go executeContextLoop(ctx)

	if ctx.Err() != nil {
		fmt.Println("Finished executing, exceeded")
	}
}
