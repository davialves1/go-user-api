package controllers

import (
	"example/server/config"
	"example/server/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func CreateTechnicalUser(c *gin.Context) {
	var techUser models.TechnicalUser
	err := c.ShouldBindJSON(&techUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	techUser.ID = uuid.New()
	techUser.CreatedAt = time.Now()
	result := config.DB.Create(&techUser)
	if result.Error != nil {
		fmt.Printf("Failed to create new technical user. %v", result.Error)
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"error":   http.StatusInternalServerError,
				"details": result.Error.Error(),
			})
		return
	}
	c.JSON(http.StatusOK, techUser)
}
