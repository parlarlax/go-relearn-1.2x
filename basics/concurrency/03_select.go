package concurrency

import (
	"fmt"
	"time"
)

func ExampleSelect() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "from ch1"
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch2 <- "from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println(msg)
		case msg := <-ch2:
			fmt.Println(msg)
		}
	}
}

func ExampleSelectDefault() {
	ch := make(chan string)

	select {
	case msg := <-ch:
		fmt.Println(msg)
	default:
		fmt.Println("no message ready")
	}
}
