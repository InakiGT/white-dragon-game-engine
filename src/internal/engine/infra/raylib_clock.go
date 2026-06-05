package infra

import (
	"time"

	"github.com/InakiGT/white-dragon-engine/src/internal/engine/domain"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RaylibClock struct{}

func NewRaylibClock() *RaylibClock {
	return &RaylibClock{}
}

func (c *RaylibClock) GetFrameContext() domain.FrameContext {
	return *domain.NewFramwContext(
		rl.GetFrameTime(),
		time.Duration(rl.GetTime()*float64(time.Second)),
	)
}
