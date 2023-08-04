package handler

import (
	"example/gin-with-tests/platform/newsfeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

type input struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

func NewsfeedPost(repo newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		body := input{}
		c.Bind(&body)

		item := newsfeed.Item{
			Title: body.Title,
			Post:  body.Post,
		}

		repo.Add(item)

		c.Status(http.StatusNoContent)
	}
}
