package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes registers all the user/auth routes to the Gin router
func RegisterUserRoutes(r *gin.Engine) {
	// Public routes (no authentication required)
	publicRoutes := r.Group("/api")
	{
		// Auth routes
		publicRoutes.POST("/auth/login", controllers.Login)
		publicRoutes.POST("/users/register", controllers.CreateUser)
	}

	// Protected routes (authentication required)
	protectedRoutes := r.Group("/api")
	protectedRoutes.Use(middleware.AuthMiddleware())
	{
		// Current user routes
		protectedRoutes.GET("/users/me", controllers.GetCurrentUser)
	}
}
