package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Title     string
	Width     int32
	Height    int32
	TargetFPS int32

	scene IScene
}

func NewGame(title string, width int32, height int32, targetFPS int32) *Game {
	return &Game{
		Title:     title,
		Width:     width,
		Height:    height,
		TargetFPS: targetFPS,
	}
}

func (g *Game) SetActiveScene(s IScene) {
	g.scene = s
}

func (g *Game) Start() {
	rl.InitWindow(g.Width, g.Height, g.Title)
	defer rl.CloseWindow()

	rl.SetTargetFPS(g.TargetFPS)

	for !rl.WindowShouldClose() {
		// Scene Update
		res := g.scene.Update()
		if res.NextScene != nil && res.NextScene.GetID() != g.scene.GetID() {
			g.SetActiveScene(res.NextScene)
		}

		rl.BeginDrawing()

		// Scene Draw
		g.scene.Draw()

		rl.DrawFPS(0, 0)
		rl.EndDrawing()
	}
}
