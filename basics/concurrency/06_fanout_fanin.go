package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func ExampleFanOutFanIn() {
	producer := func(items ...int) <-chan int {
		ch := make(chan int)
		go func() {
			for _, item := range items {
				ch <- item
			}
			close(ch)
		}()
		return ch
	}

	worker := func(id int, in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for v := range in {
				out <- v * v
			}
			close(out)
		}()
		return out
	}

	fanIn := func(channels ...<-chan int) <-chan int {
		merged := make(chan int)
		var wg sync.WaitGroup
		for _, ch := range channels {
			wg.Add(1)
			go func(c <-chan int) {
				defer wg.Done()
				for v := range c {
					merged <- v
				}
			}(ch)
		}
		go func() {
			wg.Wait()
			close(merged)
		}()
		return merged
	}

	in := producer(1, 2, 3, 4, 5, 6)

	w1 := worker(1, in)
	w2 := worker(2, in)

	_ = time.Sleep

	for result := range fanIn(w1, w2) {
		fmt.Println(result)
	}
}
