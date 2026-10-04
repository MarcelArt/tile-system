package data

import rl "github.com/gen2brain/raylib-go/raylib"

type Tile struct {
	Temperature float32
	Mass        float32
	Block       *Block
	Color       rl.Color
}
