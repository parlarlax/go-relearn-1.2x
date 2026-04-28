package store

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Order struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Item   string `json:"item"`
	Qty    int    `json:"qty"`
}

type UserStore interface {
	Create(name, email string) (*User, error)
	GetByID(id int) (*User, error)
	List() ([]*User, error)
	Update(id int, name, email string) (*User, error)
	Delete(id int) error
}

type OrderStore interface {
	Create(userID int, item string, qty int) (*Order, error)
	ListByUser(userID int) ([]*Order, error)
	List() ([]*Order, error)
}
