package player

import "github.com/arq_hexagonal/internal/ports"

type Handler struct {
	PlayerService ports.PlayerService
}
