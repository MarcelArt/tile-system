package entities

import (
	"log"
	"math"

	"github.com/MarcelArt/tile-system/internal/data"
	"github.com/MarcelArt/tile-system/pkg/constructor"
	"github.com/MarcelArt/tile-system/pkg/engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Position      rl.Vector2
	Width, Height float32

	velocity   rl.Vector2
	isGrounded bool
	texture    rl.Texture2D
}

func NewPlayer(opts ...constructor.OptionFunc[Player]) *Player {
	e := &Player{
		Position:   rl.Vector2Zero(),
		texture:    engine.GetTexture2D("player"),
		isGrounded: true,
	}

	return constructor.New(e, opts...)
}

func WithPlayerPosition(pos rl.Vector2) constructor.OptionFunc[Player] {
	return func(e *Player) {
		e.Position = pos
	}
}

func WithPlayerWidthAndHeight(width, height float32) constructor.OptionFunc[Player] {
	return func(p *Player) {
		p.Width = width
		p.Height = height
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

func (e *Player) Update(dt float32, grid *TileGrid) {
	e.fall(dt, grid)
}

func (e *Player) Bounds() rl.Rectangle {
	return rl.NewRectangle(e.Position.X, e.Position.Y, e.Width, e.Height)
}

func (e *Player) fall(dt float32, grid *TileGrid) {
	log.Println("e.Position :>> ", e.Position)
	e.velocity.Y = min(e.velocity.Y+(data.Gravity*dt), data.MaxFall)

	e.Position.X += e.velocity.X * dt
	if grid.OverlapsSolid(e.Bounds()) {
		if e.velocity.X > 0 {
			e.Position.X = float32(math.Floor(float64(e.Bounds().X+e.Width)/32))*32 - e.Width
		} else {
			e.Position.X = (float32(math.Floor(float64(e.Bounds().X)/32)) + 1) * 32
		}
		e.velocity.X = 0
	}

	// --- Y axis ---
	e.Position.Y += e.velocity.Y * dt
	e.isGrounded = false
	if grid.OverlapsSolid(e.Bounds()) {
		if e.velocity.Y > 0 { // falling: snap feet to tile top
			e.Position.Y = float32(math.Floor(float64(e.Bounds().Y+e.Height)/32))*32 - e.Height
			e.isGrounded = true
		} else { // jumping: snap head to tile bottom
			e.Position.Y = (float32(math.Floor(float64(e.Bounds().Y)/32)) + 1) * 32
		}
		e.velocity.Y = 0
	}
}
