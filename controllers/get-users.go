package controllers

import (
	"net/http"
	"strings"
	"user/server/config"
	"user/server/models"

	"github.com/gin-gonic/gin"
)

// GetUser
// Get the first user on the database
// Returns
//   - models.User: Technical User
func GetUserById(c *gin.Context) {
	var user models.User
	err := config.DB.Raw("Select * FROM users WHERE id = ?", c.Param("id")).Scan(&user).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user.ToDto()})
}

// GetAllUsers
// Get all the users from the database and map to DTO
func GetAllUsers(c *gin.Context) {
	// Fetch users from database
	var allUsers []models.User
	err := config.DB.Find(&allUsers).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Map users to DTO
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
		c.JSON(http.StatusOK, make([]string, 0))
		return
	}

	// Fetch from database
	var technicalUsers []models.User
	query = "%" + query + "%"
	err := config.DB.Where("Email ILIKE ? OR Name ILIKE ?", query, query).Find(&technicalUsers).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Map to DTOs
	userDtos := make([]models.UserDto, len(technicalUsers))
	for index, user := range technicalUsers {
		userDtos[index] = user.ToDto()
	}
	c.JSON(http.StatusOK, userDtos)
}
