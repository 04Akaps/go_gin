package main

import (
	"context"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/jjimgo/go_gin/models"
	"github.com/jjimgo/go_gin/services"
	"github.com/jjimgo/go_gin/controllers"
	redis "github.com/redis/go-redis/v9"
)

var (
	server		*gin.Engine
	err			error
)

var (
	ctx			context.Context
	us			services.UserService
	uc			controllers.UserController
	userList	[]*models.User
)

var (
	rctx		context.Context
	rds 		services.RedisService
	rdsc		controllers.RedisController
)

func init() {
	// ---- user ---
	ctx = context.TODO()

	userOne := models.User{"testName1",  1}
	userTwo := models.User{"testName2",  2}

	userList := []*models.User{&userOne, &userTwo}
	us = services.NewUserService(userList,ctx)
	uc = controllers.New(us)

	// ---- redis ---
	rctx = context.Background()
	rgb := redis.NewClient(&redis.Options{
		Addr:	  "redis-16990.c289.us-west-1-2.ec2.cloud.redislabs.com:16990",
		Password: "MpE2LxEUJwWAYdfPV6NUX2S8ywN9Gcjg", 
		DB:		  0,  
	})

	rds = services.NewRedisService(rgb, rctx)
	rdsc = controllers.NewRedisController(rds)

	server = gin.Default()
}

func main() {
	basePath := server.Group("/")
	uc.RegisterUserRoutes(basePath)
	rdsc.RegisterRedisRouter(basePath)

	log.Fatal(server.Run(":8080"))
}