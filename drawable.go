package main

import "github.com/forestgiant/eff"

type tetrisDrawable struct {
	initialized bool
	blocks      []block
	t           float64
}

func (t *tetrisDrawable) Init(c eff.Canvas) {
	t.blocks = testBlocks()
	t.initialized = true
}

func (t *tetrisDrawable) Draw(c eff.Canvas) {
	index := int(float64(matrixWidth*matrixHeight) * t.t)
	for i := 0; i <= index; i++ {
		t.blocks[i].draw(c)
	}
}

func (t *tetrisDrawable) Update(c eff.Canvas) {
	t.t += 0.001
	if t.t > 1 {
		t.t = 0
	}
}

func (t *tetrisDrawable) Initialized() bool {
	return t.initialized
}

type block struct {
	point eff.Point
	color eff.Color
}

func (b *block) draw(c eff.Canvas) {
	borderColor := eff.Color{
		R: 255,
		G: 255,
		B: 255,
	}
	x := b.point.X * squareSize
	y := b.point.Y * squareSize
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

func testBlocks() []block {
	var blocks []block
	for i := 0; i < matrixWidth*matrixHeight; i++ {
		b := block{}
		b.point.X = i % matrixWidth
		b.point.Y = int(float64(i) / float64(matrixWidth))
		b.color = eff.RandomColor()
		blocks = append(blocks, b)
	}

	return blocks
}
