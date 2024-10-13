package ports

import "github.com/arq_hexagonal/internal/domain"

/*
* Definimos los puertos de los servicios.
* Es decir aqui tendremos definidos todos nuestro metodos que incluye un servicio.
* En este caso Player
*
**/

type PlayerService interface {
	//CreatePlayerService(player domain.Player) (id interface{}, err error)
	Create(player domain.Player) (id interface{}, err error) //Definimos con solo Create ya que sabemos que hace referencia a Player
}

type PlayerRepository interface {
	Insert(player domain.Player) (id interface{}, err error)
}
