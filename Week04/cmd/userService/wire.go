package main

import (
	"Go-000/Week04/internal/biz"
	"Go-000/Week04/internal/data"

	"github.com/google/wire"
)

func InitializeUserService() (biz.UserService, error) {
	wire.Build(biz.NewUserService, data.NewUserRepo)
	return biz.UserService{}, nil
}
