package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"user/server/config"
	"user/server/models"
)

// GetUser
// Get the first user on the database
// Returns
//   - models.User: Technical User
func GetUser(c *gin.Context) {
	var user models.User
	result := config.DB.First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user.ToDto()})
}

// GetAllUsers
// Get all the users from the database and map to DTO
func GetAllUsers(c *gin.Context) {
	var allUsers []models.User
	result := config.DB.Find(&allUsers)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	allUsersDto := make([]models.UserDto, len(allUsers))
	for i, v := range allUsers {
		allUsersDto[i] = v.ToDto()
	}
	c.JSON(http.StatusOK, allUsersDto)
}

// SearchForUser
// Search for technical users based on a query provided by the user.
//
// Parameters:
//   - gin.Context: c
//
// Returns:
//   - User: The first result
func SearchForUser(c *gin.Context) {
	query := c.Query("query")
	query = strings.ToLower(query)
	// Handle empty search
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Query parameter is empty",
		})
		return
	}
	var technicalUsers []models.User
	query = "%" + query + "%"
	err := config.DB.Where("Email ILIKE ? OR Name ILIKE ?", query, query).Find(&technicalUsers).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, technicalUsers[0].ToDto())
	return
}
