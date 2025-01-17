package controllers

import (
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
