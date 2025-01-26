package controllers

import (
	"hotel-backend/database"
	"hotel-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateHotel(c *gin.Context) {
	var hotel models.Hotel
	if err := c.ShouldBindJSON(&hotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&hotel)
	c.JSON(http.StatusCreated, hotel)
}

func GetHotels(c *gin.Context) {
	var hotels []models.Hotel
	database.DB.Preload("Rooms").Find(&hotels)
	c.JSON(http.StatusOK, hotels)
}

func UpdateHotel(c *gin.Context) {
	var hotel models.Hotel
	id := c.Param("id")
	if err := database.DB.First(&hotel, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Hotel not found"})
		return
	}
	if err := c.ShouldBindJSON(&hotel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&hotel)
	c.JSON(http.StatusOK, hotel)
}

func DeleteHotel(c *gin.Context) {
	id := c.Param("id")
	database.DB.Delete(&models.Hotel{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Hotel deleted"})
}
