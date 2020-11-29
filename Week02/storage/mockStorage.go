package storage

import (
	"database/sql"
	"homework/week02/users"
)

type MockUserStorage struct {
	userRepository users.Repository
}

func NewMockUserStorage() *MockUserStorage {
	return new(MockUserStorage)
}

func (us *MockUserStorage) Get(id string) (users.User, error) {
	var user users.User
	return user, sql.ErrNoRows
}
