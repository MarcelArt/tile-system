package entities

import (
	"github.com/MarcelArt/tile-system/pkg/constructor"
	"github.com/MarcelArt/tile-system/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type OptionFunc[T any] func(*T)

type Player struct {
	Position rl.Vector2

	texture rl.Texture2D
}

func NewPlayer(opts ...constructor.OptionFunc[Player]) *Player {
	e := &Player{
		Position: rl.Vector2Zero(),
		texture:  engine.GetTexture2D("player"),
	}

	return constructor.New(e, opts...)
}

func WithPosition(pos rl.Vector2) constructor.OptionFunc[Player] {
	return func(e *Player) {
		e.Position = pos
	}
}

func (e *Player) Draw() {
	height := float32(64)
	width := height * float32(e.texture.Width) / float32(e.texture.Height)

	source := rl.NewRectangle(
		0, 0,
		float32(e.texture.Width),
		float32(e.texture.Height),
	)

	destination := rl.NewRectangle(
		e.Position.X,
		e.Position.Y,
		width,
		height,
	)

	rl.DrawTexturePro(
		e.texture,
		source,
		destination,
		rl.Vector2Zero(),
		0,
		rl.White,
	)
}
