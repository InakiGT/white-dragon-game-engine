package application

import (
	"github.com/InakiGT/white-dragon-engine/src/internal/engine/domain"
	rendererApp "github.com/InakiGT/white-dragon-engine/src/internal/renderer/application"
	rendererDom "github.com/InakiGT/white-dragon-engine/src/internal/renderer/domain"
)

type GameLoop struct {
	clock         domain.Clock
	behaviors     []domain.Behavior
	renderer      *rendererApp.DrawSceneService
	entityManager *EntityManager
}

func NewGameLoop(r *rendererApp.DrawSceneService, e *EntityManager, c domain.Clock) *GameLoop {
	return &GameLoop{
		renderer:      r,
		clock:         c,
		entityManager: e,
		behaviors:     make([]domain.Behavior, 0),
	}
}

func (g *GameLoop) AddBehaviour(b domain.Behavior) {
	g.behaviors = append(g.behaviors, b)
}

func (g *GameLoop) Run(width, height int, title string) {
	g.renderer.Init(width, height, title)
	defer g.renderer.Close()

	for _, b := range g.behaviors {
		b.Start()
	}

	for !g.renderer.ShouldClose() {
		ctx := g.clock.GetFrameContext()

		for _, b := range g.behaviors {
			b.Update(ctx)
		}

		g.entityManager.Flush()

		renderables := make([]*rendererDom.Rect, 0)

		for _, v := range g.entityManager.Registry.GetAll() {
			comp, ok := v.Components[domain.ComponentRenderer]
			if !ok {
				continue
			}
			renderable := comp.(*rendererDom.Rect)
			renderable.Position.X = v.Transform.X
			renderable.Position.Y = v.Transform.Y
			renderables = append(renderables, renderable)
		}
		g.renderer.Draw(renderables)
	}
}
