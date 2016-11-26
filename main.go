package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/sdl"
)

const (
	matrixWidth      int = 10
	matrixHeight     int = 20
	squareSize       int = 30
	letterBlockSize  int = 30
	scoreboardHeight int = 50
)

func main() {
	canvas := sdl.NewCanvas(
		"Eff-Tetris",
		matrixWidth*squareSize,
		matrixHeight*squareSize+scoreboardHeight,
		eff.Color{R: 0x00, G: 0x00, B: 0x00, A: 0xFF},
		60,
		true,
	)

	canvas.Run(func() {
		rand.Seed(time.Now().UnixNano())
		font := eff.Font{
			Path: "assets/fonts/roboto/Roboto-Bold.ttf",
		}

		err := canvas.SetFont(font, 24)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		td := tetris{}

		m := menu{}

		showingMenu := true

		startGame := func() {
			if !showingMenu {
				return
			}
			showingMenu = false
			canvas.RemoveDrawable(&m)
			td = tetris{}
			td.gameOverCallback = func() {
				if showingMenu {
					return
				}
				canvas.RemoveDrawable(&td)
				m = menu{}
				canvas.AddDrawable(&m)
				showingMenu = true
			}
			canvas.AddDrawable(&td)
		}

		canvas.AddDrawable(&m)

		canvas.AddKeyDownEnumHandler(func(key sdl.Keycode) {
			switch key {
			case sdl.KeyA:
				fallthrough
			case sdl.KeyLeft:
				if !showingMenu {
					td.moveLeft()
				}
			case sdl.KeyD:
				fallthrough
			case sdl.KeyRight:
				if !showingMenu {
					td.moveRight()
				}
			case sdl.KeySpace:
				if !showingMenu {
					td.dropTetrimino()
				}
			case sdl.KeyR:
				fallthrough
			case sdl.KeyUp:
				if !showingMenu {
					td.rotate()
				}
			case sdl.KeyS:
				fallthrough
			case sdl.KeyDown:
				if !showingMenu {
					td.moveTetrimino()
				}
			case sdl.KeyP:
				if !showingMenu {
					td.togglePause(canvas)
				}
			case sdl.KeyReturn:
				startGame()
			}
		})
	})
}
