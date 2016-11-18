package main

import "github.com/forestgiant/eff"

type menu struct {
	letterBlocks []letterBlock
	initialized  bool
	progress     float64
	progressStep float64
}

func (m *menu) Init(c eff.Canvas) {
	m.progressStep = float64(1) / float64(25)
	blockSize := 30

	m.letterBlocks = append(m.letterBlocks, letterBlocksForString("EFF", eff.Point{X: 0, Y: 200}, c, false)...)
	m.letterBlocks = append(m.letterBlocks, letterBlocksForString("TETRIS", eff.Point{X: 0, Y: 200 + blockSize}, c, true)...)
	m.initialized = true
}

func (m *menu) Initialized() bool {
	return m.initialized
}

func (m *menu) Draw(c eff.Canvas) {
	for _, block := range m.letterBlocks {
		block.draw(c)
	}
}

func (m *menu) Update(c eff.Canvas) {
	if m.progress < 1 {
		m.progress += m.progressStep
	}

	for i := range m.letterBlocks {
		m.letterBlocks[i].rect = m.letterBlocks[i].mover(m.letterBlocks[i].rect, m.progress)
	}
}
