package main

import (

	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv" // Import the godotenv package
)

type Value struct {
	Currency int `json:"currency"`
}

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize the product and department database connections
	r := gin.Default()
	r.Static("/downloads", "./public")

	// Enable CORS with specific configurations
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // it should accept everything
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	fmt.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
