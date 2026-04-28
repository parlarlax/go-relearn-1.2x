package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"weak"
)

type Resource struct {
	ID int
}

func weakPointerExample() {
	r := &Resource{ID: 42}
	wp := weak.Make(r)

	if got := wp.Value(); got != nil {
		fmt.Println("weak pointer alive:", got.ID)
	}

	runtime.GC()

	if got := wp.Value(); got != nil {
		fmt.Println("still alive (not collected yet):", got.ID)
	} else {
		fmt.Println("collected!")
	}

	_ = r
}

func osRootExample() {
	root, err := os.OpenRoot(".")
	if err != nil {
		fmt.Println("open root error:", err)
		return
	}

	f, err := root.Open("README.md")
	if err != nil {
		fmt.Println("open error:", err)
		return
	}
	f.Close()
	fmt.Println("os.Root: opened README.md safely")

	_, err = root.Open("../../../etc/passwd")
	if err != nil {
		fmt.Println("os.Root: blocked path traversal:", err)
	}
}

func cleanupExample() {
	r := &Resource{ID: 99}
	runtime.AddCleanup(r, func(id int) {
		fmt.Printf("cleanup: resource %d collected\n", id)
	}, r.ID)
	fmt.Println("cleanup registered for resource", r.ID)
}

func waitGroupGoExample() {
	var wg sync.WaitGroup

	tasks := []string{"alpha", "beta", "gamma"}
	for _, task := range tasks {
		t := task
		wg.Go(func() {
			fmt.Printf("  WaitGroup.Go: %s done\n", t)
		})
	}
	wg.Wait()
	fmt.Println("WaitGroup.Go: all tasks complete")
}

func main() {
	fmt.Println("=== weak pointer ===")
	weakPointerExample()

	fmt.Println("\n=== os.Root ===")
	osRootExample()

	fmt.Println("\n=== runtime.AddCleanup ===")
	cleanupExample()

	fmt.Println("\n=== sync.WaitGroup.Go (Go 1.25+) ===")
	waitGroupGoExample()
}
