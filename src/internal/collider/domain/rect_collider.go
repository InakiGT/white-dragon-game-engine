package domain

import (
	engDom "github.com/InakiGT/white-dragon-engine/src/internal/engine/domain"
)

type RectCollider struct {
	Width  float32
	Height float32
	entity *engDom.Entity
}

func (c *RectCollider) Intersects(other *RectCollider) bool {
	a := c.entity.Transform
	b := other.entity.Transform

	return a.X < b.X+other.Width &&
		a.X+c.Width > b.X &&
		a.Y < b.Y+other.Height &&
		a.Y+c.Height > b.Y
}

func (c *RectCollider) Type() engDom.ComponentType { return engDom.ComponentCollider }
func (c *RectCollider) OnDestroy()
