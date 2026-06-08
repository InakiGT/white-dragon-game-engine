package scripts

import (
	engDom "github.com/InakiGT/white-dragon-engine/src/internal/engine/domain"
	inputDom "github.com/InakiGT/white-dragon-engine/src/internal/input/domain"
)

var direction *float32

type PlayerMovement struct {
	Player *engDom.Entity
	Input  inputDom.InputProvider
	Speed  float32
}

func (s *PlayerMovement) Start() {
	direction = &s.Player.Transform.X
}

func (s *PlayerMovement) Update(ctx engDom.FrameContext) {
	*direction += (s.Speed * 3) * ctx.DeltaTime

	if s.Input.IsActionActive("WalkingRight") {
		direction = &s.Player.Transform.X
	}

	if s.Input.IsActionActive("WalkingDown") {
		direction = &s.Player.Transform.Y
	}
}
