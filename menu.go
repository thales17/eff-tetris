package main

import (
	"time"

	"github.com/forestgiant/easing"
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
	h := c.Height()
	effStr := "EFF"
	effPt := eff.Point{X: (c.Width() - len(effStr)*letterBlockSize) / 2, Y: (c.Height() - letterBlockSize*2) / 2}
	m.effLetters = append(m.effLetters, letterBlocksForString(effStr, eff.Point{X: effPt.X, Y: -1 * letterBlockSize})...)
	tetrisStr := "TETRIS"
	tetrisPt := eff.Point{X: (c.Width() - len(tetrisStr)*letterBlockSize) / 2, Y: (c.Height()-letterBlockSize*2)/2 + letterBlockSize}
	m.tetrisLetters = append(m.tetrisLetters, letterBlocksForString(tetrisStr, eff.Point{X: tetrisPt.X, Y: h + letterBlockSize})...)

	m.tweener = tween.NewTweener(time.Millisecond*500, func(progress float64) {
		startY := -1 * letterBlockSize
		// startY := 100
		endY := effPt.Y
		y := int(float64(endY-startY) * progress)
		for i := range m.effLetters {
			m.effLetters[i].rect.Y = startY + y
		}

		startY = h + letterBlockSize
		// startY = h - 100
		endY = tetrisPt.Y
		y = int(float64(endY-startY) * progress)
		for i := range m.tetrisLetters {
			m.tetrisLetters[i].rect.Y = startY + y
		}
	}, false, false, nil, easing.BounceOut)

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
