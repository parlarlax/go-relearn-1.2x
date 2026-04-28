package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func cancelExample() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("  child: work done")
		case <-ctx.Done():
			fmt.Println("  child:", ctx.Err())
		}
	}()

	time.Sleep(100 * time.Millisecond)
	cancel()
	time.Sleep(100 * time.Millisecond)
}

func timeoutExample() {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Println("  work finished")
	case <-ctx.Done():
		fmt.Println("  timeout:", ctx.Err())
	}
}

func valueExample() {
	type requestIDKey struct{}
	ctx := context.WithValue(context.Background(), requestIDKey{}, "req-12345")

	if id, ok := ctx.Value(requestIDKey{}).(string); ok {
		fmt.Println("  request ID:", id)
	}
}

func mutexExample() {
	var mu sync.Mutex
	counter := 0
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("  counter:", counter, "(expected 1000)")
}

func main() {
	fmt.Println("=== context cancel ===")
	cancelExample()

	fmt.Println("\n=== context timeout ===")
	timeoutExample()

	fmt.Println("\n=== context value ===")
	valueExample()

	fmt.Println("\n=== mutex (safe counter) ===")
	mutexExample()
}
