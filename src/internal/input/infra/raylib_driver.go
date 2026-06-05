package infra

import (
	"github.com/InakiGT/white-dragon-engine/src/internal/input/domain"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RaylibDriver struct {
	nativeKeys map[domain.Key]int32
}

func NewRaylibDriver() *RaylibDriver {
	return &RaylibDriver{
		nativeKeys: map[domain.Key]int32{
			domain.KeyA:         rl.KeyA,
			domain.KeyW:         rl.KeyW,
			domain.KeyS:         rl.KeyS,
			domain.KeyD:         rl.KeyD,
			domain.KeySpace:     rl.KeySpace,
			domain.KeyLeftShift: rl.KeyLeftShift,
		},
	}
}

func (d *RaylibDriver) IsKeyDown(key domain.Key) bool {
	rlKey, exists := d.nativeKeys[key]
	if !exists {
		return false
	}

	return rl.IsKeyDown(rlKey)
}
