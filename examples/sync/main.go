package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("=== 1. Mutex / RWMutex ===")
	var mu sync.RWMutex
	counter := 0
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("counter:", counter)

	fmt.Println("\n=== 2. WaitGroup.Go (Go 1.25+) ===")
	var wg2 sync.WaitGroup
	for _, task := range []string{"alpha", "beta", "gamma"} {
		t := task
		wg2.Go(func() {
			fmt.Printf("  task %s done\n", t)
		})
	}
	wg2.Wait()

	fmt.Println("\n=== 3. Once ===")
	var once sync.Once
	for i := 0; i < 5; i++ {
		once.Do(func() {
			fmt.Println("  this runs only once!")
		})
	}

	fmt.Println("\n=== 4. WaitGroup + result collection ===")
	results := make([]string, 3)
	var wg3 sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg3.Add(1)
		go func(idx int) {
			defer wg3.Done()
			time.Sleep(time.Duration(idx*10) * time.Millisecond)
			results[idx] = fmt.Sprintf("result-%d", idx)
		}(i)
	}
	wg3.Wait()
	fmt.Println("results:", results)

	fmt.Println("\n=== 5. Cond ===")
	var cond = sync.NewCond(&sync.Mutex{})
	ready := false
	go func() {
		cond.L.Lock()
		ready = true
		cond.Broadcast()
		cond.L.Unlock()
	}()
	cond.L.Lock()
	for !ready {
		cond.Wait()
	}
	cond.L.Unlock()
	fmt.Println("cond: ready signal received")

	fmt.Println("\n=== 6. Pool ===")
	pool := &sync.Pool{
		New: func() any {
			fmt.Println("  pool: allocating new buffer")
			return make([]byte, 1024)
		},
	}
	buf1 := pool.Get().([]byte)
	fmt.Printf("  pool: got buf len=%d\n", len(buf1))
	pool.Put(buf1)
	buf2 := pool.Get().([]byte)
	fmt.Printf("  pool: reused buf len=%d\n", len(buf2))

	fmt.Println("\n=== 7. Map (concurrent map) ===")
	var m sync.Map
	m.Store("key1", "value1")
	m.Store("key2", "value2")
	if v, ok := m.Load("key1"); ok {
		fmt.Println("  map load:", v)
	}
	m.Range(func(k, v any) bool {
		fmt.Printf("  map: %v = %v\n", k, v)
		return true
	})
	m.Delete("key1")
	if _, ok := m.Load("key1"); !ok {
		fmt.Println("  map: key1 deleted")
	}
}
