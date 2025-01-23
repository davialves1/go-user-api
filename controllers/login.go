package controllers

import (
	"fmt"
	"net/http"
	"user/server/config"
	"user/server/models"
	"user/server/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var request struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		fmt.Println("Failed to bind user to request struct")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authentication failed."})
		return
	}

	// Get user by email
	var user models.User
	result := config.DB.Raw("SELECT id, name, email, hashed_password FROM users WHERE email = ?", request.Email).Scan(&user)
	if result.Error != nil {
		fmt.Println("Failed to retrieve user from database")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authentication failed."})
		return
	}

	// Validate password
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(request.Password))
	if err != nil {
		fmt.Println("Failed to validate password")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authentication failed."})
		return
	}

	// Generate jwt token
	token, err := utils.GenerateToken(user.Email, user.ID.String())
	if err != nil {
		fmt.Printf("\nFailed to generate JWT token:\n%v\n\n", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authentication failed."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"login": "success!", "token": token})

}
