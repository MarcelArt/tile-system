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

const defaultSimStep = 0.2

// TileGrid
type TileGrid struct {
	Offset   rl.Vector2
	Tiles    array2d.Array2D[*data.Tile]
	SimStep  float32
	TileSize int32

	simStepCounter float32
}

type TileSystemOption func(*TileGrid)

func NewTileSystem(opts ...TileSystemOption) *TileGrid {
	e := &TileGrid{
		TileSize:       32,
		Offset:         rl.NewVector2(0, 0),
		Tiles:          array2d.New[*data.Tile](8, 8),
		SimStep:        defaultSimStep,
		simStepCounter: 0,
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

func WithSimStep(simStep float32) TileSystemOption {
	return func(tg *TileGrid) {
		tg.SimStep = simStep
	}
}

func (e *TileGrid) Generate() {
	w := e.Tiles.GetW()
	h := e.Tiles.GetH()

	for x := range w {
		for y := range h {
			blockID := rng.Int32(0, int32(data.BlockIDLength))
			temp := rng.Float32(20, 30)
			mass := rng.Float32(200, 600)

			if blockID == int32(data.BlockVacuum) {
				temp = -273
				mass = 0
			}

			tile := &data.Tile{
				Temperature: temp,
				Mass:        mass,
				Block:       data.Blocks[data.BlockID(blockID)],
				Color:       data.BlockColors[data.BlockID(blockID)],
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
	if e.simStepCounter >= e.SimStep {
		e.heatTransfer(dt)
		e.simStepCounter = 0
	}
	e.simStepCounter += dt

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
				e.exchangeHeat(x, y, x+1, y, dt)
			}
			if y+1 < h {
				e.exchangeHeat(x, y, x, y+1, dt)
			}
		}
	}
}

func (e *TileGrid) exchangeHeat(x1, y1, x2, y2 int32, dt float32) {
	current, err1 := e.Tiles.Get(x1, y1)
	neighbour, err2 := e.Tiles.Get(x2, y2)
	if err1 != nil || err2 != nil {
		return
	}

	if !e.canTransferHeat(current) || !e.canTransferHeat(neighbour) {
		return
	}

	deltaT := current.Temperature - neighbour.Temperature
	if deltaT == 0 {
		return
	}

	hot, cold := current, neighbour
	if deltaT < 0 {
		hot, cold = neighbour, current
		deltaT = -deltaT
	}

	k := math.Sqrt(float64(hot.Block.ThermalConductivity) * float64(cold.Block.ThermalConductivity))
	cHot := hot.Mass * hot.Block.SpecificHeatCapacity
	cCold := cold.Mass * cold.Block.SpecificHeatCapacity

	q := deltaT * dt * float32(k) * 1
	qMax := deltaT * (cHot * cCold) / (cHot + cCold)
	q = min(q, qMax)

	hot.Temperature -= q / cHot
	cold.Temperature += q / cCold
}

func (e *TileGrid) canTransferHeat(tile *data.Tile) bool {
	return tile.Block.SpecificHeatCapacity > 0 && tile.Block.ThermalConductivity > 0 && tile.Mass > 0
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
