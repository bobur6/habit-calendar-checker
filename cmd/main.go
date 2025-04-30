package main

import (
	handlers2 "go-rest-project/internal/delivery/handlers"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go-rest-project/internal/db"
	"go-rest-project/internal/repository"
	"go-rest-project/internal/routes"
	"go-rest-project/internal/services"
)

func main() {
	// Initialize database
	db.InitDB()

	// Initialize repositories
	userRepo := repository.NewUserRepository(db.DB)
	habitListRepo := repository.NewHabitListRepository(db.DB)
	habitRepo := repository.NewHabitRepository(db.DB)
	habitCheckRepo := repository.NewHabitCheckRepository(db.DB)

	// Initialize services
	userService := services.NewUserService(userRepo)
	habitListService := services.NewHabitListService(habitListRepo)
	habitService := services.NewHabitService(habitRepo, habitListRepo)
	habitCheckService := services.NewHabitCheckService(habitCheckRepo)

	// Initialize handlers
	userHandler := handlers2.NewUserHandler(userService)
	habitListHandler := handlers2.NewHabitListHandler(habitListService)
	habitHandler := handlers2.NewHabitHandler(habitService)
	habitCheckHandler := handlers2.NewHabitCheckHandler(habitCheckService)

	// Setup Gin router
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:3002"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	routes.SetupRoutes(r, userHandler, habitListHandler, habitHandler, habitCheckHandler)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	err := r.Run(":" + port)
	if err != nil {
		log.Fatal(err)
	}
}
