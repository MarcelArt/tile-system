package entities

import (
	"fmt"
	"log"
	"math"

	"github.com/MarcelArt/tile-system/internal/data"
	"github.com/MarcelArt/tile-system/pkg/array2d"
	"github.com/MarcelArt/tile-system/pkg/rng"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// TileGrid
type TileGrid struct {
	TileSize int32
	Offset   rl.Vector2
	Tiles    array2d.Array2D[*data.Tile]
}

type TileSystemOption func(*TileGrid)

func NewTileSystem(opts ...TileSystemOption) *TileGrid {
	e := &TileGrid{
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
	return func(ts *TileGrid) {
		ts.TileSize = tileSize
	}
}

func WithWidthAndHeight(width, height int32) TileSystemOption {
	return func(ts *TileGrid) {
		ts.Tiles = array2d.New[*data.Tile](int32(width), int32(height))
	}
}

func WithOffset(offset rl.Vector2) TileSystemOption {
	return func(ts *TileGrid) {
		ts.Offset = offset
	}
}

func (e *TileGrid) Generate() {
	w := e.Tiles.GetW()
	h := e.Tiles.GetH()

	for x := range w {
		for y := range h {
			blockRNG := rng.Int32(0, int32(data.BlockIDLength))
			tempRNG := rng.Float32(20, 30)
			massRNG := rng.Float32(200, 600)
			tile := &data.Tile{
				Temperature: tempRNG,
				Mass:        massRNG,
				Block:       data.Blocks[data.BlockID(blockRNG)],
				Color:       data.BlockColors[data.BlockID(blockRNG)],
			}
			e.Tiles.Set(x, y, tile)
		}
	}
}

func (e *TileGrid) Draw() {
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

func (e *TileGrid) Update(dt float32) {
	e.debugTile()
}

func (e *TileGrid) TileToWorldPoint(x, y int32) (int32, int32) {
	return x * e.TileSize, y * e.TileSize
}

func (e *TileGrid) WorldToTileCoord(pos rl.Vector2) (int32, int32) {
	x := float64(pos.X / float32(e.TileSize))
	x = math.Floor(x)

	y := float64(pos.Y / float32(e.TileSize))
	y = math.Floor(y)

	return int32(x) - int32(e.Offset.X), int32(y) - int32(e.Offset.Y)
}

func (e *TileGrid) heatTransfer(dt float32) {
	w := e.Tiles.GetW()
	h := e.Tiles.GetH()

	for x := range w {
		for y := range h {
			if x+1 < w {

			}
			if y+1 < h {

			}
		}
	}
}

func (e *TileGrid) debugTile() {
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
	fmt.Printf("Block		: %s\n", tile.Block.Name)
	fmt.Printf("Mass		: %f\n", tile.Mass)
	fmt.Printf("Temperature	: %f\n", tile.Temperature)
	fmt.Printf("Mouse Pos	: (%d, %d)\n", x, y)
	fmt.Printf("Tile Coord	: (%f, %f)\n", mousePos.X, mousePos.Y)
	fmt.Println("===============================================")
}

// End of TileSystem
