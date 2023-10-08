package main

import (
	"github.com/AlmedinNasufi/go-crud/initializers"
	"github.com/AlmedinNasufi/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
