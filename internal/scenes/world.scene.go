package scenes

import (
	"fmt"

	"github.com/MarcelArt/tile-system/internal/entities"
	"github.com/MarcelArt/tile-system/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	World           = "world"
	defaultTickRate = 0.2
)

type WorldScene struct {
	tileGrid    *entities.TileGrid
	tickCounter float32
}

func NewWorldScene() *WorldScene {
	tileSystem := entities.NewTileSystem(
		entities.WithOffset(rl.NewVector2(4, 4)),
		entities.WithTileSize(32),
		entities.WithWidthAndHeight(8, 8),
	)
	tileSystem.Generate()

	return &WorldScene{
		tileGrid:    tileSystem,
		tickCounter: 0,
	}
}

// Draw implements [engine.IScene].
func (s *WorldScene) Draw() {
	rl.ClearBackground(rl.Black)
	s.tileGrid.Draw()
	s.drawTileTooltip()
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

	if s.tickCounter >= defaultTickRate {
		s.SimUpdate(dt)
		s.tickCounter = 0
	}
	s.tickCounter += dt

	return result
}

func (s *WorldScene) SimUpdate(dt float32) {
	s.tileGrid.SimUpdate(dt)
}

func (s *WorldScene) drawTileTooltip() {
	mousePos := rl.GetMousePosition()
	x, y := s.tileGrid.WorldToTileCoord(mousePos)

	tile, err := s.tileGrid.Tiles.Get(x, y)
	if err != nil || tile == nil || tile.Block == nil {
		return
	}

	const (
		fontSize     int32 = 20
		padding      int32 = 10
		lineSpacing  int32 = 6
		cursorOffset int32 = 16
	)
	lines := []string{
		fmt.Sprintf("Block: %s", tile.Block.Name),
		fmt.Sprintf("Mass: %.2f kg", tile.Mass),
		fmt.Sprintf("Temperature: %.2f C", tile.Temperature),
	}

	var textWidth int32
	for _, line := range lines {
		textWidth = max(textWidth, rl.MeasureText(line, fontSize))
	}
	width := textWidth + 2*padding
	height := int32(len(lines))*(fontSize+lineSpacing) - lineSpacing + 2*padding
	tooltipX := int32(mousePos.X) + cursorOffset
	tooltipY := int32(mousePos.Y) + cursorOffset

	// Flip the tooltip beside the cursor when it would cross a window edge.
	if tooltipX+width > int32(rl.GetScreenWidth()) {
		tooltipX = int32(mousePos.X) - width - cursorOffset
	}
	if tooltipY+height > int32(rl.GetScreenHeight()) {
		tooltipY = int32(mousePos.Y) - height - cursorOffset
	}
	tooltipX = max(0, tooltipX)
	tooltipY = max(0, tooltipY)

	rl.DrawRectangle(tooltipX, tooltipY, width, height, rl.NewColor(24, 28, 36, 245))
	rl.DrawRectangleLines(tooltipX, tooltipY, width, height, rl.Gray)
	for i, line := range lines {
		lineY := tooltipY + padding + int32(i)*(fontSize+lineSpacing)
		rl.DrawText(line, tooltipX+padding, lineY, fontSize, rl.RayWhite)
	}
}

var _ engine.IScene = &WorldScene{}
