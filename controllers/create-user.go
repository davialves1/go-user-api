package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"user/server/config"
	"user/server/models"
)

func CreateUser(c *gin.Context) {
	var request models.UserRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	user.New(request.Email, request.Name)
	result := config.DB.Create(&user)
	if result.Error != nil {
		fmt.Printf("Failed to create new technical user. %v", result.Error)
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"error":   http.StatusInternalServerError,
				"details": result.Error.Error(),
			})
		return
	}
	c.JSON(http.StatusOK, user.ToDto())
}
