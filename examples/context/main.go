package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== 1. WithTimeout ===")
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Println("work done")
	case <-ctx.Done():
		fmt.Println("timeout:", ctx.Err())
	}

	fmt.Println("\n=== 2. WithCancel ===")
	ctx2, cancel2 := context.WithCancel(context.Background())
	go func() {
		<-ctx2.Done()
		fmt.Println("  goroutine: cancelled:", ctx2.Err())
	}()
	cancel2()
	time.Sleep(50 * time.Millisecond)

	fmt.Println("\n=== 3. WithValue ===")
	type key string
	ctx3 := context.Background()
	ctx3 = context.WithValue(ctx3, key("requestID"), "req-12345")
	ctx3 = context.WithValue(ctx3, key("userID"), 42)
	if reqID, ok := ctx3.Value(key("requestID")).(string); ok {
		fmt.Println("  requestID:", reqID)
	}
	if uid, ok := ctx3.Value(key("userID")).(int); ok {
		fmt.Println("  userID:", uid)
	}

	fmt.Println("\n=== 4. Propagation — parent cancels children ===")
	parent, cancelParent := context.WithCancel(context.Background())
	defer cancelParent()
	child1, cancel1 := context.WithTimeout(parent, 10*time.Second)
	defer cancel1()
	child2, cancel2 := context.WithCancel(parent)
	defer cancel2()
	cancelParent()
	fmt.Println("  child1:", child1.Err())
	fmt.Println("  child2:", child2.Err())

	fmt.Println("\n=== 5. Deadline ===")
	ctx4, cancel4 := context.WithDeadline(context.Background(), time.Now().Add(1*time.Hour))
	defer cancel4()
	if dl, ok := ctx4.Deadline(); ok {
		fmt.Printf("  deadline: %v (in %v)\n", dl, time.Until(dl).Round(time.Second))
	}

	fmt.Println("\n=== 6. Pattern: select with context ===")
	ctx5, cancel5 := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel5()
	result := slowOperation(ctx5)
	fmt.Println("  result:", result)

	fmt.Println("\n=== 7. Pattern: context in function signature ===")
	doWork(context.Background())
}

func slowOperation(ctx context.Context) string {
	select {
	case <-time.After(200 * time.Millisecond):
		return "completed"
	case <-ctx.Done():
		return ctx.Err().Error()
	}
}

func doWork(ctx context.Context) {
	if deadline, ok := ctx.Deadline(); ok {
		fmt.Printf("  work: deadline in %v\n", time.Until(deadline).Round(time.Second))
	} else {
		fmt.Println("  work: no deadline")
	}
}
