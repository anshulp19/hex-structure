package ports

import "hex-structure/internal/core/domain"

// communicate with actors

type GameRepositories interface {
	Get(id string) (domain.Game, error)
	Save(domain.Game) error
}
