package main

import (
	"math/rand"

	"github.com/forestgiant/eff"
)

type tetrimino struct {
	color       eff.Color
	points      [][]eff.Point
	piece       rune
	rotateIndex int
	point       eff.Point
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

func (t *tetrimino) blocksWithPoint() []block {
	blocks := t.blocks()
	for i := 0; i < 4; i++ {
		blocks[i].X += t.point.X
		blocks[i].Y += t.point.Y
	}

	return blocks
}

func (t *tetrimino) originPoints() []eff.Point {
	return t.points[t.rotateIndex]
}

func (t *tetrimino) currentPoints() []eff.Point {
	return t.testPoints(eff.Point{X: 0, Y: 0})
}

func (t *tetrimino) testPoints(tp eff.Point) []eff.Point {
	var cp []eff.Point
	for _, p := range t.originPoints() {
		np := eff.Point{
			X: p.X + t.point.X + tp.X,
			Y: p.Y + t.point.Y + tp.Y,
		}

		cp = append(cp, np)
	}

	return cp
}

func (t *tetrimino) nextRotationPoints() []eff.Point {
	nextIndex := t.rotateIndex + 1
	if nextIndex >= len(t.points) {
		nextIndex = 0
	}
	var nrp []eff.Point
	for _, p := range t.points[nextIndex] {
		np := eff.Point{
			X: p.X + t.point.X,
			Y: p.Y + t.point.Y,
		}

		nrp = append(nrp, np)
	}

	return nrp
}

func (t *tetrimino) width() int {
	points := t.originPoints()
	w := 0
	for i := 0; i < 4; i++ {
		if points[i].X > w {
			w = points[i].X
		}
	}

	return w + 1
}

func (t *tetrimino) draw(c eff.Canvas) {
	tetriminoBlocks := t.blocks()

	p := eff.Point{
		X: t.point.X,
		Y: t.point.Y,
	}

	for i := 0; i < 4; i++ {
		tetriminoBlocks[i].drawWithPoint(p, c)
	}
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
	t := tetriminoForRune(pieces[rand.Intn(len(pieces))])
	t.point.X = (matrixWidth - t.width()) / 2
	t.point.Y = 0
	return t
}
