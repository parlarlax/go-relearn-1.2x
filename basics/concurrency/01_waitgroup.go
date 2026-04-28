package concurrency

import (
	"fmt"
	"sync"
)

func ExampleWaitGroup() {
	var wg sync.WaitGroup

	tasks := []string{"task-A", "task-B", "task-C"}

	for _, task := range tasks {
		wg.Add(1)
		go func(t string) {
			defer wg.Done()
			fmt.Printf("finished %s\n", t)
		}(task)
	}

	wg.Wait()
	fmt.Println("all tasks done")
}
