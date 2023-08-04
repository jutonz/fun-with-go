package controllers

import (
	"example/api-testin-v2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)
	c.IndentedJSON(http.StatusOK, gin.H{"data": books})
}

func GetBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")

	if err := models.DB.Where("id = ?", id).First(&book).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"data": "No matching book"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"data": book})
}

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

func CreateBook(c *gin.Context) {
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{
		Title:  input.Title,
		Author: input.Author,
	}
	models.DB.Create(&book)

	c.IndentedJSON(http.StatusCreated, gin.H{"data": book})
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func UpdateBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := models.DB.Where("id = ?", id).First(&book).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"data": "no matching book"})
		return
	}

	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"data": err.Error()})
		return
	}

	models.DB.Model(&book).Updates(input)
	c.IndentedJSON(http.StatusOK, gin.H{"data": book})
}

func DeleteBook(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	if err := models.DB.Where("id = ?", id).First(&book).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"data": "no matching book"})
		return
	}

	models.DB.Delete(&book)
	c.IndentedJSON(http.StatusOK, gin.H{"data": "deleted book"})
}
