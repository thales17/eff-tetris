package main

import (
	"time"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/component/tween"
)

type menu struct {
	effLetters    []letterBlock
	tetrisLetters []letterBlock
	initialized   bool
	progress      float64
	progressStep  float64
	tweener       tween.Tweener
}

func (m *menu) Init(c eff.Canvas) {
	m.progressStep = float64(1) / float64(25)
	blockSize := 30
	h := c.Height()
	m.effLetters = append(m.effLetters, letterBlocksForString("EFF", eff.Point{X: 0, Y: -200})...)
	m.tetrisLetters = append(m.tetrisLetters, letterBlocksForString("TETRIS", eff.Point{X: 0, Y: h + blockSize})...)

	m.tweener = tween.NewTweener(time.Millisecond*500, func(progress float64) {
		startY := -200
		endY := 200
		y := int(float64(endY-startY) * progress)
		for i := range m.effLetters {
			m.effLetters[i].rect.Y = startY + y
		}

		startY = h + blockSize
		endY = 200 + blockSize
		y = int(float64(endY-startY) * progress)
		for i := range m.tetrisLetters {
			m.tetrisLetters[i].rect.Y = startY + y
		}
	}, false)

	m.initialized = true
}

func (m *menu) Initialized() bool {
	return m.initialized
}

func (m *menu) Draw(c eff.Canvas) {
	for _, block := range m.effLetters {
		block.draw(c)
	}

	for _, block := range m.tetrisLetters {
		block.draw(c)
	}
}

func (m *menu) Update(c eff.Canvas) {
	if m.progress < 1 {
		m.progress += m.progressStep
	}

	if !m.tweener.Done {
		m.tweener.Tween()
	}

}
