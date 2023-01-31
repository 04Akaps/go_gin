package controllers

import (
	"github.com/jjimgo/go_gin/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func New(u services.UserService) UserController {
	return UserController {
		UserService : u
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H("message" : err.Error()))
		return
	}
	err : = uc.UserService.CreateUser(&user)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H("message " : err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Params("name")

	users, err : = uc.UserService.GetUser()

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) DeleteUser(ctx *gin.Context)  {
	username := ctx.Param("name")

		// need logic
	err : = uc.UserService.DeleteUser(username)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")
	userRoute.POST("/create", uc.CreateUser)
	userRoute.GET("/getUser/:name", uc.GetUser)
	userRoute.DELETE("/deleteUser/:name", uc.DeleteUser)
}