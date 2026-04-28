package user

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func New(name, email string) *User {
	return &User{Name: name, Email: email}
}
