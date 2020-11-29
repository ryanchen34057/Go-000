package getting

import (
	"database/sql"
	"homework/week02/users"

	"github.com/pkg/errors"
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
	user, err := s.uR.Get(id)
	if errors.Is(err, sql.ErrNoRows) {
		return user, errors.Wrap(err, "User not found")
	}
	return user, nil
}
