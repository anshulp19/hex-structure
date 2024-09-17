package gamesrepo

import (
	"encoding/json"
	"errors"
	"hex-structure/internal/core/domain"
)

type memkvs struct {
	kvs map[string][]byte
}

func NewMemKVS() *memkvs {
	return &memkvs{kvs: map[string][]byte{}}
}

func (repo *memkvs) Get(id string) (domain.Game, error) {
	if value, ok := repo.kvs[id]; ok {
		game := domain.Game{}
		err := json.Unmarshal(value, &game)
		if err != nil {
			return domain.Game{}, errors.New("failed to get value from kvs")
		}

		return game, nil
	}
	return domain.Game{}, errors.New("Game not found in kvs")
}

func (repo *memkvs) Save(game domain.Game) error {
	bytes, err := json.Marshal(game)
	if err != nil {
		return errors.New("Unable to encode the game")
	}
	repo.kvs[game.ID] = bytes
	return nil
}
