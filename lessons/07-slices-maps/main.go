package main

import "fmt"

func main() {
	fmt.Println("=== slice basics ===")
	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println("fruits:", fruits)
	fruits = append(fruits, "date")
	fmt.Println("after append:", fruits)

	fmt.Println("\n=== slice operations ===")
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println("full:", nums)
	fmt.Println("nums[1:3]:", nums[1:3])
	fmt.Println("nums[:3]:", nums[:3])
	fmt.Println("nums[2:]:", nums[2:])
	fmt.Println("len:", len(nums), "cap:", cap(nums))

	fmt.Println("\n=== make + copy ===")
	a := make([]int, 3, 10)
	fmt.Println("make(3,10):", a, "len:", len(a), "cap:", cap(a))

	src := []int{10, 20, 30}
	dst := make([]int, len(src))
	n := copy(dst, src)
	fmt.Println("copied:", n, "dst:", dst)

	fmt.Println("\n=== map basics ===")
	scores := map[string]int{
		"Alice": 100,
		"Bob":   85,
	}
	scores["Charlie"] = 90
	fmt.Println("scores:", scores)

	val := scores["Alice"]
	fmt.Println("Alice:", val)

	val, ok := scores["David"]
	fmt.Println("David:", val, "exists:", ok)

	delete(scores, "Bob")
	fmt.Println("after delete:", scores)

	fmt.Println("\n=== range iteration ===")
	for i, v := range fruits {
		fmt.Printf("  fruits[%d] = %s\n", i, v)
	}
	for k, v := range scores {
		fmt.Printf("  %s = %d\n", k, v)
	}
}
