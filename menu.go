package main

import (
	"fmt"
	"os"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/util"
)

type letterBlock struct {
	letter      rune
	letterColor eff.Color
	color       eff.Color
	rect        eff.Rect
}

func (l *letterBlock) draw(c eff.Canvas) {
	c.FillRect(l.rect, l.color)

	t := string(l.letter)
	lp, err := util.CenterTextInRect(t, l.rect, c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = c.DrawText(t, l.letterColor, lp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type menu struct {
	letterBlocks []letterBlock
	initialized  bool
}

func (m *menu) Init(c eff.Canvas) {
	font := eff.Font{
		Path: "assets/fonts/roboto/Roboto-Bold.ttf",
	}

	err := c.SetFont(font, 24)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	blockSize := 30

	m.letterBlocks = append(m.letterBlocks, letterBlock{
		letter:      'E',
		letterColor: eff.RandomColor(),
		color:       eff.RandomColor(),
		rect: eff.Rect{
			X: 0,
			Y: 0,
			W: blockSize,
			H: blockSize,
		},
	})

	m.letterBlocks = append(m.letterBlocks, letterBlock{
		letter:      'F',
		letterColor: eff.RandomColor(),
		color:       eff.RandomColor(),
		rect: eff.Rect{
			X: blockSize,
			Y: 0,
			W: blockSize,
			H: blockSize,
		},
	})

	m.letterBlocks = append(m.letterBlocks, letterBlock{
		letter:      'F',
		letterColor: eff.RandomColor(),
		color:       eff.RandomColor(),
		rect: eff.Rect{
			X: blockSize * 2,
			Y: 0,
			W: blockSize,
			H: blockSize,
		},
	})

	m.letterBlocks = append(m.letterBlocks, letterBlock{
		letter:      'T',
		letterColor: eff.RandomColor(),
		color:       eff.RandomColor(),
		rect: eff.Rect{
			X: 0,
			Y: blockSize,
			W: blockSize,
			H: blockSize,
		},
	})

	m.letterBlocks = append(m.letterBlocks, letterBlock{
		letter:      'E',
		letterColor: eff.RandomColor(),
		color:       eff.RandomColor(),
		rect: eff.Rect{
			X: blockSize,
			Y: blockSize,
			W: blockSize,
			H: blockSize,
		},
	})

	m.letterBlocks = append(m.letterBlocks, letterBlock{
		letter:      'T',
		letterColor: eff.RandomColor(),
		color:       eff.RandomColor(),
		rect: eff.Rect{
			X: blockSize * 2,
			Y: blockSize,
			W: blockSize,
			H: blockSize,
		},
	})

	m.letterBlocks = append(m.letterBlocks, letterBlock{
		letter:      'R',
		letterColor: eff.RandomColor(),
		color:       eff.RandomColor(),
		rect: eff.Rect{
			X: blockSize * 3,
			Y: blockSize,
			W: blockSize,
			H: blockSize,
		},
	})

	m.letterBlocks = append(m.letterBlocks, letterBlock{
		letter:      'I',
		letterColor: eff.RandomColor(),
		color:       eff.RandomColor(),
		rect: eff.Rect{
			X: blockSize * 4,
			Y: blockSize,
			W: blockSize,
			H: blockSize,
		},
	})

	m.letterBlocks = append(m.letterBlocks, letterBlock{
		letter:      'S',
		letterColor: eff.RandomColor(),
		color:       eff.RandomColor(),
		rect: eff.Rect{
			X: blockSize * 5,
			Y: blockSize,
			W: blockSize,
			H: blockSize,
		},
	})
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

}
