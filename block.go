package main

import "github.com/forestgiant/eff"

type block struct {
	eff.Point
	color eff.Color
}

func (b *block) draw(c eff.Canvas) {
	borderColor := eff.Color{
		R: 255,
		G: 255,
		B: 255,
		A: 255,
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
