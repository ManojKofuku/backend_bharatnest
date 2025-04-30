package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterUserRoutes registers all the user/auth routes to the Gin router
func RegisterUserRoutes(r *gin.Engine) {
	userRoutes := r.Group("/api/users")
	{
		userRoutes.POST("", controllers.CreateUser) // admin-only endpoint
		userRoutes.GET("", controllers.GetUser) // admin-only endpoint
	}

	// Auth routes
	authRoutes := r.Group("/api/auth")
	{
		authRoutes.POST("/login", controllers.Login)
	}
}
