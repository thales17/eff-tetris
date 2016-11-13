package main

import "github.com/forestgiant/eff"

type menu struct {
	initialized bool
}

func (m *menu) Init(c eff.Canvas) {

	m.initialized = true
}

func (m *menu) Initialized() bool {
	return m.initialized
}

func (m *menu) Draw(c eff.Canvas) {
	r := eff.Rect{
		X: 0,
		Y: 0,
		W: 10,
		H: 10,
	}
	c.FillRect(r, eff.RandomColor())
}

func (m *menu) Update(c eff.Canvas) {

}
