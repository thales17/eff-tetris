package main

import (
	"fmt"
	"math/rand"

	"github.com/forestgiant/eff"
)

type tetris struct {
	initialized bool
	blocks      []block
	tetrimino   []block
	t           float64
	updateT     int
	gameOver    bool
}

func (t *tetris) Init(c eff.Canvas) {
	t.tetrimino = randomTetrimino()
	t.initialized = true
}

func (t *tetris) Draw(c eff.Canvas) {
	for i := 0; i < 4; i++ {
		t.tetrimino[i].draw(c)
	}
	for i := 0; i < len(t.blocks); i++ {
		t.blocks[i].draw(c)
	}
}

func (t *tetris) Update(c eff.Canvas) {
	if t.gameOver {
		return
	}
	t.t += 0.01
	if t.t > 1 {
		t.t = 0
	}

	if int(float64(10.0)*t.t) != t.updateT {
		t.updateT = int(float64(10.0) * t.t)
		if !t.moveTetrimino() {
			t.blocks = append(t.blocks, t.tetrimino...)
			t.tetrimino = randomTetrimino()
		}
	}

	t.gameOver = t.isGameOver()
	if t.gameOver {
		fmt.Println("Game over man!")
	}
}

func (t *tetris) Initialized() bool {
	return t.initialized
}

func (t *tetris) moveTetrimino() bool {
	for i := 0; i < 4; i++ {
		b := t.tetrimino[i]
		if b.Y == matrixHeight-1 {
			return false
		}
		for j := 0; j < len(t.blocks); j++ {
			if b.X == t.blocks[j].X && b.Y+1 == t.blocks[j].Y {
				return false
			}
		}
	}

	for i := 0; i < 4; i++ {
		t.tetrimino[i].Y++
	}

	return true
}

func (t *tetris) gravity() bool {
	fallingBlocks := false
	for i := 0; i < len(t.blocks); i++ {
		//Check if on the bottom
		if t.blocks[i].Y == matrixHeight-1 {
			continue
		}

		// Check for a block below
		belowBlock := false
		for j := 0; j < len(t.blocks); j++ {
			if t.blocks[j].X == t.blocks[i].X && t.blocks[j].Y == t.blocks[i].Y+1 {
				belowBlock = true
				break
			}
		}

		// No block below and not on the bottom moving the block down
		if !belowBlock {
			t.blocks[i].Y++
			fallingBlocks = true
		}
	}

	return fallingBlocks
}

func (t *tetris) isGameOver() bool {
	for i := 0; i < len(t.blocks); i++ {
		if t.blocks[i].Y == 0 {
			return true
		}
	}

	return false
}

func (t *tetris) moveLeft() {
	for i := 0; i < 4; i++ {
		b := t.tetrimino[i]
		if b.X == 0 {
			return
		}
		for j := 0; j < len(t.blocks); j++ {
			if b.Y == t.blocks[j].Y && b.X-1 == t.blocks[j].X {
				return
			}
		}
	}

	for i := 0; i < 4; i++ {
		t.tetrimino[i].X--
	}

	return
}

func (t *tetris) moveRight() {
	for i := 0; i < 4; i++ {
		b := t.tetrimino[i]
		if b.X == matrixWidth-1 {
			return
		}
		for j := 0; j < len(t.blocks); j++ {
			if b.Y == t.blocks[j].Y && b.X+1 == t.blocks[j].X {
				return
			}
		}
	}

	for i := 0; i < 4; i++ {
		t.tetrimino[i].X++
	}

	return
}

type block struct {
	eff.Point
	color eff.Color
}

func (b *block) draw(c eff.Canvas) {
	borderColor := eff.Color{
		R: 255,
		G: 255,
		B: 255,
	}
	x := b.X * squareSize
	y := b.Y * squareSize
	spacing := squareSize / 10

	borderRect := eff.Rect{
		X: x,
		Y: y,
		W: squareSize,
		H: squareSize,
	}

	fillRect := eff.Rect{
		X: x + spacing,
		Y: y + spacing,
		W: squareSize - (spacing * 2),
		H: squareSize - (spacing * 2),
	}

	c.DrawRect(borderRect, borderColor)
	c.FillRect(fillRect, b.color)
}

func nextBlock() block {
	block := block{}
	block.X = rand.Intn(matrixWidth)
	block.Y = -1
	block.color = eff.RandomColor()

	return block
}

func tetriminoForRune(piece rune) []block {
	var t []block

	switch piece {
	case 'i':
		color := eff.Color{R: 45, G: 255, B: 254, A: 255}
		for i := 0; i < 4; i++ {
			b := block{}
			b.X = i
			b.Y = 0
			b.color = color
			t = append(t, b)
		}
	case 'j':
		color := eff.Color{R: 11, G: 36, B: 251, A: 255}
		for i := 0; i < 3; i++ {
			b := block{}
			b.X = i
			b.Y = 0
			b.color = color
			t = append(t, b)
		}
		b := block{}
		b.X = 2
		b.Y = 1
		b.color = color
		t = append(t, b)
	case 'l':
		color := eff.Color{R: 253, G: 164, B: 40, A: 255}
		for i := 0; i < 3; i++ {
			b := block{}
			b.X = i
			b.Y = 0
			b.color = color
			t = append(t, b)
		}
		b := block{}
		b.X = 0
		b.Y = 1
		b.color = color
		t = append(t, b)
	case 'o':
		color := eff.Color{R: 255, G: 253, B: 56, A: 255}
		for i := 0; i < 4; i++ {
			b := block{}
			b.X = i % 2
			b.Y = (i / 2)
			b.color = color
			t = append(t, b)
		}
	case 's':
		color := eff.Color{R: 41, G: 253, B: 47, A: 255}
		for i := 0; i < 4; i++ {
			b := block{}
			b.X = i
			if i < 2 {
				b.Y = 1
			} else {
				b.Y = 0
				b.X--
			}
			b.color = color
			t = append(t, b)
		}
	case 't':
		color := eff.Color{R: 169, G: 38, B: 251, A: 255}
		for i := 0; i < 3; i++ {
			b := block{}
			b.X = i
			b.Y = 0
			b.color = color
			t = append(t, b)
		}
		b := block{}
		b.X = 1
		b.Y = 1
		b.color = color
		t = append(t, b)
	case 'z':
		color := eff.Color{R: 252, G: 13, B: 27, A: 255}
		for i := 0; i < 4; i++ {
			b := block{}
			b.X = i
			if i < 2 {
				b.Y = 0
			} else {
				b.Y = 1
				b.X--
			}
			b.color = color
			t = append(t, b)
		}
	}

	return t
}

func randomTetrimino() []block {
	pieces := []rune{'i', 'j', 'l', 'o', 's', 't', 'z'}
	return tetriminoForRune(pieces[rand.Intn(len(pieces))])
	// return tetriminoForRune('j')
}
