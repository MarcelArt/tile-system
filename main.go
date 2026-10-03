package main

import (
	"github.com/MarcelArt/tile-system/internal/scenes"
	"github.com/MarcelArt/tile-system/pkg/engine"
)

func main() {
	g := engine.NewGame("Tile System", 1280, 720, 60)
	g.SetActiveScene(scenes.NewWorldScene())
	g.Start()
}
