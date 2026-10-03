package scenes

import (
	"github.com/MarcelArt/tile-system/internal/entities"
	"github.com/MarcelArt/tile-system/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const World = "world"

type WorldScene struct {
	tileSystem *entities.TileSystem
}

func NewWorldScene() *WorldScene {
	tileSystem := entities.NewTileSystem(
		entities.WithOffset(rl.NewVector2(4, 4)),
		entities.WithTileSize(32),
		entities.WithWidthAndHeight(8, 8),
	)
	tileSystem.Generate()

	return &WorldScene{
		tileSystem: tileSystem,
	}
}

// Draw implements [engine.IScene].
func (s *WorldScene) Draw() {
	rl.ClearBackground(rl.Black)
	s.tileSystem.Draw()
}

// GetID implements [engine.IScene].
func (s *WorldScene) GetID() string {
	return World
}

// Update implements [engine.IScene].
func (s *WorldScene) Update() engine.SceneResult {
	var result engine.SceneResult

	dt := rl.GetFrameTime()
	s.tileSystem.Update(dt)

	return result
}

var _ engine.IScene = &WorldScene{}
