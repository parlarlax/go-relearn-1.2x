package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func (u User) Greet() string {
	return fmt.Sprintf("Hi, I'm %s, age %d", u.Name, u.Age)
}

func (u *User) Birthday() {
	u.Age++
}

type Admin struct {
	User
	Role string
}

func main() {
	fmt.Println("=== struct literal ===")
	u1 := User{Name: "Alice", Age: 25}
	u2 := User{"Bob", 30}
	fmt.Println(u1, u2)

	fmt.Println("\n=== constructor pattern ===")
	u3 := NewUser("Charlie", 35)
	fmt.Println(u3.Greet())

	fmt.Println("\n=== pointer receiver modifies ===")
	fmt.Println("before birthday:", u1.Age)
	u1.Birthday()
	fmt.Println("after birthday:", u1.Age)

	fmt.Println("\n=== struct embedding (composition) ===")
	admin := Admin{
		User: User{Name: "Boss", Age: 40},
		Role: "superadmin",
	}
	fmt.Println(admin.Name)
	fmt.Println(admin.Greet())
	fmt.Printf("admin: %+v\n", admin)
}
