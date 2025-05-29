package main

import (
	"log"
	"mini-loja/internal/routers"
	"mini-loja/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		if err := godotenv.Load("../../.env"); err != nil {
			log.Println("No .env file found")
		}
	}

	db := database.Connect()
	defer db.Close()

	router := gin.Default()
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		log.Fatal(err)
	}

	router.Use(gin.Logger(), gin.Recovery())

	// Register routes
	routers.RegisterRoutes(router, db)

	router.Run(":8080")
}
