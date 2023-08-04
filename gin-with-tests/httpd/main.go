package main

import (
	"example/gin-with-tests/db"
	"example/gin-with-tests/httpd/handler"
	"example/gin-with-tests/models"
	"example/gin-with-tests/platform/newsfeed"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func setupDatabase() *gorm.DB {
	repo := db.Connect()
	repo.AutoMigrate(models.Post{})
	return repo
}

func hostname() string {
	return fmt.Sprintf("%v:%v", os.Getenv("HOST"), os.Getenv("PORT"))
}

func main() {
	loadEnv()
	database := setupDatabase()

	router := gin.Default()
	repo := newsfeed.New() // or DB, or any other shared dependency

	router.GET("/ping", handler.PingGet())

	router.GET("/feeds", handler.NewsfeedGetAll(repo))
	router.POST("/feeds", handler.NewsfeedPost(repo))

	router.GET("/posts", handler.PostsList(database))
	router.POST("/posts", handler.PostCreate(database))
	router.DELETE("/posts/:id", handler.PostDelete(database))

	handler.NewHandler(&handler.Config{
		Router: router,
	})

	router.Run(hostname())
}
