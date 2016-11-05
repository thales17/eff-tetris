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
		canvas.AddDrawable(&td)
		canvas.AddKeyDownHandler(func(key string) {
			switch key {
			case "A":
				td.moveLeft()
			case "D":
				td.moveRight()
			case "R":
				td.rotate()
			case "S":
				td.dropTetrimino()
			}
		})
	})
}
