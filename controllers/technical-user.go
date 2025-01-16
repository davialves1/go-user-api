package controllers

import (
	"example/server/config"
	"example/server/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// GetTechnicalUser
// Get the first user on the database
// Returns
//   - models.TechnicalUser: Technical User
func GetTechnicalUser(c *gin.Context) {
	var technicalUser models.TechnicalUser
	result := config.DB.First(&technicalUser)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": technicalUser})
}

// SearchForTechnicalUser
// Search for technical users based on a query provided by the user.
//
// Parameters:
//   - gin.Context: c
//
// Returns:
//   - TechnicalUser: The first result
func SearchForTechnicalUser(c *gin.Context) {
	query := c.Query("query")
	query = strings.ToLower(query)
	// Handle empty search
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Query parameter is empty",
		})
		return
	}
	var technicalUsers []models.TechnicalUser
	query = "%" + query + "%"
	err := config.DB.Where("Email ILIKE ? OR Name ILIKE ?", query, query).Find(&technicalUsers).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, technicalUsers[0])
	return
}
