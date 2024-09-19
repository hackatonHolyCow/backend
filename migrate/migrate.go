package main

import (
	"hackathon/backend/initializers"
	"hackathon/backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnecToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
