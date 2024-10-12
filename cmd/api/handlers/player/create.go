/***********************************************
 * Esta sección corresponde a la capa de ENTITIES.
 * Las entities encapsulan la lógica y los datos
 * principales del dominio. En este caso, domain.Player
 * es una entidad que representa un jugador dentro
 * del sistema, conteniendo atributos esenciales
 * como el nombre, edad y tiempo de creación.
 *
 * Las entities en arquitectura hexagonal son independientes
 * de cualquier infraestructura externa (bases de datos, APIs).
 * Aquí la entidad domain.Player no tiene dependencias hacia
 * cómo o dónde se almacena, lo que facilita su
 * reutilización en diferentes capas y adaptadores.
 ***********************************************/
package player

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/arq_hexagonal/internal/domain"
)

func CreatePlayer(c *gin.Context) {
	//1. Traducir el request -> http
	//	 -validaciones
	//2. Consumir un servicio -> funcion que tendra la logica de nuestro handlers
	//3. Traducir el response -> http

	///////////////////////////////////////////////////
	////////////////Traducir Request///////////////////
	///////////////////////////////////////////////////
	var player domain.Player                    //////
	if err := c.BindJSON(&player); err != nil { //////
		c.JSON(400, gin.H{"error": err.Error()}) ////// Aqui es donde ocurre la validacion
		return                                   ////// En este caso de alguna manera estamos validando como esta en la linea 36
	}
	///////////////////////////////////////////////////
	///////////////////////////////////////////////////

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
	///////////////////////////////////////////////////////////////////////////////////////////////////////
	///////////////////////////////////////////////FIN//////////////////////////////////////////////////////
	///////////////////////////////////////////////////////////////////////////////////////////////////////

	//////////////////////////////////////////////////////////
	//////////////Traducir el response////////////////////////
	//////////////////////////////////////////////////////////
	c.JSON(200, gin.H{"player_id": insertResult.InsertedID}) //
	//////////////////////////////////////////////////////////
	/////////////////////////FIN//////////////////////////////
}
