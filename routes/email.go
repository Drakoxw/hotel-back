package routes

import (
	"hotel-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterEmailsRoutes(r *gin.Engine) {
	reservations := r.Group("/mail")
	{
		reservations.POST("/new-reservation", controllers.SendEmailReservationCreated)
	}
}
