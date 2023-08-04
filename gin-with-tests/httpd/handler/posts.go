package handler

import (
	"example/gin-with-tests/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PostsList(repo *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		posts := models.PostsList(repo)
		c.IndentedJSON(http.StatusOK, posts)
	}
}

type PostCreateInput struct {
	Title string `json:"title" binding:"required"`
	URL   string `json:"url" binding:"required"`
}

func PostCreate(repo *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
    var input PostCreateInput
    if err := c.ShouldBind(&input); err != nil {
      c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
      return
    }

		post := models.Post{
			Title: input.Title,
			URL: input.URL,
		}

		err := repo.Create(&post).Error
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{"data": post})
	}
}

func PostDelete(repo *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		err := repo.Delete(&models.Post{}, id).Error
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{
				"error": "failed to delete",
			})
			return
		}

		c.IndentedJSON(http.StatusOK, gin.H{
			"data": "deleted successfully",
		})
	}
}
