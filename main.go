package main

import (
	"hotel-backend/database"
	"hotel-backend/routes"
	"hotel-backend/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	// // Migraciones
	// database.DB.AutoMigrate(
	// 	&models.Hotel{},
	// 	&models.Room{},
	// 	&models.Reservation{},
	// 	&models.EmergencyContact{},
	// )

	r := gin.Default()
	r.Use(utils.CORSMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  200,
			"success": true,
			"message": "Hola mundo",
		})
	})

	// Registrar rutas
	// routes.RegisterHotelRoutes(r)       // Rutas de hoteles
	// routes.RegisterRoomRoutes(r)        // Rutas de habitaciones
	// routes.RegisterReservationRoutes(r) // Rutas reservas
	routes.RegisterEmailsRoutes(r) // Rutas para email

	port := utils.GetPort()
	r.Run(port)
}
