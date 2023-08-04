package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {}

type Config struct {
  Router *gin.Engine
}

func NewHandler(c *Config) {
  handler := &Handler{}

  group := c.Router.Group("/api/account")

  group.GET("/me", handler.Me)
}

func (h *Handler) Me(c *gin.Context) {
  c.IndentedJSON(http.StatusOK, gin.H{
    "data": "it's me :)",
  })
}
