package main

import "github.com/forestgiant/eff/sdl"

const (
	matrixWidth  int = 10
	matrixHeight int = 20
	squareSize   int = 25
)

func main() {
	canvas := sdl.NewCanvas(
		"Eff-Tetris",
		matrixWidth*squareSize,
		matrixHeight*squareSize,
		60,
		true,
	)

	canvas.Run(func() {
		td := tetrisDrawable{}
		canvas.AddDrawable(&td)
	})
}
