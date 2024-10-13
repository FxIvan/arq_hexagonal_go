package player

import (
	"github.com/arq_hexagonal/internal/ports"
)

type Service struct {
	Repo ports.PlayerRepository
}
