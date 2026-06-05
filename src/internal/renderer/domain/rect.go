package domain

import "github.com/InakiGT/white-dragon-engine/src/internal/engine/domain"

type Rect struct {
	Entity   domain.Entity
	Position Vector2
	With     float32
	Height   float32
}

func NewRect(id string, position Vector2, with, height float32) *Rect {
	return &Rect{*domain.NewEntity(id), position, with, height}
}
