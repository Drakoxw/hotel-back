package controllers

import (
	"hotel-backend/database"
	"hotel-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Crear una nueva reserva
func CreateReservation(c *gin.Context) {
	var reservation models.Reservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&reservation)
	c.JSON(http.StatusCreated, reservation)
}

// Obtener la lista de reservas
func GetReservations(c *gin.Context) {
	var reservations []models.Reservation
	database.DB.Preload("EmergencyContact").Find(&reservations) // Preload carga el contacto de emergencia
	c.JSON(http.StatusOK, reservations)
}

// Obtener los detalles de una reserva específica
func GetReservationDetails(c *gin.Context) {
	var reservation models.Reservation
	id := c.Param("id")
	if err := database.DB.Preload("EmergencyContact").First(&reservation, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}
	c.JSON(http.StatusOK, reservation)
}

// Eliminar una reserva
func DeleteReservation(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Reservation{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Reservation deleted"})
}
