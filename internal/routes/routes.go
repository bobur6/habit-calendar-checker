package routes

import (
	"github.com/gin-gonic/gin"
	handlers2 "go-rest-project/internal/delivery/handlers"
	"go-rest-project/internal/middleware"
)

func healthHandler(c *gin.Context) {
	c.JSON(200, gin.H{"status": "ok"})
}

func SetupRoutes(
	r *gin.Engine,
	userHandler *handlers2.UserHandler,
	habitListHandler *handlers2.HabitListHandler,
	habitHandler *handlers2.HabitHandler,
	habitCheckHandler *handlers2.HabitCheckHandler,
) {

	// Глобальный обработчик 404
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "Not found"})
	})
	// Глобальный обработчик 405
	r.NoMethod(func(c *gin.Context) {
		c.JSON(405, gin.H{"error": "Method not allowed"})
	})

	r.GET("/health", healthHandler)

	r.POST("/api/auth/register", userHandler.Register)
	r.POST("/api/auth/login", userHandler.Login)

	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	{
		// User profile
		auth.GET("/users/profile", userHandler.GetProfile)
		auth.PUT("/users/profile", userHandler.UpdateProfile)
		auth.DELETE("/users/profile", userHandler.DeleteProfile)

		// Habit lists
		auth.POST("/habit-lists", habitListHandler.CreateHabitList)
		auth.GET("/habit-lists", habitListHandler.GetUserHabitLists)
		auth.GET("/habit-lists/:id", habitListHandler.GetHabitList)
		auth.PUT("/habit-lists/:id", habitListHandler.UpdateHabitList)
		auth.DELETE("/habit-lists/:id", habitListHandler.DeleteHabitList)

		// Habits
		auth.POST("/habits", habitHandler.CreateHabit)
		auth.GET("/habits", habitHandler.GetHabitsByList)
		auth.GET("/habits/:id", habitHandler.GetHabit)
		auth.PUT("/habits/:id", habitHandler.UpdateHabit)
		auth.DELETE("/habits/:id", habitHandler.DeleteHabit)

		// Habit checks
		auth.POST("/habit-checks", habitCheckHandler.CreateHabitCheck)
		auth.GET("/habit-checks", habitCheckHandler.GetHabitChecks)
		auth.GET("/habit-checks/:id", habitCheckHandler.GetHabitCheck)
		auth.GET("/habit-checks/date", habitCheckHandler.GetHabitCheckByDate)
		auth.PUT("/habit-checks/:id", habitCheckHandler.UpdateHabitCheck)
		auth.DELETE("/habit-checks/:id", habitCheckHandler.DeleteHabitCheck)
	}
}
