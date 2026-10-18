package data

import (
	"github.com/MarcelArt/tile-system/pkg/no"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type BlockState uint

const (
	StateGas BlockState = iota
	StateLiquid
	StateSolid
)

type Block struct {
	_ no.Copy

	ID                   BlockID
	Name                 string
	ThermalConductivity  float32
	SpecificHeatCapacity float32
	Modifier             float32
	Color                rl.Color
	BaseMass             float32
	State                BlockState
}

func (d *Block) TextureSource() rl.Rectangle {
	id := d.ID
	column := id % 2
	row := id / 2

	return rl.NewRectangle(float32(column)*627, float32(row)*627, 627, 627)
}

type BlockID = uint

const (
	BlockVacuum BlockID = iota
	BlockDirt
	BlockSandstone
	BlockCopperOre

	// Always put on last, only used for rng max value exclusivity
	BlockIDLength
)

var Blocks = map[BlockID]*Block{
	BlockVacuum: {
		ID:                   BlockVacuum,
		Name:                 "Vacuum",
		ThermalConductivity:  0,
		SpecificHeatCapacity: 0,
		Modifier:             0,
		Color:                rl.NewColor(0, 0, 0, 255),
		BaseMass:             0,
		State:                StateGas,
	},
	BlockDirt: {
		ID:                   BlockDirt,
		Name:                 "Dirt",
		ThermalConductivity:  2,
		SpecificHeatCapacity: 1.48,
		Modifier:             1,
		Color:                rl.NewColor(139, 69, 19, 255),
		BaseMass:             500,
		State:                StateSolid,
	},
	BlockSandstone: {
		ID:                   BlockSandstone,
		Name:                 "Sandstone",
		ThermalConductivity:  2.9,
		SpecificHeatCapacity: 0.8,
		Modifier:             1,
		Color:                rl.NewColor(210, 180, 140, 255),
		BaseMass:             800,
		State:                StateSolid,
	},
	BlockCopperOre: {
		ID:                   BlockCopperOre,
		Name:                 "Copper Ore",
		ThermalConductivity:  4.5,
		SpecificHeatCapacity: 0.386,
		Modifier:             1,
		Color:                rl.NewColor(184, 115, 51, 255),
		BaseMass:             400,
		State:                StateSolid,
	},
}

// var BlockColors = map[BlockID]rl.Color{
// 	BlockVacuum:    rl.NewColor(0, 0, 0, 255),       // Black color for vacuum
// 	BlockDirt:      rl.NewColor(139, 69, 19, 255),   // Brown color for dirt
// 	BlockSandstone: rl.NewColor(210, 180, 140, 255), // Tan color for sandstone
// 	BlockCopperOre: rl.NewColor(184, 115, 51, 255),  // Copper color for copper ore
// }
