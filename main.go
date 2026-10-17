package main

import (
	"github.com/MarcelArt/tile-system/internal/scenes"
	"github.com/MarcelArt/tile-system/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	g := engine.NewGame("Tile System", 1280, 720, 60)
	g.OnStart = func() {
		g.SetActiveScene(scenes.NewWorldScene())
		blockTexture := engine.SetTexture2D("blocks", "assets/sprites/blocks.png")
		rl.SetTextureFilter(blockTexture, rl.FilterPoint)
	}
	g.OnExit = func() {
		engine.UnloadAllTexture2D()
	}

	g.Start()

}
