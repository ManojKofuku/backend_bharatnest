package main

import (
	"backend/initializers"
	"backend/models"
	"fmt"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Hotel{}, &models.Room{}, &models.Manager{})
	fmt.Println("✅ Migration completed successfully!")
}
