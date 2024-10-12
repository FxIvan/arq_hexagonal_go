package player

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/arq_hexagonal/internal/domain"
)

//////////////////////////////////////// Consumision de servicio //////////////////////////////////////

func (s Service) Create(player domain.Player) (id interface{}, err error) {
	/**
	* Responsabilidades que tiene que tener un servicio:
	* Set creation time
	* Save to repo
	* Responder con el id del recurso creado
	**/

	// Este es el primer paso que es -> Set creation time
	player.CreationTime = time.Now().UTC()
	// ================================ REPO
	insertResult, err := Save(player)
	if err != nil {
		return nil, err
	}
	// ================================ REPO
	// Este es el tercer paso -> Responder con el id del recurso creado
	return insertResult, nil
}

func Save(player domain.Player) (id interface{}, err error) {
	// ================================ REPO
	// Connection to the DB, if the connection to the DB takes longer than 10 seconds
	// then the connection is cancelled
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to the MongoDB server
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Call the ping method yo verify that the connection has been stablished successfully
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Conectamos a una base de datos("go-l") y obtenemos una coleccion("players")
	collection := client.Database(os.Getenv("MONGO_NAME_DB")).Collection(os.Getenv("MONGO_NAME_COLLECTION"))
	insertResult, err := collection.InsertOne(ctx, player)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return insertResult.InsertedID, nil
	// ================================ REPO
}
