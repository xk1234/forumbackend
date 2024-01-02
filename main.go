package main

import (
	"cvwoapi/controller"
	"cvwoapi/database"
	"cvwoapi/middleware"
	"cvwoapi/model"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var Database *gorm.DB

func main() {
    loadEnv()
    loadDatabase()
    runApp()
}

func loadEnv() {
    err := godotenv.Load(".env.local")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func loadDatabase() {
    database.Connect()
    database.Database.AutoMigrate(&model.User{})
    database.Database.AutoMigrate(&model.Post{})
	database.Database.AutoMigrate(&model.Comment{})
}

func runApp() {
    router := gin.Default()
	protR := router.Group("/api")
    protR.Use(middleware.JWTAuthMiddleware())
    protR.POST("/posts", controller.AddPost)
	protR.POST("/comments/:id", controller.AddComment)
	protR.GET("/posts", controller.GetUserPosts)
    pubR := router.Group("/public")
    pubR.POST("/signup", controller.SignUp)
    pubR.POST("/login", controller.Login)
	pubR.GET("/posts", controller.GetAllPosts)
	pubR.GET("/comments/:id", controller.GetComments)

    router.Run(":3000")
    fmt.Println("Server running on port 3000")
}