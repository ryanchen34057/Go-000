package biz

import "fmt"

// User
type User struct {
	ID    string
	Email string
	Name  string
}

// UserRepo
type UserRepo interface {
	CreateUser(*User)
	GetUserByID(id string) (*User, error)
	EmailExists(email string) bool
}

// UserService
type UserService struct {
	repo UserRepo
}

// NewUserService
func NewUserService(repo UserRepo) *UserService {
	return &UserService{repo: repo}
}

// Create user
func (us *UserService) CreateUser(user *User) error {
	if us.repo.EmailExists(user.Email) {
		return fmt.Errorf("Email already exists")
	}
	us.repo.CreateUser(user)
	return nil
}

// Get user
func (us *UserService) GetUserByID(id string) (*User, error) {
	user, err := us.repo.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
