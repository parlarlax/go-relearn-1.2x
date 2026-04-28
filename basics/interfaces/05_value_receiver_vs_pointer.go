package interfaces

import "fmt"

type Counter struct {
	count int
}

func (c Counter) Value() int {
	return c.count
}

func (c *Counter) Increment() {
	c.count++
}

type Valuer interface {
	Value() int
}

type Incrementer interface {
	Increment()
}

func ExampleReceiverTypes() {
	c := Counter{count: 0}

	var v Valuer = c
	fmt.Println("value receiver:", v.Value())

	c.Increment()

	var i Incrementer = &c
	i.Increment()
	fmt.Println("after pointer calls:", c.Value())

	var v2 Valuer = &c
	fmt.Println("pointer satisfies value interface:", v2.Value())
}
