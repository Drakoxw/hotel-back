package controllers

import (
	"hotel-backend/database"
	"hotel-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Crear una nueva habitación
func CreateRoom(c *gin.Context) {
	var room models.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&room)
	c.JSON(http.StatusCreated, room)
}

// Obtener la lista de habitaciones
func GetRooms(c *gin.Context) {
	var rooms []models.Room
	database.DB.Find(&rooms)
	c.JSON(http.StatusOK, rooms)
}

// Actualizar una habitación
func UpdateRoom(c *gin.Context) {
	var room models.Room
	id := c.Param("id")
	if err := database.DB.First(&room, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&room)
	c.JSON(http.StatusOK, room)
}

// Eliminar una habitación
func DeleteRoom(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&models.Room{}, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Room deleted"})
}
