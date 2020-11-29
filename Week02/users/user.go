package users

type User struct {
	ID   string
	Name string
}

type Repository interface {
	Get(id string) (User, error)
}
