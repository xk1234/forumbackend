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

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
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
	router.Use(CORSMiddleware())
	protectedRoutes := router.Group("/api")
    protectedRoutes.Use(middleware.JWTAuthMiddleware())
    protectedRoutes.POST("/posts", controller.AddPost)
	protectedRoutes.PUT("/posts/:id", controller.UpdatePost)
	protectedRoutes.DELETE("/posts/:id", controller.DeletePost)
	protectedRoutes.GET("/posts", controller.GetUserPosts)
	protectedRoutes.POST("/comments/:pid", controller.AddComment)
	protectedRoutes.PUT("/comments/:id", controller.UpdateComment)
	protectedRoutes.DELETE("/comments/:id", controller.DeleteComment)

    publicRoutes := router.Group("/public")
    publicRoutes.POST("/signup", controller.SignUp)
    publicRoutes.POST("/login", controller.Login)
	publicRoutes.GET("/posts", controller.GetAllPosts)
	publicRoutes.GET("/comments/:pid", controller.GetComments)
	publicRoutes.GET("/post/:id", controller.GetPost)

    router.Run(":4000")
    fmt.Println("Server running on port 4000")
}