package handler

import (
	"example/clean-arch/model"

	"github.com/gin-gonic/gin"
)

type Handler struct{
	UserService model.UserService
}

type Config struct {
	Router *gin.Engine
	UserService model.UserService
}

func NewHandler(c *Config) {
	handler := &Handler{
		UserService: c.UserService,
	}

	group := c.Router.Group("/api/account")

	group.GET("/me", handler.Me)
	group.POST("/signup", handler.Signup)
}
