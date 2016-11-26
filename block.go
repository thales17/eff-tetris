package main

import "github.com/forestgiant/eff"

type block struct {
	eff.Point
	color       eff.Color
	size        int
	offsetPoint eff.Point
}

func (b *block) draw(c eff.Canvas) {
	s := b.size
	if s == 0 {
		s = squareSize
	}

	borderColor := eff.Color{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
	}
	x := b.X * s
	y := b.Y*s + scoreboardHeight
	spacing := s / 10

	borderRect := eff.Rect{
		X: x,
		Y: y,
		W: s,
		H: s,
	}

	fillRect := eff.Rect{
		X: x + spacing,
		Y: y + spacing,
		W: s - (spacing * 2),
		H: s - (spacing * 2),
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
