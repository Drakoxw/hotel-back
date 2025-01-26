package routes

import (
	"hotel-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoomRoutes(r *gin.Engine) {
	rooms := r.Group("/rooms") // Grupo de rutas con prefijo "/rooms"
	{
		rooms.POST("/", controllers.CreateRoom)      // Crear una habitación
		rooms.GET("/", controllers.GetRooms)         // Obtener todas las habitaciones
		rooms.PUT("/:id", controllers.UpdateRoom)    // Actualizar una habitación por ID
		rooms.DELETE("/:id", controllers.DeleteRoom) // Eliminar una habitación por ID
	}
}
