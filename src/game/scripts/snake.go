package scripts

import (
	"fmt"

	engDom "github.com/InakiGT/white-dragon-engine/src/internal/engine/domain"
	inputDom "github.com/InakiGT/white-dragon-engine/src/internal/input/domain"
	renderDom "github.com/InakiGT/white-dragon-engine/src/internal/renderer/domain"
)

var direction *float32

type PlayerMovement struct {
	Player *renderDom.Rect
	Input  inputDom.InputProvider
	Speed  float32
}

func (s *PlayerMovement) Start() {
	direction = &s.Player.Position.X
}

func (s *PlayerMovement) Update(ctx engDom.FrameContext) {
	*direction += (s.Speed * 3) * ctx.DeltaTime

	if s.Input.IsActionActive("WalkingRight") {
		fmt.Println("HEMOS CAIDO VALE")
		direction = &s.Player.Position.X
	}
	if s.Input.IsActionActive("WalkingDown") {
		direction = &s.Player.Position.Y
	}
}
