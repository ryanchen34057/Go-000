package user

import (
	"Go-000/Week04/internal/biz"
	context "context"

	"github.com/google/uuid"
)

type userServer struct {
	UnimplementedUserServiceServer
	us *biz.UserService
}

func NewUserServer(us *biz.UserService) *userServer {
	return &userServer{us: us}
}

func (u *userServer) CreateUser(context context.Context, userMessage *User) (*User, error) {
	user := &biz.User{}
	user.Email = userMessage.Email
	user.Name = userMessage.Name
	user.ID = uuid.New().String()

	err := u.us.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return userMessage, nil
}
func (u *userServer) GetUser(context context.Context, request *GetUserRequest) (*User, error) {
	user, err := u.us.GetUserByID(request.UserID)
	if err != nil {
		return nil, err
	}

	return &User{ID: user.ID, Name: user.Name, Email: user.Email}, nil

}
