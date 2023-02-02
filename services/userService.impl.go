package services

import (
	userModel "github.com/jjimgo/go_gin/models"

	"context"
	"reflect"
	"errors"
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

func (u *UserServiceImpl) DeleteUser(str string) error {
	checkExisted := false

	for index, user := range u.userList {
		r := reflect.ValueOf(user)
		name := reflect.Indirect(r).FieldByName("Name") // get Name
		if(name.String() == str) {
			checkExisted = true
			u.userList = append(u.userList[:index], u.userList[index+1:]...)
		}
	}

	if (!checkExisted) {
		return errors.New("존재하지 않는 이름")
	}

	return nil
}