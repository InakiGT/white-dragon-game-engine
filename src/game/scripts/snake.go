package scripts

import (
	engDom "github.com/InakiGT/white-dragon-engine/src/internal/engine/domain"
	inputDom "github.com/InakiGT/white-dragon-engine/src/internal/input/domain"
	renderDom "github.com/InakiGT/white-dragon-engine/src/internal/renderer/domain"
)

type PlayerMovement struct {
	Player *renderDom.Rect
	Input  inputDom.InputProvider
	Speed  float32
}

func (s *PlayerMovement) Update(ctx engDom.FrameContext) {
	s.Player.Position.X += (s.Speed * 3) * ctx.DeltaTime
	if s.Input.IsActionActive("Dash") {
	}

	if s.Input.IsActionActive("WalkingRight") {
		s.Player.Position.X += s.Speed * ctx.DeltaTime
	}
}
