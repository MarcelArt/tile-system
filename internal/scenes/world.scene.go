package scenes

import (
	"github.com/MarcelArt/tile-system/internal/entities"
	"github.com/MarcelArt/tile-system/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const World = "world"

type WorldScene struct {
	tileGrid *entities.TileGrid
}

func NewWorldScene() *WorldScene {
	tileSystem := entities.NewTileSystem(
		entities.WithOffset(rl.NewVector2(4, 4)),
		entities.WithTileSize(32),
		entities.WithWidthAndHeight(8, 8),
	)
	tileSystem.Generate()

	return &WorldScene{
		tileGrid: tileSystem,
	}
}

// Draw implements [engine.IScene].
func (s *WorldScene) Draw() {
	rl.ClearBackground(rl.Black)
	s.tileGrid.Draw()
}

// GetID implements [engine.IScene].
func (s *WorldScene) GetID() string {
	return World
}

// Update implements [engine.IScene].
func (s *WorldScene) Update() engine.SceneResult {
	var result engine.SceneResult

	dt := rl.GetFrameTime()
	s.tileGrid.Update(dt)

	return result
}

var _ engine.IScene = &WorldScene{}
