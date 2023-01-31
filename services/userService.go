package services

import (
	userModel "github.com/jjimgo/go_gin/models"
)

type UserService interface {
	CreateUser(*userModel.User) error
	GetUser() ([]*userModel.User, error)
	DeleteUser(*string) error
}