package domain

type Renderer interface {
	Close()
	// Sends the draw to the monitor
	Present()
	EndFrame()
	BeginFrame()
	ShouldClose() bool
	Clear(color Color)
	DrawRect(square Rect, color Color)
	Init(with, height int, title string)
}
