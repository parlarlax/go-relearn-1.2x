package main

import (
	"fmt"
	"maps"
	"slices"
	"sort"
)

func main() {
	fmt.Println("=== slices package ===")

	nums := []int{5, 3, 1, 4, 2}

	fmt.Println("\n--- slices.Sort ---")
	slices.Sort(nums)
	fmt.Println("sorted:", nums)

	fmt.Println("\n--- slices.SortFunc (custom) ---")
	words := []string{"banana", "apple", "cherry"}
	slices.SortFunc(words, func(a, b string) int {
		return len(a) - len(b)
	})
	fmt.Println("by length:", words)

	fmt.Println("\n--- slices.Contains ---")
	fmt.Println("contains 3:", slices.Contains(nums, 3))
	fmt.Println("contains 9:", slices.Contains(nums, 9))

	fmt.Println("\n--- slices.Index ---")
	fmt.Println("index of 4:", slices.Index(nums, 4))

	fmt.Println("\n--- slices.Insert / Delete ---")
	s := []int{1, 2, 3, 4, 5}
	s = slices.Insert(s, 2, 99)
	fmt.Println("insert 99 at 2:", s)
	s = slices.Delete(s, 2, 3)
	fmt.Println("delete index 2:", s)

	fmt.Println("\n--- slices.Replace ---")
	s = slices.Replace(s, 1, 3, 10, 20)
	fmt.Println("replace [1:3] with 10,20:", s)

	fmt.Println("\n--- slices.Clone ---")
	clone := slices.Clone(s)
	clone[0] = 999
	fmt.Println("original:", s, "clone:", clone)

	fmt.Println("\n--- slices.Reverse ---")
	slices.Reverse(s)
	fmt.Println("reversed:", s)

	fmt.Println("\n--- slices.Min / Max ---")
	fmt.Println("min:", slices.Min(nums))
	fmt.Println("max:", slices.Max(nums))

	fmt.Println("\n--- slices.BinarySearch ---")
	sort.Ints(nums)
	i, found := slices.BinarySearch(nums, 3)
	fmt.Println("binary search 3:", "index=", i, "found=", found)

	fmt.Println("\n--- slices.Concat ---")
	joined := slices.Concat([]int{1, 2}, []int{3, 4}, []int{5})
	fmt.Println("concat:", joined)

	fmt.Println("\n--- slices.Collect (from iterator) ---")
	seq := slices.Values([]string{"x", "y", "z"})
	collected := slices.Collect(seq)
	fmt.Println("collect from iterator:", collected)

	fmt.Println("\n\n=== maps package ===")

	m := map[string]int{"apple": 5, "banana": 3, "cherry": 7}

	fmt.Println("\n--- maps.Keys / maps.Values ---")
	keys := maps.Keys(m)
	kSlice := slices.Collect(keys)
	slices.Sort(kSlice)
	fmt.Println("keys:", kSlice)
	vals := maps.Values(m)
	vSlice := slices.Collect(vals)
	slices.Sort(vSlice)
	fmt.Println("values:", vSlice)

	fmt.Println("\n--- maps.Clone ---")
	m2 := maps.Clone(m)
	m2["durian"] = 10
	fmt.Println("original:", m, "clone:", m2)

	fmt.Println("\n--- maps.DeleteFunc ---")
	maps.DeleteFunc(m, func(k string, v int) bool {
		return v < 5
	})
	fmt.Println("after delete v<5:", m)

	fmt.Println("\n--- maps.Equal ---")
	a := map[string]int{"x": 1, "y": 2}
	b := map[string]int{"x": 1, "y": 2}
	c := map[string]int{"x": 1, "y": 3}
	fmt.Println("a==b:", maps.Equal(a, b), "a==c:", maps.Equal(a, c))

	fmt.Println("\n--- maps.Collect (from iterator) ---")
	pairs := maps.All(map[string]int{"a": 1, "b": 2})
	collected2 := maps.Collect(pairs)
	fmt.Println("collect from iterator:", collected2)
}
