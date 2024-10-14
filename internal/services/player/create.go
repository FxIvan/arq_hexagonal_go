package player

import (
	"time"

	"github.com/arq_hexagonal/internal/domain"
)

func (s Service) Create(player domain.Player) (id interface{}, err error) {
	player.CreationTime = time.Now().UTC()

	insertResult, err := s.Repo.Insert(player)
	if err != nil {
		return nil, err
	}

	return insertResult, nil
}
