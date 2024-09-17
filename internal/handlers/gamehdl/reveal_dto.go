package gamehdl

import "hex-structure/internal/core/domain"

type BodyRevealCell struct {
	Row  uint `json:"uint"`
	Cell uint `json:uint`
}

type ResponseRevealCell domain.Game

func BuildResponseRevealCell(model domain.Game) ResponseRevealCell {
	return ResponseRevealCell(model)
}
