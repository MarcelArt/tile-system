package scenes

import (
	"fmt"

	"github.com/MarcelArt/tile-system/internal/data"
	"github.com/MarcelArt/tile-system/internal/entities"
	"github.com/MarcelArt/tile-system/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	World           = "world"
	defaultTickRate = 0.2
	cameraSpeed     = 600
)

type WorldScene struct {
	tileGrid    *entities.TileGrid
	tickCounter float32
	camera      rl.Camera2D
	player      *entities.Player
}

func NewWorldScene() *WorldScene {
	tileSystem := entities.NewTileSystem(
		// entities.WithOffset(rl.NewVector2(4, 4)),
		entities.WithTileSize(64),
		entities.WithWidthAndHeight(64, 44),
	)
	tileSystem.Generate()

	h := tileSystem.Tiles.GetH()
	x := tileSystem.Tiles.GetW() / 2
	var playerPos rl.Vector2
	for y := range h {
		tile, _ := tileSystem.Tiles.Get(x, y+1)
		if tile.Block.ID != data.BlockVacuum {
			playerX, playerY := tileSystem.TileToWorldPoint(x, y)
			playerPos.X = float32(playerX)
			playerPos.Y = float32(playerY)
			break
		}
	}

	screenWidth := rl.GetScreenWidth()
	screenHeight := rl.GetScreenHeight()

	return &WorldScene{
		tileGrid:    tileSystem,
		tickCounter: 0,
		camera: rl.NewCamera2D(
			rl.NewVector2(float32(screenWidth)/2, float32(screenHeight)/2),
			rl.Vector2Zero(),
			0,
			1,
		),
		player: entities.NewPlayer(
			entities.WithPlayerPosition(playerPos),
			entities.WithPlayerWidthAndHeight(64, 64),
		),
	}
}

// Draw implements [engine.IScene].
func (s *WorldScene) Draw() {
	rl.ClearBackground(rl.Black)
	rl.BeginMode2D(s.camera)

	s.tileGrid.Draw()
	s.player.Draw()
	s.drawTileTooltip()

	s.collisionDebug()

	rl.EndMode2D()
}

// GetID implements [engine.IScene].
func (s *WorldScene) GetID() string {
	return World
}

// Update implements [engine.IScene].
func (s *WorldScene) Update() engine.SceneResult {
	var result engine.SceneResult

	dt := rl.GetFrameTime()
	s.tileGrid.Update(dt, s.camera)
	s.player.Update(dt, s.tileGrid)
	s.moveCamera(dt)

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

func (s *WorldScene) Destroy() {
	s.tileGrid = nil
}

func (s *WorldScene) drawTileTooltip() {
	mousePos := engine.GetWorldMousePosition2D(s.camera)
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

func (s *WorldScene) moveCamera(dt float32) {
	if rl.IsKeyDown(rl.KeyA) {
		s.camera.Offset.X += cameraSpeed * dt
	}
	if rl.IsKeyDown(rl.KeyD) {
		s.camera.Offset.X -= cameraSpeed * dt
	}
	if rl.IsKeyDown(rl.KeyW) {
		s.camera.Offset.Y += cameraSpeed * dt
	}
	if rl.IsKeyDown(rl.KeyS) {
		s.camera.Offset.Y -= cameraSpeed * dt
	}
	s.clampCamera()
}

func (s *WorldScene) clampCamera() {
	worldW, worldH := s.tileGrid.WorldPixelSize()
	screenW := rl.GetScreenWidth()
	screenH := rl.GetScreenHeight()

	s.camera.Offset.X = rl.Clamp(s.camera.Offset.X, float32(screenW)-worldW, 0)
	s.camera.Offset.Y = rl.Clamp(s.camera.Offset.Y, float32(screenH)-worldH, 0)
}

// func (s *WorldScene) checkCollisions() {
// 	playerCol := s.player.CollisionRect()
// 	x0, y0 := s.tileGrid.WorldToTileCoord(rl.NewVector2(playerCol.X, playerCol.Y))
// 	x1, y1 := s.tileGrid.WorldToTileCoord(rl.NewVector2(playerCol.X+playerCol.Width, playerCol.Y+playerCol.Height))

// }

func (s *WorldScene) collisionDebug() {
	playerCol := s.player.CollisionRect()

	playerX, playerY := s.tileGrid.WorldToTileCoord(rl.NewVector2(playerCol.X, playerCol.Y))
	nwRect := s.tileGrid.CollisionRect(playerX-1, playerY-1)
	nRect := s.tileGrid.CollisionRect(playerX, playerY-1)
	neRect := s.tileGrid.CollisionRect(playerX+1, playerY-1)

	wRect := s.tileGrid.CollisionRect(playerX-1, playerY)
	eRect := s.tileGrid.CollisionRect(playerX+1, playerY)

	swRect := s.tileGrid.CollisionRect(playerX-1, playerY+1)
	sRect := s.tileGrid.CollisionRect(playerX, playerY+1)
	seRect := s.tileGrid.CollisionRect(playerX+1, playerY+1)

	// x0, y0 := s.tileGrid.WorldToTileCoord(rl.NewVector2(playerCol.X, playerCol.Y))
	// x1, y1 := s.tileGrid.WorldToTileCoord(rl.NewVector2(playerCol.X+playerCol.Width, playerCol.Y+playerCol.Height))

	// nwRect := s.tileGrid.CollisionRect(x0, y0)
	// seRect := s.tileGrid.CollisionRect(x1, y1)

	// log.Println("nwRect :>> ", nwRect)
	// log.Println("seRect :>> ", seRect)

	rl.DrawRectangleLinesEx(nwRect, 1, rl.Red)
	rl.DrawRectangleLinesEx(nRect, 1, rl.Red)
	rl.DrawRectangleLinesEx(neRect, 1, rl.Red)

	rl.DrawRectangleLinesEx(wRect, 1, rl.Red)
	rl.DrawRectangleLinesEx(eRect, 1, rl.Red)

	rl.DrawRectangleLinesEx(swRect, 1, rl.Red)
	rl.DrawRectangleLinesEx(sRect, 1, rl.Red)
	rl.DrawRectangleLinesEx(seRect, 1, rl.Red)

	s.player.CollisionDebug()
}

var _ engine.IScene = &WorldScene{}
