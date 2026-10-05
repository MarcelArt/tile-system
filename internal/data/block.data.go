package data

import (
	"github.com/MarcelArt/tile-system/pkg/no"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Block struct {
	_                    no.Copy
	Name                 string
	ThermalConductivity  float32
	SpecificHeatCapacity float32
	Modifier             float32
}

type BlockID uint

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
		Name:                 "Vacuum",
		ThermalConductivity:  0,
		SpecificHeatCapacity: 0,
		Modifier:             0,
	},
	BlockDirt: {
		Name:                 "Dirt",
		ThermalConductivity:  2,
		SpecificHeatCapacity: 1.48,
		Modifier:             1,
	},
	BlockSandstone: {
		Name:                 "Sandstone",
		ThermalConductivity:  2.9,
		SpecificHeatCapacity: 0.8,
		Modifier:             1,
	},
	BlockCopperOre: {
		Name:                 "Copper Ore",
		ThermalConductivity:  4.5,
		SpecificHeatCapacity: 0.386,
		Modifier:             1,
	},
}

var BlockColors = map[BlockID]rl.Color{
	BlockVacuum:    rl.NewColor(0, 0, 0, 255),       // Black color for vacuum
	BlockDirt:      rl.NewColor(139, 69, 19, 255),   // Brown color for dirt
	BlockSandstone: rl.NewColor(210, 180, 140, 255), // Tan color for sandstone
	BlockCopperOre: rl.NewColor(184, 115, 51, 255),  // Copper color for copper ore
}
