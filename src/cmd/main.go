package main

import (
	"runtime"

	"github.com/InakiGT/white-dragon-engine/src/game/scripts"
	engineApp "github.com/InakiGT/white-dragon-engine/src/internal/engine/application"
	engineDom "github.com/InakiGT/white-dragon-engine/src/internal/engine/domain"
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

	inputManager.BindAction("WalkingLeft", inputDom.KeyA)
	inputManager.BindAction("WalkingRight", inputDom.KeyD)
	inputManager.BindAction("WalkingUp", inputDom.KeyW)
	inputManager.BindAction("WalkingDown", inputDom.KeyS)

	drawSceneService := rendererApp.NewDrawSceneService(rendererAdapter)

	snakeRenderer := rendererDom.NewRect(
		*rendererDom.NewVector2(0, 0),
		60,
		10,
	)
	pointRenderer := rendererDom.NewRect(*rendererDom.NewVector2(20, 20), 10, 10)
	snakeEntity := engineDom.NewEntity("0", nil)
	snakeEntity.AddComponent(snakeRenderer)

	pointEntity := engineDom.NewEntity("1", nil)
	pointEntity.AddComponent(pointRenderer)

	entities := map[engineDom.EntityID]*engineDom.Entity{}
	entities[pointEntity.Id] = pointEntity
	entities[snakeEntity.Id] = snakeEntity

	entityManager := engineApp.NewEntityManager(engineDom.NewEntityRegistry(entities))
	gameLoop := engineApp.NewGameLoop(drawSceneService, entityManager, clockAdapter)

	script := &scripts.PlayerMovement{
		Player: snakeRenderer,
		Input:  inputManager,
		Speed:  15.0,
	}

	gameLoop.AddBehaviour(script)

	gameLoop.Run(800, 450, "White Dragon Engine")
}
