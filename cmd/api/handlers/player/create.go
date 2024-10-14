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
	var player domain.Player
	if err := c.BindJSON(&player); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	insertedId, err := h.PlayerService.Create(player)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "oops !"})
	}

	c.JSON(200, gin.H{"player_id": insertedId}) //
}
