package main

import (
	"fmt"
	"sync"
	"time"
)

func basicGoroutine() {
	go func() {
		fmt.Println("goroutine: hello from background!")
	}()
	fmt.Println("main: hello from main")
	time.Sleep(100 * time.Millisecond)
}

func waitGroup() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("  worker %d done\n", id)
		}(i)
	}

	fmt.Println("waiting for workers...")
	wg.Wait()
	fmt.Println("all workers done")
}

func channel() {
	ch := make(chan string)

	go func() {
		ch <- "hello from channel"
	}()

	msg := <-ch
	fmt.Println("received:", msg)
}

func bufferedChannel() {
	ch := make(chan int, 3)

	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Println("buffered:", <-ch, <-ch, <-ch)
}

func selectExample() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "fast"
	}()

	go func() {
		time.Sleep(150 * time.Millisecond)
		ch2 <- "slow"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("ch1:", msg)
		case msg := <-ch2:
			fmt.Println("ch2:", msg)
		}
	}
}

func producerConsumer() {
	ch := make(chan int, 5)

	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			fmt.Printf("  produced: %d\n", i)
		}
		close(ch)
	}()

	for val := range ch {
		fmt.Printf("  consumed: %d\n", val)
	}
}

func main() {
	fmt.Println("=== basic goroutine ===")
	basicGoroutine()

	fmt.Println("\n=== waitgroup ===")
	waitGroup()

	fmt.Println("\n=== channel ===")
	channel()

	fmt.Println("\n=== buffered channel ===")
	bufferedChannel()

	fmt.Println("\n=== select ===")
	selectExample()

	fmt.Println("\n=== producer-consumer ===")
	producerConsumer()
}
