package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/arq_hexagonal/cmd/api/handlers/player"
	"github.com/arq_hexagonal/internal/repositories/mongo"
	playerMongo "github.com/arq_hexagonal/internal/repositories/mongo/player"
	playerService "github.com/arq_hexagonal/internal/services/player"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ginEngine := gin.Default()

	client, err := mongo.ConnectClient(os.Getenv("MONGO_URI"))
	if err != nil {
		log.Fatal(err.Error())
	}

	playerRepo := playerMongo.Repository{
		Client: client,
	}

	playerSrv := playerService.Service{
		Repo: playerRepo,
	}

	playerHandler := player.Handler{
		PlayerService: playerSrv,
	}

	ginEngine.POST("/players", playerHandler.CreatePlayer)

	log.Fatal(ginEngine.Run(":8001"))
}
