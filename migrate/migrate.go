package main

import (
	"github.com/tridims/go-crud/initializers"
	"github.com/tridims/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
