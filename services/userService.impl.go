package services

import (
	userModel "github.com/jjimgo/go_gin/models"

	"context"
)

type UserServiceImpl struct {
	userList 	[]*userModel.User
	ctx			context.Context
}

func NewUserService(userList []*userModel.User, ctx context.Context) UserService {
	return &UserServiceImpl {
		userList : userList,
		ctx : ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *userModel.User) error {
	u.userList = append(u.userList, user);
	return nil
}

func (u *UserServiceImpl) GetUser() ([]*userModel.User, error) {
	return u.userList, nil
}

func (u *UserServiceImpl) DeleteUser(str *string) error {
		return nil
}