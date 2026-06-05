package application

import "github.com/InakiGT/white-dragon-engine/src/internal/renderer/domain"

type DrawSceneService struct {
	renderer domain.Renderer
	elements []*domain.Rect
}

func NewDrawSceneService(renderer domain.Renderer) *DrawSceneService {
	return &DrawSceneService{
		renderer: renderer,
		elements: make([]*domain.Rect, 0),
	}
}

func (s *DrawSceneService) AddElements(element []*domain.Rect) {
	s.elements = append(s.elements, element...)
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

func (s *DrawSceneService) Draw() {
	s.renderer.BeginFrame()
	s.renderer.Clear(domain.ColorBlack)

	for _, el := range s.elements {
		s.renderer.DrawRect(*el, domain.ColorRed)
	}

	s.renderer.EndFrame()
}
