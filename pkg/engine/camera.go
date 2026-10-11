package engine

import rl "github.com/gen2brain/raylib-go/raylib"

func GetWorldMousePosition2D(camera rl.Camera2D) rl.Vector2 {
	mousePos := rl.GetMousePosition()
	return rl.GetScreenToWorld2D(mousePos, camera)
}
