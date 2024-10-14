package ports

import "github.com/arq_hexagonal/internal/domain"

/*
* Definimos los puertos de los servicios.
* Es decir aqui tendremos definidos todos nuestro metodos que incluye un servicio.
**/

type PlayerService interface {
	Create(player domain.Player) (id interface{}, err error)
}

type PlayerRepository interface {
	Insert(player domain.Player) (id interface{}, err error)
}
