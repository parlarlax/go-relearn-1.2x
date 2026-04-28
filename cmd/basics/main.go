package main

import (
	"fmt"
	"os"

	"github.com/lax/go-relearn/basics/interfaces"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Go Re-learning Lab — Zero to Hero")
		fmt.Println()
		fmt.Println("Usage: go run . <topic>")
		fmt.Println()
		fmt.Println("Topics:")
		fmt.Println("  interfaces   — Phase 1.1: Interface & Methods")
		fmt.Println()
		fmt.Println("Or run experiments directly:")
		fmt.Println("  go run ./experiments/slog-demo")
		fmt.Println("  go run ./experiments/http-mux")
		fmt.Println("  go run ./experiments/iterators-lab")
		return
	}

	switch os.Args[1] {
	case "interfaces":
		interfaces.ExampleBasic()
		interfaces.ExampleComposition()
		interfaces.ExampleEmptyInterface()
		interfaces.ExampleTypeAssertion()
		interfaces.ExampleReceiverTypes()
		interfaces.ExampleNilInterface()
	default:
		fmt.Printf("unknown topic: %s\n", os.Args[1])
	}
}
