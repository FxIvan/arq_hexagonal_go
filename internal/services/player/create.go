package player

import (
	"time"

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
	insertResult, err := s.Repo.Insert(player)
	if err != nil {
		return nil, err
	}
	// ================================ REPO
	// Este es el tercer paso -> Responder con el id del recurso creado
	return insertResult, nil
}
