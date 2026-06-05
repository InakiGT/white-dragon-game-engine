package main

import (
	"runtime"

	"github.com/InakiGT/white-dragon-engine/src/game/scripts"
	engineApp "github.com/InakiGT/white-dragon-engine/src/internal/engine/application"
	engineInfra "github.com/InakiGT/white-dragon-engine/src/internal/engine/infra"
	inputApp "github.com/InakiGT/white-dragon-engine/src/internal/input/application"
	inputDom "github.com/InakiGT/white-dragon-engine/src/internal/input/domain"
	inputInfra "github.com/InakiGT/white-dragon-engine/src/internal/input/infra"
	rendererApp "github.com/InakiGT/white-dragon-engine/src/internal/renderer/application"
	rendererDom "github.com/InakiGT/white-dragon-engine/src/internal/renderer/domain"
	rendererInfra "github.com/InakiGT/white-dragon-engine/src/internal/renderer/infra"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	rendererAdapter := rendererInfra.NewRaylibAdapter()
	clockAdapter := engineInfra.NewRaylibClock()

	raylibDriver := inputInfra.NewRaylibDriver()

	inputManager := inputApp.NewInputManager(raylibDriver)

	inputManager.BindAction("WalkingRight", inputDom.KeyD)
	inputManager.BindAction("Dash", inputDom.KeySpace, inputDom.KeyLeftShift)

	drawSceneService := rendererApp.NewDrawSceneService(rendererAdapter)

	gameLoop := engineApp.NewGameLoop(drawSceneService, clockAdapter)

	snakeSquare := rendererDom.NewRect(
		"0",
		*rendererDom.NewVector2(0, 0),
		60,
		10,
	)

	point := rendererDom.NewRect("1", *rendererDom.NewVector2(20, 20), 10, 10)

	drawSceneService.AddElements([]*rendererDom.Rect{snakeSquare, point})

	script := &scripts.PlayerMovement{
		Player: snakeSquare,
		Input:  inputManager,
		Speed:  15.0,
	}

	gameLoop.AddBehaviour(script)

	gameLoop.Run(800, 450, "White Dragon Engine")
}
