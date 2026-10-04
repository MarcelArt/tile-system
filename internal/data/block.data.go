package data

import "github.com/MarcelArt/tile-system/pkg/no"

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
)

var Blocks = map[BlockID]Block{
	BlockVacuum: Block{
		Name:                 "Vacuum",
		ThermalConductivity:  0,
		SpecificHeatCapacity: 0,
		Modifier:             0,
	},
	BlockDirt: Block{
		Name:                 "Dirt",
		ThermalConductivity:  2,
		SpecificHeatCapacity: 1.48,
		Modifier:             1,
	},
	BlockSandstone: Block{
		Name:                 "Sandstone",
		ThermalConductivity:  2.9,
		SpecificHeatCapacity: 0.8,
		Modifier:             1,
	},
	BlockCopperOre: Block{
		Name: "Copper Ore",
	},
}
