package main

import (
	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/sdl"
)

const (
	matrixWidth  int = 10
	matrixHeight int = 20
	squareSize   int = 30
)

func main() {
	canvas := sdl.NewCanvas(
		"Eff-Tetris",
		matrixWidth*squareSize,
		matrixHeight*squareSize,
		eff.Color{R: 0x00, G: 0x00, B: 0x00, A: 0xFF},
		60,
		true,
	)

	canvas.Run(func() {
		td := tetris{}
		m := menu{}
		showingMenu := true
		swapMenuGame := func() {
			if showingMenu {
				canvas.RemoveDrawable(&m)
				canvas.AddDrawable(&td)
			} else {
				canvas.RemoveDrawable(&td)
				canvas.AddDrawable(&m)
			}

			showingMenu = !showingMenu
		}

		canvas.AddDrawable(&m)

		canvas.AddKeyDownEnumHandler(func(key sdl.Keycode) {
			switch key {
			case sdl.KeyA:
				fallthrough
			case sdl.KeyLeft:
				td.moveLeft()
			case sdl.KeyD:
				fallthrough
			case sdl.KeyRight:
				td.moveRight()
			case sdl.KeySpace:
				td.dropTetrimino()
			case sdl.KeyR:
				fallthrough
			case sdl.KeyUp:
				td.rotate()
			case sdl.KeyS:
				fallthrough
			case sdl.KeyDown:
				td.moveTetrimino()
			case sdl.KeyP:
				td.paused = !td.paused
			case sdl.KeyReturn:
				swapMenuGame()
			}
		})
	})
}
