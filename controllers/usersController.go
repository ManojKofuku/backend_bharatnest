package controllers

import (
	"net/http"
	"os"
	"time"

	"backend/initializers"
	"backend/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)
func init() {
	initializers.LoadEnvVariables()
}
// Secret key for JWT (you should move this to a safer place like environment variables)
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func CreateUser(c *gin.Context) {
	var body struct {
		FullName    string `json:"fullName"`
		Email       string `json:"email" binding:"required,email"`
		Password    string `json:"password" binding:"required,min=8"` // Added password field with validation
		DateOfBirth string `json:"dateOfBirth"` // Format: dd/mm/yyyy
		Gender      string `json:"gender"`
		Role        string `json:"role" binding:"required"` // Make role required
	}

	// Parse the incoming JSON request body
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate the role field
	if body.Role == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Role cannot be empty"})
		return
	}

	// Parse date of birth
	dob, err := time.Parse("02/01/2006", body.DateOfBirth)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format (use dd/mm/yyyy)"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Generate UUID for the new user
	newUUID := uuid.New().String()

	// Check if the role exists, otherwise create it
	var role models.Role
	if err := initializers.DB.Where("role = ?", body.Role).First(&role).Error; err != nil {
		// Role does not exist, create a new role
		role = models.Role{
			ID:          uuid.New().String(),
			Role:        body.Role, // Role passed from request body
			Description: body.Role + " role",
		}
		// Save the new role to the database
		if err := initializers.DB.Create(&role).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create role"})
			return
		}
	}

	// Create a user with the valid role_id
	user := models.User{
		ID:             newUUID,
		FullName:       body.FullName,
		Email:          body.Email,
		HashedPassword: string(hashedPassword), // Store the hashed password
		DateOfBirth:    dob.Format("02/01/2006"),
		IsInvited:      true,
		RoleID:         role.ID, // Assign the valid role ID
	}

	// Save the user to the database
	if err := initializers.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Create JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"email":  user.Email,
		"role":   role.Role,
		"exp":    time.Now().Add(24 * 60 * 60 * time.Second).Unix(), // token expiry in 24 hours
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Respond with the success message and token
	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"token":   tokenString,
	})
}

func GetUser(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	var user models.User
	if err := initializers.DB.Where("email = ?", email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User found",
		"user":    user,
	})
}


