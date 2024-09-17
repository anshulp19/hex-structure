package gamehdl

import "hex-structure/internal/core/domain"

type BodyCreate struct {
	Name  string `json:"string"`
	Size  uint   `json:"uint"`
	Bombs uint   `json:"uint"`
}

type ResponseCreate domain.Game

func BuildResponseCreate(model domain.Game) ResponseCreate {
	return ResponseCreate(model)
}
