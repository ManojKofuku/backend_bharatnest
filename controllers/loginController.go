// Login authenticates a user and issues a JWT token
package controllers

import (
	"backend/initializers"
	"backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var body struct {
		EmailOrPhone bool   `json:"email_or_phone_number" binding:"required"` // true means email, false means phone
		Input        string `json:"input" binding:"required"`
		Password     string `json:"password" binding:"required"`
	}

	// Parse and validate the request body
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Find the user by email or phone based on the flag
	var user models.User
	var err error

	if body.EmailOrPhone {
		err = initializers.DB.Where("email = ?", body.Input).First(&user).Error
	} else {
		err = initializers.DB.Where("phone_number = ?", body.Input).First(&user).Error
	}

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email/phone or password"})
		return
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email/phone or password"})
		return
	}

	// Get user role
	var role models.Role
	if err := initializers.DB.First(&role, "id = ?", user.RoleID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve role information"})
		return
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"email":  user.Email,
		"phoneNumber": user.PhoneNumber,
		"role":   role.Role,
		"exp":    time.Now().Add(24 * 60 * 60 * time.Second).Unix(), // token expiry in 24 hours
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Respond with token and message
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   tokenString,
	})
}