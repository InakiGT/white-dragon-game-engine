package application

import (
	"github.com/InakiGT/white-dragon-engine/src/internal/engine/domain"
	rendererApp "github.com/InakiGT/white-dragon-engine/src/internal/renderer/application"
)

type GameLoop struct {
	clock     domain.Clock
	behaviors []domain.Behavior
	renderer  *rendererApp.DrawSceneService
}

func NewGameLoop(r *rendererApp.DrawSceneService, c domain.Clock) *GameLoop {
	return &GameLoop{
		renderer:  r,
		clock:     c,
		behaviors: make([]domain.Behavior, 0),
	}
}

func (g *GameLoop) AddBehaviour(b domain.Behavior) {
	g.behaviors = append(g.behaviors, b)
}

func (g *GameLoop) Run(width, height int, title string) {
	g.renderer.Init(width, height, title)
	defer g.renderer.Close()

	for !g.renderer.ShouldClose() {
		ctx := g.clock.GetFrameContext()

		for _, b := range g.behaviors {
			b.Update(ctx)

		}

		g.renderer.Draw()
	}
}
