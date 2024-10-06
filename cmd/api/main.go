package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Player struct {
	Name         string    `json:"name"	binding:"required"`
	Age          int       `json:"age"		binding:"required"`
	CreationTime time.Time `json:"creation_time"`
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ginEngine := gin.Default()

	ginEngine.POST("/players", func(c *gin.Context) {
		var player Player
		if err := c.BindJSON(&player); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		player.CreationTime = time.Now().UTC()

		// Connection to the DB, if the connection to the DB takes longer than 10 seconds
		// then the connection is cancelled
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Connect to the MongoDB server
		client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
		if err != nil {
			log.Fatal(err)
		}

		// Call the ping method yo verify that the connection has been stablished successfully
		err = client.Ping(ctx, nil)
		if err != nil {
			log.Fatal(err)
		}

		// Conectamos a una base de datos("go-l") y obtenemos una coleccion("players")
		collection := client.Database("tangowallet").Collection("players")
		insertResult, err := collection.InsertOne(ctx, player)
		if err != nil {
			log.Fatal(err)
		}

		c.JSON(200, gin.H{"player_id": insertResult.InsertedID})
	})

	log.Fatal(ginEngine.Run(":8001"))
}
