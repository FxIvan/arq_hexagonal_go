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

	//playerHandler es una variable de tipo player.Handler{} que dentro tiene esto
	/*
		type Handler struct {
			PlaterService ports.PlayerService
		}
		Este tiene a disposicion un componente de tipo ports.PlayerService
		Este ports.PlayerService tiene lo siguiente:
		type PlayerService interface {
			//CreatePlayerService(player domain.Player) (id interface{}, err error)
			Create(player domain.Player) (id interface{}, err error) //Definimos con solo Create ya que sabemos que hace referencia a Player
		}
		Este expone los servicios de nuestra aplicacion, en este caso tiene nuestro:
		Create(player domain.Player) (id interface{}, err error)
	*/

	playerRepo := playerMongo.Repository{
		Client: client,
	}

	playerSrv := playerService.Service{
		Repo: playerRepo,
	}

	playerHandler := player.Handler{
		PlayerService: playerSrv,
	}

	// playerHandler := player.Handler{
	// 	PlayerService: nil,
	// }

	ginEngine.POST("/players", playerHandler.CreatePlayer)

	log.Fatal(ginEngine.Run(":8001"))
}
