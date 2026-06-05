package domain

type Transform struct{ X, Y float32 }

func NewTransform(x, y float32) *Transform {
	return &Transform{x, y}
}
