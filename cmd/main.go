package main

import (
	"fmt"
	"user/server/config"
	"user/server/controllers"
	"user/server/middlewares"
	"user/server/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	server := gin.Default()
	setupServer(server)
	setupDataBase()
	loadEnvVariables()
	err := server.Run("localhost:8080")
	if err != nil {
		return
	}
}

func loadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load JWT secret key")
	}
}

func setupDataBase() {
	config.ConnectDatabase()
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		fmt.Printf("Failed to create tables")
	}
}

func setupServer(server *gin.Engine) {
	addCorsConfig(server)
	addRouters(server)
	addProxies(server)
}

func addCorsConfig(server *gin.Engine) {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:4200"}
	server.Use(cors.New(corsConfig))
}

func addRouters(server *gin.Engine) {
	authenticated := server.Group("/user")
	authenticated.Use(middlewares.Authenticate)
	authenticated.GET("/:id", controllers.GetUserById)
	authenticated.GET("/all", controllers.GetAllUsers)
	authenticated.GET("/search", controllers.SearchForUser)

	server.POST("/sign-up", controllers.CreateUser)
	server.POST("/login", controllers.Login)
}

func addProxies(server *gin.Engine) {
	err := server.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		fmt.Printf("Failed to set proxy %v\n", err)
		return
	}
}
