package routes

import (
	"hotel-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterReservationRoutes(r *gin.Engine) {
	reservations := r.Group("/reservations") // Grupo de rutas con prefijo "/reservations"
	{
		reservations.POST("/", controllers.CreateReservation)       // Crear una reserva
		reservations.GET("/", controllers.GetReservations)          // Obtener todas las reservas
		reservations.GET("/:id", controllers.GetReservationDetails) // Obtener detalles de una reserva
	}
}
