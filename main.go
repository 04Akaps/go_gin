package main

import (
	"context"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/jjimgo/go_gin/models"
	"github.com/jjimgo/go_gin/services"
	"github.com/jjimgo/go_gin/controllers"
)

var (
	ctx			context.Context
	server		*gin.Engine
	us			services.UserService
	uc			controllers.UserController
	err			error
	userList	[]*models.User
)


func init() {
	ctx = context.TODO()

	userOne := models.User{"testName1",  1}
	userTwo := models.User{"testName2",  2}

	userList := []*models.User{&userOne, &userTwo}
	us = services.NewUserService(userList,ctx)
	uc = controllers.New(us)

	server = gin.Default()
}

func main() {
	basePath := server.Group("/")
	uc.RegisterUserRoutes(basePath)

	log.Fatal(server.Run(":8080"))
}