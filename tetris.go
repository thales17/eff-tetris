package main

import (
	"fmt"
	"math/rand"
	"time"

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
	rand.Seed(time.Now().UnixNano())
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

		if !t.isBlockClear(eff.Point{X: x, Y: y + 1}) {
			return false
		}
	}

	t.tPoint.Y++

	return true
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
		x := t.tetrimino.currentPoints()[i].X + t.tPoint.X
		y := t.tetrimino.currentPoints()[i].Y + t.tPoint.Y

		if !t.isBlockClear(eff.Point{X: x - 1, Y: y + 1}) {
			return
		}
	}

	t.tPoint.X--

	return
}

func (t *tetris) moveRight() {
	for i := 0; i < 4; i++ {
		x := t.tetrimino.currentPoints()[i].X + t.tPoint.X
		y := t.tetrimino.currentPoints()[i].Y + t.tPoint.Y

		if !t.isBlockClear(eff.Point{X: x + 1, Y: y + 1}) {
			return
		}
	}

	t.tPoint.X++

	return
}

func (t *tetris) rotate() {
	nextPoints := t.tetrimino.nextPoints()
	for _, p := range nextPoints {
		if !t.isBlockClear(eff.Point{X: p.X + t.tPoint.X, Y: p.Y + t.tPoint.Y}) {
			return
		}
	}

	t.tetrimino.rotate()
}

func (t *tetris) isBlockClear(p eff.Point) bool {
	if p.X < 0 || p.Y < 0 {
		return false
	}

	if p.X >= matrixWidth || p.Y >= matrixHeight {
		return false
	}

	for _, b := range t.blocks {
		if b.X == p.X && b.Y == p.Y {
			return false
		}
	}

	return true
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

func (t *tetrimino) nextPoints() []eff.Point {
	nextIndex := t.rotateIndex + 1
	if nextIndex >= len(t.points) {
		nextIndex = 0
	}

	return t.points[nextIndex]
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
