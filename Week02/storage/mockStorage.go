package storage

import (
	"Go-000/Week02/users"
	"database/sql"

	"github.com/pkg/errors"
)

type MockUserStorage struct {
	userRepository users.Repository
}

func NewMockUserStorage() *MockUserStorage {
	return new(MockUserStorage)
}

func (us *MockUserStorage) Get(id string) (users.User, error) {
	var user users.User
	if user, err := getFromDB(id); errors.Is(err, sql.ErrNoRows) {
		return user, errors.Wrapf(err, "User not found")
	}
	return user, nil
}

func getFromDB(id string) (users.User, error) {
	return users.User{}, sql.ErrNoRows
}
