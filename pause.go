package main

import "github.com/forestgiant/eff"

type pauseScreen struct {
	initialized bool
}

func (p *pauseScreen) Init(c eff.Canvas) {

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

	c.FillRect(r, eff.Color{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xEE})
}

func (p *pauseScreen) Update(c eff.Canvas) {

}
