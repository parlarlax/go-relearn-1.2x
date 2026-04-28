package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	name := "Go learner"
	age := 25
	fmt.Printf("Name: %s, Age: %d\n", name, age)

	msg := fmt.Sprintf("Welcome, %s!", name)
	fmt.Println(msg)

	person := struct {
		Name string
		Age  int
	}{"Alice", 30}

	fmt.Printf("default:  %v\n", person)
	fmt.Printf("with key: %+v\n", person)
	fmt.Printf("type:     %T\n", person)
}
