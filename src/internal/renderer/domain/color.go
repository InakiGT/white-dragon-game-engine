package domain

type Color struct {
	R, G, B, A uint8
}

func NewColor(R, G, B, A uint8) *Color {
	return &Color{R, G, B, A}
}

var (
	ColorBlack = Color{0, 0, 0, 255}
	ColorRed   = Color{255, 0, 0, 255}
	ColorWhite = Color{255, 255, 255, 255}
)
