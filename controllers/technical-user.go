package controllers

import (
	"example/server/config"
	"example/server/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetTechnicalUser(c *gin.Context) {
	var technicalUser models.TechnicalUser
	result := config.DB.First(&technicalUser)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": technicalUser})
}

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
	config.DB.Find(&technicalUsers)
	// Perform search
	for _, v := range technicalUsers {
		if searchTechnicalUser(query, &v) {
			c.JSON(http.StatusOK, gin.H{
				"data": v,
			})
			return
		}
	}

	// Returns null if it doesn't find anything
	c.JSON(http.StatusOK, gin.H{
		"data": nil,
	})
}

func searchTechnicalUser(query string, user *models.TechnicalUser) bool {
	return strings.Contains(strings.ToLower(user.Name), query) ||
		strings.Contains(strings.ToLower(user.Email), query) ||
		strings.Contains(strings.ToLower(user.Gid), query) ||
		strings.Contains(strings.ToLower(user.ID.String()), query)
}
