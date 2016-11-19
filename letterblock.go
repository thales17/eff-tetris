package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/forestgiant/eff"
	"github.com/forestgiant/eff/util"
)

type letterBlock struct {
	letter rune
	color  eff.Color
	rect   eff.Rect
}

func (l *letterBlock) draw(c eff.Canvas) {
	c.FillRect(l.rect, l.color)

	t := string(l.letter)
	lp, err := util.CenterTextInRect(t, l.rect, c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	textColor := eff.Black()
	if rand.Intn(10) >= 5 {
		textColor = eff.White()
	}

	err = c.DrawText(t, textColor, lp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func letterBlocksForString(s string, offset eff.Point) []letterBlock {
	var colors []eff.Color
	colors = append(colors, eff.Color{R: 45, G: 255, B: 254, A: 255})
	colors = append(colors, eff.Color{R: 11, G: 36, B: 251, A: 255})
	colors = append(colors, eff.Color{R: 253, G: 164, B: 40, A: 255})
	colors = append(colors, eff.Color{R: 255, G: 253, B: 56, A: 255})
	colors = append(colors, eff.Color{R: 41, G: 253, B: 47, A: 255})
	colors = append(colors, eff.Color{R: 169, G: 38, B: 251, A: 255})
	colors = append(colors, eff.Color{R: 252, G: 13, B: 27, A: 255})

	blockSize := 30

	var letterBlocks []letterBlock
	for i, letter := range s {
		letterBlocks = append(letterBlocks, letterBlock{
			letter: letter,
			color:  colors[rand.Intn(len(colors))],
			rect: eff.Rect{
				X: offset.X + i*blockSize,
				Y: offset.Y,
				W: blockSize,
				H: blockSize,
			},
		})
	}

	return letterBlocks
}
