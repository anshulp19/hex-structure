package gameserv

import (
	"errors"
	"hex-structure/internal/core/domain"
	"hex-structure/internal/core/ports"
	"hex-structure/pkg/apperrors"
	"hex-structure/pkg/uidgen"
)

// implement individual port

type service struct {
	gamesRepository ports.GameRepositories
	uidGen          uidgen.UIDGen
}

func New(gamesRepository ports.GameRepositories, uidgen uidgen.UIDGen) *service {
	return &service{
		gamesRepository: gamesRepository,
		uidGen:          uidgen,
	}
}

/**
Get(id string) (domain.Game, error)
	Create(name string, size uint, bombs uint) (domain.Game, error)
	Reveal(id string, row uint, col uint) (domain.Game, error)
*/

func (srv *service) Get(id string) (domain.Game, error) {
	game, err := srv.gamesRepository.Get(id)
	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			return domain.Game{}, errors.New("game not dound")
		}
		return domain.Game{}, errors.New("get game from repository has failed")
	}

	game.Board = game.Board.HideBombs()

	return game, nil
}

func (srv *service) Create(name string, size uint, bombs uint) (domain.Game, error) {
	if bombs >= size*size {
		return domain.Game{}, errors.New(apperrors.InavlidInput.Error())
	}

	game := domain.NewGame(srv.uidGen.New(), name, size, bombs)

	if err := srv.gamesRepository.Save(game); err != nil {
		return domain.Game{}, errors.New(apperrors.Internal.Error())
	}

	game.Board = game.Board.HideBombs()

	return game, nil
}

func (srv *service) Reveal(id string, row uint, col uint) (domain.Game, error) {
	game, err := srv.gamesRepository.Get(id)
	if err != nil {
		return domain.Game{}, errors.New(apperrors.Internal.Error())
	}
	if !game.Board.IsValidPosition(row, col) {
		return domain.Game{}, errors.New(apperrors.InavlidInput.Error())
	}

	if game.IsOver() {
		return domain.Game{}, errors.New(apperrors.IllegalOperation.Error())
	}

	if game.Board.Contains(row, col, domain.CELL_BOMB) {
		game.State = domain.GAME_STATE_LOST
	} else {
		game.Board.Set(row, col, domain.CELL_REVEALED)

		hasEmptyCell := game.Board.HasEmptyCells()
		if !hasEmptyCell {
			game.State = domain.GAME_STATE_WON
		}
	}

	if err := srv.gamesRepository.Save(game); err != nil {
		return domain.Game{}, errors.New(apperrors.Internal.Error())
	}

	game.Board = game.Board.HideBombs()

	return game, nil

}
