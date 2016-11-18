package main

import "github.com/forestgiant/eff"

type pauseScreen struct {
	letterBlocks []letterBlock
	initialized  bool
	progress     float64
	progressStep float64
}

func (p *pauseScreen) Init(c eff.Canvas) {
	p.progressStep = float64(1) / float64(25)
	p.letterBlocks = append(p.letterBlocks, letterBlocksForString("PAUSE", eff.Point{X: 0, Y: 0}, c, false)...)
	p.initialized = true
}

func (p *pauseScreen) Initialized() bool {
	return p.initialized
}

func (p *pauseScreen) Draw(c eff.Canvas) {
	r := eff.Rect{
		X: 0,
		Y: 0,
		W: c.Width(),
		H: c.Height(),
	}

	c.FillRect(r, eff.Color{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xCC})

	for _, block := range p.letterBlocks {
		block.draw(c)
	}
}

func (p *pauseScreen) Update(c eff.Canvas) {
	if p.progress < 1 {
		p.progress += p.progressStep
	}

	for i := range p.letterBlocks {
		p.letterBlocks[i].rect = p.letterBlocks[i].mover(p.letterBlocks[i].rect, p.progress)
	}
}
