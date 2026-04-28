package interfaces

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return fmt.Sprintf("%s says: Woof!", d.Name)
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return fmt.Sprintf("%s says: Meow!", c.Name)
}

func MakeItSpeak(s Speaker) string {
	return s.Speak()
}

func ExampleBasic() {
	d := Dog{Name: "Rex"}
	c := Cat{Name: "Whiskers"}

	fmt.Println(MakeItSpeak(d))
	fmt.Println(MakeItSpeak(c))
}
