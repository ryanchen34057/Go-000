package data

import "Go-000/Week04/internal/biz"

var _ biz.UserRepo = (biz.UserRepo)(nil)

// NewUserRepo
func NewUserRepo() biz.UserRepo {
	return new(userRepo)
}

type userRepo struct{}

// Create User
func (ur *userRepo) CreateUser(user *biz.User) {
	return
}

// Get user by id
func (ur *userRepo) GetUserByID(id string) (*biz.User, error) {
	return nil, nil
}

// Check email exists
func (ur *userRepo) EmailExists(email string) bool {
	return false
}
