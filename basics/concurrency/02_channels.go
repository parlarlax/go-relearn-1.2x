package concurrency

import (
	"fmt"
)

func ExampleChannels() {
	ch := make(chan string, 2)

	ch <- "first"
	ch <- "second"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func ExampleChannelDirection() {
	producer := func(ch chan<- string) {
		ch <- "sent from producer"
	}

	consumer := func(ch <-chan string) string {
		return <-ch
	}

	ch := make(chan string, 1)
	producer(ch)
	msg := consumer(ch)
	fmt.Println(msg)
}
