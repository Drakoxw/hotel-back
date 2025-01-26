package routes

import (
	"hotel-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterHotelRoutes(r *gin.Engine) {
	hotels := r.Group("/hotels")
	{
		hotels.POST("/", controllers.CreateHotel)
		hotels.GET("/", controllers.GetHotels)
		hotels.PUT("/:id", controllers.UpdateHotel)
		hotels.DELETE("/:id", controllers.DeleteHotel)
	}
}
