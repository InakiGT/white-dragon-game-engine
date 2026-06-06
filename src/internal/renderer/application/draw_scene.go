package application

import "github.com/InakiGT/white-dragon-engine/src/internal/renderer/domain"

type DrawSceneService struct {
	renderer domain.Renderer
}

func NewDrawSceneService(renderer domain.Renderer) *DrawSceneService {
	return &DrawSceneService{
		renderer: renderer,
	}
}

func (s *DrawSceneService) Init(width, height int, title string) {
	s.renderer.Init(width, height, title)
}

func (s *DrawSceneService) ShouldClose() bool {
	return s.renderer.ShouldClose()
}

func (s *DrawSceneService) Close() {
	s.renderer.Close()
}

func (s *DrawSceneService) Draw(elements []*domain.Rect) {
	s.renderer.BeginFrame()
	s.renderer.Clear(domain.ColorBlack)

	for _, el := range elements {
		s.renderer.DrawRect(*el, domain.ColorRed)
	}

	s.renderer.EndFrame()
}
