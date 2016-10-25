package main

import (
	"fmt"

	"github.com/forestgiant/eff"
)

type tetris struct {
	initialized bool
	blocks      []block
	tetrimino   []block
	t           float64
	updateT     int
	gameOver    bool
	tPoint      eff.Point
}

func (t *tetris) Init(c eff.Canvas) {
	t.tetrimino = randomTetrimino()
	t.initialized = true
	t.tPoint.X = 5
	t.tPoint.Y = 0
}

func (t *tetris) Draw(c eff.Canvas) {
	for i := 0; i < 4; i++ {
		t.tetrimino[i].drawWithPoint(t.tPoint, c)
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
			for i := 0; i < 4; i++ {
				b := t.tetrimino[i]
				b.X += t.tPoint.X
				b.Y += t.tPoint.Y
				t.blocks = append(t.blocks, b)
			}
			t.tPoint.X = 5
			t.tPoint.Y = 0
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
		x := b.X + t.tPoint.X
		y := b.Y + t.tPoint.Y
		if y == matrixHeight-1 {
			return false
		}
		for j := 0; j < len(t.blocks); j++ {
			if x == t.blocks[j].X && y+1 == t.blocks[j].Y {
				return false
			}
		}
	}

	t.tPoint.Y++

	return true
}

// func (t *tetris) gravity() bool {
// 	fallingBlocks := false
// 	for i := 0; i < len(t.blocks); i++ {
// 		//Check if on the bottom
// 		if t.blocks[i].Y == matrixHeight-1 {
// 			continue
// 		}

// 		// Check for a block below
// 		belowBlock := false
// 		for j := 0; j < len(t.blocks); j++ {
// 			if t.blocks[j].X == t.blocks[i].X && t.blocks[j].Y == t.blocks[i].Y+1 {
// 				belowBlock = true
// 				break
// 			}
// 		}

// 		// No block below and not on the bottom moving the block down
// 		if !belowBlock {
// 			t.blocks[i].Y++
// 			fallingBlocks = true
// 		}
// 	}

// 	return fallingBlocks
// }

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
		x := b.X + t.tPoint.X
		y := b.Y + t.tPoint.Y
		if x == 0 {
			return
		}
		for j := 0; j < len(t.blocks); j++ {
			if y == t.blocks[j].Y && x-1 == t.blocks[j].X {
				return
			}
		}
	}

	t.tPoint.X--

	return
}

func (t *tetris) moveRight() {
	for i := 0; i < 4; i++ {
		b := t.tetrimino[i]
		x := b.X + t.tPoint.X
		y := b.Y + t.tPoint.Y
		if x == matrixWidth-1 {
			return
		}
		for j := 0; j < len(t.blocks); j++ {
			if y == t.blocks[j].Y && x+1 == t.blocks[j].X {
				return
			}
		}
	}

	t.tPoint.X++

	return
}

func (t *tetris) rotate() {

	for i := 0; i < 4; i++ {
		x := t.tetrimino[i].X
		t.tetrimino[i].X = t.tetrimino[i].Y * -1
		t.tetrimino[i].Y = x
	}
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

func (b *block) drawWithPoint(p eff.Point, c eff.Canvas) {
	bPrime := block{}
	bPrime.X = b.X + p.X
	bPrime.Y = b.Y + p.Y
	bPrime.color = b.color
	bPrime.draw(c)
}

func tetriminoForRune(piece rune) []block {
	var t []block
	switch piece {
	case 'i':
		color := eff.Color{R: 45, G: 255, B: 254, A: 255}
		for i := -2; i < 2; i++ {
			b := block{}
			b.X = i
			b.Y = 0
			b.color = color
			t = append(t, b)
		}
	case 'j':
		color := eff.Color{R: 11, G: 36, B: 251, A: 255}
		for i := -1; i < 2; i++ {
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
	case 'l':
		color := eff.Color{R: 253, G: 164, B: 40, A: 255}
		for i := -1; i < 2; i++ {
			b := block{}
			b.X = i
			b.Y = 0
			b.color = color
			t = append(t, b)
		}
		b := block{}
		b.X = -1
		b.Y = 1
		b.color = color
		t = append(t, b)
	case 'o':
		color := eff.Color{R: 255, G: 253, B: 56, A: 255}
		for i := 0; i < 4; i++ {
			b := block{}
			b.X = i%2 - 1
			b.Y = (i / 2)
			b.color = color
			t = append(t, b)
		}
	case 's':
		color := eff.Color{R: 41, G: 253, B: 47, A: 255}
		for i := -2; i < 2; i++ {
			b := block{}
			b.X = i
			if i < 0 {
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
		for i := -1; i < 2; i++ {
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
	case 'z':
		color := eff.Color{R: 252, G: 13, B: 27, A: 255}
		for i := -2; i < 2; i++ {
			b := block{}
			b.X = i
			if i < 0 {
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
	// pieces := []rune{'i', 'j', 'l', 'o', 's', 't', 'z'}
	// return tetriminoForRune(pieces[rand.Intn(len(pieces))])
	return tetriminoForRune('z')
}
