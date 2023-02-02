package controllers

import (
	"net/http"
	"github.com/jjimgo/go_gin/services"
	"github.com/gin-gonic/gin"
)

type RedisController struct {
	RedisService services.RedisService
}

func NewRedisController(redis services.RedisService) RedisController {
	return RedisController {
		RedisService : redis,
	}
}

type RedisReq struct {
	Key	string `json:"key"`
	Value string `json:"value"`
}

type RedisUriReq struct {
	Key string `uri:"key" binding:"required"`
}

func (redis *RedisController) SetRedis(ctx *gin.Context) {
	var req RedisReq

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	err := redis.RedisService.SetRedis(req.Key, req.Value)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message " : err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "success")
}

func (redis *RedisController) GetRedis(ctx *gin.Context) {
	var req RedisUriReq

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	result, err := redis.RedisService.GetRedis(req.Key)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message " : err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (redis *RedisController) DeleteRedis(ctx *gin.Context) {
	var req RedisUriReq

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	 err := redis.RedisService.DeleteRedis(req.Key)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message " : err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "success")
}


func (redis *RedisController) RegisterRedisRouter(rg *gin.RouterGroup) {
	redisRoute := rg.Group("/redis")

	redisRoute.POST("/create", redis.SetRedis)
	redisRoute.GET("/get/:key", redis.GetRedis)
	redisRoute.DELETE("/delete/:key", redis.DeleteRedis)
}