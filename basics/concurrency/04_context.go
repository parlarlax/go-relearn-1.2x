package concurrency

import (
	"context"
	"fmt"
	"time"
)

func ExampleContextTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Println("work done")
	case <-ctx.Done():
		fmt.Println("timeout:", ctx.Err())
	}
}

func ExampleContextCancel() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-ctx.Done()
		fmt.Println("goroutine cancelled:", ctx.Err())
	}()

	cancel()
	time.Sleep(10 * time.Millisecond)
}

func ExampleContextValue() {
	type key string
	ctx := context.WithValue(context.Background(), key("userID"), 42)

	if v, ok := ctx.Value(key("userID")).(int); ok {
		fmt.Println("userID:", v)
	}
}
