package domain

type Vector2 struct {
	X, Y float32
}

func NewVector2(X, Y float32) *Vector2 {
	return &Vector2{X, Y}
}
