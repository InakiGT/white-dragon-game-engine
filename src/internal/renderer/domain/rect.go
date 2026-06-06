package domain

import "github.com/InakiGT/white-dragon-engine/src/internal/engine/domain"

type Rect struct {
	Position Vector2
	With     float32
	Height   float32
}

func NewRect(position Vector2, with, height float32) *Rect {
	return &Rect{position, with, height}
}

func (r *Rect) Type() domain.ComponentType { return domain.ComponentRenderer }
func (r *Rect) OnDestroy()                 {}
