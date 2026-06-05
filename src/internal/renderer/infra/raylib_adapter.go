package infra

import (
	"github.com/InakiGT/white-dragon-engine/src/internal/renderer/domain"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RaylibAdapter struct{}

func NewRaylibAdapter() *RaylibAdapter {
	return &RaylibAdapter{}
}

func (a *RaylibAdapter) Init(width, height int, title string) {
	rl.InitWindow(int32(width), int32(height), title)
	rl.SetTargetFPS(60)
}

func (a *RaylibAdapter) Clear(color domain.Color) {
	c := rl.NewColor(color.R, color.G, color.B, color.A)
	rl.ClearBackground(c)
}

func (a *RaylibAdapter) DrawRect(rect domain.Rect, color domain.Color) {
	c := rl.NewColor(color.R, color.G, color.B, color.A)
	rl.DrawRectangle(
		int32(rect.Position.X),
		int32(rect.Position.Y),
		int32(rect.With),
		int32(rect.Height),
		c,
	)
}

func (a *RaylibAdapter) Present() {
	rl.EndDrawing()
	rl.BeginDrawing()
}

func (a *RaylibAdapter) Close() {
	rl.CloseWindow()
}

func (a *RaylibAdapter) ShouldClose() bool {
	return rl.WindowShouldClose()
}

func (a *RaylibAdapter) BeginFrame() {
	rl.BeginDrawing()
}

func (a *RaylibAdapter) EndFrame() {
	rl.EndDrawing()
}
