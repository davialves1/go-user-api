package controllers

import (
	"fmt"
	"net/http"
	"user/server/config"
	"user/server/models"
	"user/server/utils"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	// Format request into type
	var request models.UserRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		fmt.Printf("Bad request: %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing data or bad request format"})
		return
	}

	// Hash Password
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Search for user in the database
	var user models.User
	user = user.New(request.Email, request.Name, hashedPassword)
	err = config.DB.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"error":   http.StatusInternalServerError,
				"details": err.Error(),
			})
		return
	}
	c.JSON(http.StatusOK, user.ToDto())
}
