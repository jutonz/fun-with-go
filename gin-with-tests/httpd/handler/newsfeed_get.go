package handler

import (
	"example/gin-with-tests/platform/newsfeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewsfeedGetAll(repo *newsfeed.Repo) gin.HandlerFunc {
  return func(c *gin.Context) {
    results := repo.GetAll()
    c.IndentedJSON(http.StatusOK, results)
  }
}
