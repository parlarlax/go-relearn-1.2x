package main

import "fmt"

func main() {
	var name string = "Alice"
	var city = "Bangkok"
	age := 25

	fmt.Println(name, city, age)

	var score int
	var active bool
	var label string
	fmt.Println(score, active, label)

	const AppName = "Go Relearn"
	const (
		StatusOK    = 200
		StatusError = 500
	)
	fmt.Println(AppName, StatusOK, StatusError)

	const (
		Sunday    = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Println("Sunday =", Sunday, "Saturday =", Saturday)

	fmt.Printf("type of name: %T\n", name)
	fmt.Printf("type of age: %T\n", age)
}
