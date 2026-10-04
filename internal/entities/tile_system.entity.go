package entities

import (
	"fmt"
	"log"
	"math"

	"github.com/MarcelArt/tile-system/internal/data"
	"github.com/MarcelArt/tile-system/pkg/array2d"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// TileSystem
type TileSystem struct {
	TileSize int32
	Offset   rl.Vector2
	Tiles    array2d.Array2D[*data.Tile]
}

type TileSystemOption func(*TileSystem)

func NewTileSystem(opts ...TileSystemOption) *TileSystem {
	e := &TileSystem{
		TileSize: 32,
		Offset:   rl.NewVector2(0, 0),
		Tiles:    array2d.New[*data.Tile](8, 8),
	}

	for _, opt := range opts {
		opt(e)
	}

	return e
}

func WithTileSize(tileSize int32) TileSystemOption {
	return func(ts *TileSystem) {
		ts.TileSize = tileSize
	}
}

func WithWidthAndHeight(width, height int32) TileSystemOption {
	return func(ts *TileSystem) {
		ts.Tiles = array2d.New[*data.Tile](int32(width), int32(height))
	}
}

func WithOffset(offset rl.Vector2) TileSystemOption {
	return func(ts *TileSystem) {
		ts.Offset = offset
	}
}

func (e *TileSystem) Generate() {
	tileColors := map[TileVariant]rl.Color{
		TileVacuum:    rl.NewColor(0, 0, 0, 255),       // Black color for vacuum
		TileDirt:      rl.NewColor(139, 69, 19, 255),   // Brown color for dirt
		TileSandstone: rl.NewColor(210, 180, 140, 255), // Tan color for sandstone
		TileCopperOre: rl.NewColor(184, 115, 51, 255),  // Copper color for copper ore
	}

	for x := range e.Tiles.GetW() {
		for y := range e.Tiles.GetH() {
			rng := rl.GetRandomValue(0, 3)
			tile := &data.Tile{
				Temperature: 20.0,
				Mass:        500.0,
				Variant:     TileVariant(rng),
				Color:       tileColors[TileVariant(rng)],
			}
			e.Tiles.Set(x, y, tile)
		}
	}
}

func (e *TileSystem) Draw() {
	width := e.Tiles.GetW()
	height := e.Tiles.GetH()
	offsetX, offsetY := e.TileToWorldPoint(int32(e.Offset.X), int32(e.Offset.Y))

	for x := range width {
		for y := range height {
			tile, err := e.Tiles.Get(x, y)
			if err != nil {
				continue
			}
			worldX, worldY := e.TileToWorldPoint(int32(x), int32(y))
			rl.DrawRectangle(worldX+offsetX, worldY+offsetY, e.TileSize, e.TileSize, tile.Color)
		}
	}
}

func (e *TileSystem) Update(dt float32) {
	e.debugTile()
}

func (e *TileSystem) TileToWorldPoint(x, y int32) (int32, int32) {
	return x * e.TileSize, y * e.TileSize
}

func (e *TileSystem) WorldToTileCoord(pos rl.Vector2) (int32, int32) {
	x := float64(pos.X / float32(e.TileSize))
	x = math.Floor(x)

	y := float64(pos.Y / float32(e.TileSize))
	y = math.Floor(y)

	return int32(x) - int32(e.Offset.X), int32(y) - int32(e.Offset.Y)
}

func (e *TileSystem) debugTile() {
	mousePos := rl.GetMousePosition()

	x, y := e.WorldToTileCoord(mousePos)

	tile, err := e.Tiles.Get(int32(x), int32(y))
	if err != nil {
		log.Println(err.Error())
		fmt.Println("===============================================")
		fmt.Printf("Mouse Pos	: (%d, %d)\n", x, y)
		fmt.Printf("Tile Coord	: (%f, %f)\n", mousePos.X, mousePos.Y)
		fmt.Println("===============================================")
		return
	}

	fmt.Println("===============================================")
	fmt.Printf("Tile		: %s\n", tileVariants[tile.Variant])
	fmt.Printf("Mass		: %f\n", tile.Mass)
	fmt.Printf("Temperature	: %f\n", tile.Temperature)
	fmt.Printf("Mouse Pos	: (%d, %d)\n", x, y)
	fmt.Printf("Tile Coord	: (%f, %f)\n", mousePos.X, mousePos.Y)
	fmt.Println("===============================================")
}

// End of TileSystem
