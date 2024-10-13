package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectClient(DB_URI string) (client *mongo.Client, err error) {
	// Connection to the DB, if the connection to the DB takes longer than 10 seconds
	// then the connection is cancelled
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to the MongoDB server
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(DB_URI))
	if err != nil {
		return nil, err
	}

	// Call the ping method yo verify that the connection has been stablished successfully
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}
