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
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/arq_hexagonal/internal/domain"
)

func (h Handler) CreatePlayer(c *gin.Context) {
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

	insertedId, err := h.PlayerService.Create(player)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "oops !"})
	}
	//////////////////////////////////////////////////////////
	//////////////Traducir el response////////////////////////
	//////////////////////////////////////////////////////////
	c.JSON(200, gin.H{"player_id": insertedId}) //
	//////////////////////////////////////////////////////////
	/////////////////////////FIN//////////////////////////////
}
