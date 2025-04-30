package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

// RegisterHotelRoutes registers all the hotel routes to the Gin router
func RegisterHotelRoutes(r *gin.Engine) {
	hotelRoutes := r.Group("/api/hotels")
	{
		// Public routes available to all users
		hotelRoutes.GET("", controllers.GetAllHotelsHTTP)
		hotelRoutes.GET("/:id", controllers.GetHotelByIDHTTP)

		// The following routes might require authorization in a real app
		hotelRoutes.POST("", controllers.CreateHotelHTTP)
		hotelRoutes.PUT("/:id", controllers.UpdateHotelHTTP)
		hotelRoutes.DELETE("/:id", controllers.DeleteHotelHTTP)
	}

	// Additional routes for room management could be added here
} 