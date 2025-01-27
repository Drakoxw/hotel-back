package controllers

import (
	"hotel-backend/models"
	"hotel-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendEmailReservationCreated(c *gin.Context) {
	var email models.EmailReservation
	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"success": false,
			"message": err.Error(),
		})
		return
	}

	services.SendEmailWithTemplate(email.MailTo, "RESERVACION REGISTRADA!", "templates/newReservation.html", email, []string{})

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"success": true,
		"message": "Correo Enviado!",
	})
}
