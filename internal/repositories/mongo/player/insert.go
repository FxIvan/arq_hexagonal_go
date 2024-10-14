package player

import (
	"context"
	"log"
	"os"

	"github.com/arq_hexagonal/internal/domain"
)

func (r Repository) Insert(player domain.Player) (id interface{}, err error) {

	collection := r.Client.Database(os.Getenv("MONGO_NAME_DB")).Collection(os.Getenv("MONGO_NAME_COLLECTION"))
	insertResult, err := collection.InsertOne(context.Background(), player)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return insertResult.InsertedID, nil
}
