package main

import (
	"fmt"
	"math/rand"

	"github.com/forestgiant/eff"
)

type tetris struct {
	initialized bool
	blocks      []block
	t           float64
	updateT     int
	gameOver    bool
}

func (t *tetris) Init(c eff.Canvas) {
	t.blocks = append(t.blocks, nextBlock())
	t.initialized = true
}

func (t *tetris) Draw(c eff.Canvas) {
	for i := 0; i < len(t.blocks); i++ {
		t.blocks[i].draw(c)
	}
}

func (t *tetris) Update(c eff.Canvas) {
	if t.gameOver {
		return
	}
	t.t += 0.8
	if t.t > 1 {
		t.t = 0
	}
	if int(10.0*t.t) != t.updateT {
		t.updateT = int(10.0 * t.t)
		if !t.gravity() {
			t.blocks = append(t.blocks, nextBlock())
		}
	}

	t.gameOver = t.isGameOver()
	if t.gameOver {
		fmt.Println("Game over man!")
	}
}

func (t *tetris) Initialized() bool {
	return t.initialized
}

func (t *tetris) gravity() bool {
	fallingBlocks := false
	for i := 0; i < len(t.blocks); i++ {
		//Check if on the bottom
		if t.blocks[i].Y == matrixHeight-1 {
			continue
		}

		// Check for a block below
		belowBlock := false
		for j := 0; j < len(t.blocks); j++ {
			if t.blocks[j].X == t.blocks[i].X && t.blocks[j].Y == t.blocks[i].Y+1 {
				belowBlock = true
				break
			}
		}

		// No block below and not on the bottom moving the block down
		if !belowBlock {
			t.blocks[i].Y++
			fallingBlocks = true
		}
	}

	return fallingBlocks
}

func (t *tetris) isGameOver() bool {
	for i := 0; i < len(t.blocks); i++ {
		if t.blocks[i].Y < 0 {
			belowBlock := false
			for j := 0; j < len(t.blocks); j++ {
				if t.blocks[j].X == t.blocks[i].X && t.blocks[j].Y == t.blocks[i].Y+1 {
					belowBlock = true
					break
				}
			}
			if belowBlock {
				return true
			}
		}
	}

	return false
}

type block struct {
	eff.Point
	color eff.Color
}

func (b *block) draw(c eff.Canvas) {
	borderColor := eff.Color{
		R: 255,
		G: 255,
		B: 255,
	}
	x := b.X * squareSize
	y := b.Y * squareSize
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

func nextBlock() block {
	block := block{}
	block.X = rand.Intn(matrixWidth)
	block.Y = -1
	block.color = eff.RandomColor()

	return block
}
