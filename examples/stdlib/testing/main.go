package main

import (
	"fmt"
	"time"
)

func add(a, b int) int {
	return a + b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println("=== 1. Basic test pattern ===")
	got := add(2, 3)
	want := 5
	if got != want {
		fmt.Printf("FAIL: add(2,3) = %d, want %d\n", got, want)
	} else {
		fmt.Println("PASS: add(2,3) =", got)
	}

	fmt.Println("\n=== 2. Table-driven test ===")
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive", 2, 3, 5},
		{"negative", -1, -2, -3},
		{"zero", 0, 0, 0},
		{"mixed", 5, -3, 2},
	}
	for _, tc := range tests {
		got := add(tc.a, tc.b)
		status := "PASS"
		if got != tc.want {
			status = "FAIL"
		}
		fmt.Printf("  %s: add(%d,%d) = %d (%s)\n", tc.name, tc.a, tc.b, got, status)
	}

	fmt.Println("\n=== 3. Error testing ===")
	_, err := divide(10.0, 0)
	if err == nil {
		fmt.Println("  FAIL: expected error")
	} else {
		fmt.Println("  PASS: got error:", err)
	}

	result, err := divide(10.0, 2.0)
	if err != nil {
		fmt.Println("  FAIL: unexpected error:", err)
	} else {
		fmt.Printf("  PASS: divide(10,2) = %f\n", result)
	}

	fmt.Println("\n=== 4. Benchmark pattern ===")
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		add(i, i+1)
	}
	fmt.Printf("  1M adds in %v\n", time.Since(start))

	fmt.Println("\n=== 5. T.Context pattern (Go 1.24+) ===")
	fmt.Println("  (in real test: ctx := t.Context() auto-cancels when test ends)")

	fmt.Println("\n=== 6. B.Loop pattern (Go 1.24+) ===")
	fmt.Println("  (in real benchmark: for b.Loop() { add(1, 2) })")

	fmt.Println("\n=== 7. synctest (Go 1.25+) ===")
	fmt.Println("  (in real test: synctest.Test(t, func(t *testing.T) { ... })")
	fmt.Println("  time is virtualized, synctest.Wait() blocks until all goroutines sleep)")

	fmt.Println("\n=== 8. Subtests ===")
	cases := []string{"case_a", "case_b", "case_c"}
	for _, name := range cases {
		fmt.Printf("  === RUN  TestGroup/%s\n", name)
		fmt.Printf("  --- PASS: TestGroup/%s\n", name)
	}

	fmt.Println("\n=== 9. Testable example ===")
	fmt.Println("  Output: add(1,2) =", add(1, 2))
}
