package main

import (
	"fmt"
	"math/rand"

	"github.com/forestgiant/eff"
)

type tetris struct {
	initialized bool
	blocks      []block
	tetrimino   tetrimino
	t           float64
	updateT     int
	gameOver    bool
	tPoint      eff.Point
}

func (t *tetris) Init(c eff.Canvas) {
	t.tetrimino = randomTetrimino()
	t.initialized = true
	t.tPoint.X = (matrixWidth - t.tetrimino.width()) / 2
	t.tPoint.Y = 0
}

func (t *tetris) Draw(c eff.Canvas) {
	tetriminoBlocks := t.tetrimino.blocks()
	for i := 0; i < 4; i++ {
		tetriminoBlocks[i].drawWithPoint(t.tPoint, c)
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
			blocks := t.tetrimino.blocks()
			for i := 0; i < 4; i++ {
				b := blocks[i]
				b.X += t.tPoint.X
				b.Y += t.tPoint.Y
				t.blocks = append(t.blocks, b)
			}
			t.tetrimino = randomTetrimino()
			t.tPoint.X = (matrixWidth - t.tetrimino.width()) / 2
			t.tPoint.Y = 0

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
		x := t.tetrimino.currentPoints()[i].X + t.tPoint.X
		y := t.tetrimino.currentPoints()[i].Y + t.tPoint.Y
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
		x := t.tetrimino.currentPoints()[i].X + t.tPoint.X
		y := t.tetrimino.currentPoints()[i].Y + t.tPoint.Y
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
		x := t.tetrimino.currentPoints()[i].X + t.tPoint.X
		y := t.tetrimino.currentPoints()[i].Y + t.tPoint.Y
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
	t.tetrimino.rotate()
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

type tetrimino struct {
	color       eff.Color
	points      [][]eff.Point
	piece       rune
	rotateIndex int
}

func (t *tetrimino) rotate() {
	t.rotateIndex++
	if t.rotateIndex >= len(t.points) {
		t.rotateIndex = 0
	}
}

func (t *tetrimino) blocks() []block {
	var blocks []block
	for i := 0; i < 4; i++ {
		b := block{}
		b.X = t.points[t.rotateIndex][i].X
		b.Y = t.points[t.rotateIndex][i].Y
		b.color = t.color
		blocks = append(blocks, b)
	}

	return blocks
}

func (t *tetrimino) currentPoints() []eff.Point {
	return t.points[t.rotateIndex]
}

func (t *tetrimino) width() int {
	points := t.currentPoints()
	w := 0
	for i := 0; i < 4; i++ {
		if points[i].X > w {
			w = points[i].X
		}
	}

	return w + 1
}

func tetriminoForRune(piece rune) tetrimino {
	t := tetrimino{piece: piece}
	switch piece {
	case 'i':
		t.color = eff.Color{R: 45, G: 255, B: 254, A: 255}
		t.points = append(t.points, []eff.Point{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}})
		t.points = append(t.points, []eff.Point{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 0, Y: 3}})
	case 'j':
		t.color = eff.Color{R: 11, G: 36, B: 251, A: 255}
		t.points = append(t.points, []eff.Point{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 2, Y: 1}})
		t.points = append(t.points, []eff.Point{{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}, {X: 0, Y: 2}})
		t.points = append(t.points, []eff.Point{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}})
		t.points = append(t.points, []eff.Point{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 1, Y: 0}})
	case 'l':
		t.color = eff.Color{R: 253, G: 164, B: 40, A: 255}
		t.points = append(t.points, []eff.Point{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 0, Y: 1}})
		t.points = append(t.points, []eff.Point{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}})
		t.points = append(t.points, []eff.Point{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 2, Y: 0}})
		t.points = append(t.points, []eff.Point{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 1, Y: 2}})
	case 'o':
		t.color = eff.Color{R: 255, G: 253, B: 56, A: 255}
		t.points = append(t.points, []eff.Point{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}})
	case 's':
		t.color = eff.Color{R: 41, G: 253, B: 47, A: 255}
		t.points = append(t.points, []eff.Point{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 1, Y: 0}, {X: 2, Y: 0}})
		t.points = append(t.points, []eff.Point{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 1, Y: 2}})
	case 't':
		t.color = eff.Color{R: 169, G: 38, B: 251, A: 255}
		t.points = append(t.points, []eff.Point{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 1, Y: 1}})
		t.points = append(t.points, []eff.Point{{X: 1, Y: 0}, {X: 1, Y: 1}, {X: 1, Y: 2}, {X: 0, Y: 1}})
		t.points = append(t.points, []eff.Point{{X: 0, Y: 1}, {X: 1, Y: 1}, {X: 2, Y: 1}, {X: 1, Y: 0}})
		t.points = append(t.points, []eff.Point{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: 2}, {X: 1, Y: 1}})
	case 'z':
		t.color = eff.Color{R: 252, G: 13, B: 27, A: 255}
		t.points = append(t.points, []eff.Point{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 1}})
		t.points = append(t.points, []eff.Point{{X: 0, Y: 2}, {X: 0, Y: 1}, {X: 1, Y: 1}, {X: 1, Y: 0}})

	}

	return t
}

func randomTetrimino() tetrimino {
	pieces := []rune{'i', 'j', 'l', 'o', 's', 't', 'z'}
	return tetriminoForRune(pieces[rand.Intn(len(pieces))])
}
