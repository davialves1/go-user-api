package controllers

import (
	"net/http"
	"user/server/config"
	"user/server/models"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authentication failed."})
		return
	}

	var user models.User
	result := config.DB.Raw("SELECT id, name, email, hashed_password FROM users WHERE email = ?", request.Email).Scan(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authentication failed."})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(request.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authentication failed."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"login": "success!", "user": user.ToDto()})

}
