package main

import (
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/go-downfall/world"
)

type character struct {
	x float64
	y float64
}

var mychar *character
var myWorld world.World = world.World{}

var square *ebiten.Image

func update(screen *ebiten.Image) error {
	screen.Fill(color.RGBA{0xff, 0x00, 0x00, 0xff})
	ebitenutil.DebugPrint(screen, "Game")

	if square == nil {
		square, _ = ebiten.NewImage(16, 16, ebiten.FilterNearest)
		mychar = &character{64, 64}
	}

	square.Fill(color.White)
	opts := &ebiten.DrawImageOptions{}

	opts.GeoM.Translate(mychar.x, mychar.y)
	screen.DrawImage(square, opts)

	handleInput(screen)

	return nil
}

func handleInput(screen *ebiten.Image) {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		mychar.y -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		mychar.y += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		mychar.x -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		mychar.x += 1
	}
}

type A struct {
	Val string
}

func main() {
	ebiten.Run(update, 320, 240, 2, "DownFall")
}

