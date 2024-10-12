package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/arq_hexagonal/cmd/api/handlers/player"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ginEngine := gin.Default()

	ginEngine.POST("/players", player.CreatePlayer)

	log.Fatal(ginEngine.Run(":8001"))
}
