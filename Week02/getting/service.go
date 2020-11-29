package getting

import (
	"Go-000/Week02/users"
)

type Service interface {
	GetUser(id string) (users.User, error)
}

type service struct {
	uR users.Repository
}

func NewService(uR users.Repository) Service {
	return &service{uR}
}

func (s *service) GetUser(id string) (users.User, error) {
	return s.uR.Get(id)
}
