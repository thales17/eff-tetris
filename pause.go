package main

import (
	"math"
	"time"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/component/tween"
)

type pauseScreen struct {
	letterBlocks []letterBlock
	initialized  bool
	tweener      tween.Tweener
}

func (p *pauseScreen) Init(c eff.Canvas) {
	pauseStr := "PAUSE"
	offsetPoint := eff.Point{X: (c.Width() - (len(pauseStr) * letterBlockSize)) / 2, Y: (c.Height() - letterBlockSize) / 2}
	p.letterBlocks = append(p.letterBlocks, letterBlocksForString("PAUSE", offsetPoint)...)
	p.initialized = true
	angle := float64(0)
	p.tweener = tween.NewTweener(time.Millisecond*500, func(progress float64) {
		amp := 25
		for i := range p.letterBlocks {
			x := float64(p.letterBlocks[i].rect.X) + angle
			y := int(math.Sin(x) * float64(amp))
			p.letterBlocks[i].rect.Y = offsetPoint.Y + y
		}
		angle += 0.08
	}, true, false, nil, nil)
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

	c.FillRect(r, eff.Color{R: 0x00, G: 0x00, B: 0x00, A: 0xAA})

	for _, block := range p.letterBlocks {
		block.draw(blockFont, c)
	}
}

func (p *pauseScreen) Update(c eff.Canvas) {
	p.tweener.Tween()
}
