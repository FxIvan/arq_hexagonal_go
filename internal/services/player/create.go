package services

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/arq_hexagonal/internal/domain"
)

func CreatePlayerService(player domain.Player) (id interface{}, err error) {
	///////////////////////////////////////////////////////////////////////////////////////////////////////
	//////////////////////////////////////// Consumision de servicio //////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////////////////////////////////
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
	collection := client.Database(os.Getenv("MONGO_NAME_DB")).Collection(os.Getenv("MONGO_NAME_COLLECTION"))
	insertResult, err := collection.InsertOne(ctx, player)
	if err != nil {
		log.Fatal(err)
	}
	return insertResult, nil
	///////////////////////////////////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////FIN//////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////////////////////////////////
}
