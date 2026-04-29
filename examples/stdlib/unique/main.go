package main

import (
	"fmt"
	"unique"
)

func main() {
	fmt.Println("=== 1. unique.Make — intern strings ===")
	h1 := unique.Make("hello")
	h2 := unique.Make("hello")
	h3 := unique.Make("world")

	fmt.Println("h1 == h2:", h1 == h2)
	fmt.Println("h1 == h3:", h1 == h3)
	fmt.Printf("h1 value: %s\n", h1.Value())

	fmt.Println("\n=== 2. Pointer comparison ===")
	fmt.Printf("h1 addr: %p\n", h1)
	fmt.Printf("h2 addr: %p\n", h2)
	fmt.Printf("h3 addr: %p\n", h3)

	fmt.Println("\n=== 3. Use case: enum-like constants ===")
	statusPending := unique.Make("pending")
	statusDone := unique.Make("done")
	taskStatus := statusPending

	fmt.Println("task pending?", taskStatus == statusPending)
	fmt.Println("task done?", taskStatus == statusDone)
	taskStatus = statusDone
	fmt.Println("task done now?", taskStatus == statusDone)

	fmt.Println("\n=== 4. Use case: deduplication ===")
	names := []string{"alice", "bob", "alice", "charlie", "bob", "alice"}
	seen := make(map[unique.Handle[string]]bool)
	var unique_names []string
	for _, n := range names {
		h := unique.Make(n)
		if !seen[h] {
			seen[h] = true
			unique_names = append(unique_names, n)
		}
	}
	fmt.Println("original:", names)
	fmt.Println("deduped:", unique_names)

	fmt.Println("\n=== 5. Use case: map key (cheaper than string) ===")
	type Role unique.Handle[string]
	admin := Role(unique.Make("admin"))
	user := Role(unique.Make("user"))
	permissions := map[Role][]string{
		admin: {"read", "write", "delete"},
		user:  {"read"},
	}
	r := admin
	fmt.Println("admin perms:", permissions[r])
}
