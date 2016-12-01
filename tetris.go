package main

import (
	"fmt"

	"github.com/forestgiant/eff"
)

const (
	maxTime = 100
)

type tetris struct {
	initialized      bool
	blocks           []block
	tetrimino        tetrimino
	nextTetrimino    tetrimino
	time             int
	speed            int
	gameOver         bool
	paused           bool
	ps               pauseScreen
	linesCleared     int
	gameOverCallback func()
}

func (t *tetris) Init(c eff.Canvas) {

	t.tetrimino = randomTetrimino()
	t.nextTetrimino = randomTetrimino()
	t.initialized = true
	t.speed = 3
}

func (t *tetris) Draw(c eff.Canvas) {
	t.tetrimino.draw(c)

	for i := 0; i < len(t.blocks); i++ {
		t.blocks[i].draw(c)
	}

	t.drawScoreboard(c)
}

func (t *tetris) Update(c eff.Canvas) {
	if t.gameOver || t.paused {
		return
	}

	t.time += t.speed
	if t.time > maxTime {
		t.time = 0
		if !t.moveTetrimino() {
			t.blocks = append(t.blocks, t.tetrimino.blocksWithPoint()...)
			t.tetrimino = t.nextTetrimino
			t.nextTetrimino = randomTetrimino()

			t.clearLines()
			t.gameOver = t.isGameOver()
			if t.gameOver {
				fmt.Println("Game over man!")
				if t.gameOverCallback != nil {
					t.gameOverCallback()
				}
			}
		}
	}

}

func (t *tetris) Initialized() bool {
	return t.initialized
}

func (t *tetris) moveTetrimino() bool {
	if t.paused {
		return false
	}
	if t.arePointsClear(t.tetrimino.testPoints(eff.Point{X: 0, Y: 1})) {
		t.tetrimino.point.Y++
		return true
	}

	return false
}

func (t *tetris) dropTetrimino() {
	if t.paused {
		return
	}
	for t.moveTetrimino() {
	}
}

func (t *tetris) isGameOver() bool {
	for i := 0; i < len(t.blocks); i++ {
		if t.blocks[i].Y == 0 {
			return true
		}
	}

	return false
}

func (t *tetris) clearLines() bool {
	linesCleared := 0
	for i := 0; i < matrixHeight; i++ {
		lineBlocks := 0
		for _, b := range t.blocks {
			if b.Y == i {
				lineBlocks++
			}
		}

		if lineBlocks == matrixWidth {
			fmt.Println("Line!", i)
			linesCleared++

			var newBlocks []block
			for _, b := range t.blocks {
				if b.Y != i {
					if b.Y < i {
						b.Y++
					}
					newBlocks = append(newBlocks, b)
				}
			}

			t.blocks = newBlocks
		}
	}
	t.linesCleared += linesCleared
	return linesCleared > 0
}

func (t *tetris) moveLeft() bool {
	if t.paused {
		return false
	}
	if t.arePointsClear(t.tetrimino.testPoints(eff.Point{X: -1, Y: 0})) {
		t.tetrimino.point.X--
		return true
	}

	return false
}

func (t *tetris) moveRight() bool {
	if t.paused {
		return false
	}
	if t.arePointsClear(t.tetrimino.testPoints(eff.Point{X: 1, Y: 0})) {
		t.tetrimino.point.X++
		return true
	}

	return false
}

func (t *tetris) rotate() bool {
	if t.paused {
		return false
	}
	if t.arePointsClear(t.tetrimino.nextRotationPoints()) {
		t.tetrimino.rotate()
		return true
	}

	return false
}

func (t *tetris) isBlockClear(p eff.Point) bool {
	if p.X < 0 || p.Y < 0 {
		return false
	}

	if p.X >= matrixWidth || p.Y >= matrixHeight {
		return false
	}

	for _, b := range t.blocks {
		if b.X == p.X && b.Y == p.Y {
			return false
		}
	}

	return true
}

func (t *tetris) arePointsClear(points []eff.Point) bool {
	for _, p := range points {
		if !t.isBlockClear(p) {
			return false
		}
	}

	return true
}

func (t *tetris) togglePause(c eff.Canvas) {
	t.paused = !t.paused
	if t.paused {
		t.ps = pauseScreen{}
		c.AddDrawable(&t.ps)
	} else {
		c.RemoveDrawable(&t.ps)
	}
}

func (t *tetris) drawScoreboard(c eff.Canvas) {
	scoreStr := fmt.Sprintf("Lines Cleared: %d", t.linesCleared)
	c.DrawText(scoreboardFont, scoreStr, eff.White(), eff.Point{X: 10, Y: 10})
	tetrimino := tetriminoForRune(t.nextTetrimino.piece)
	tetrimino.drawPreview(c)
}
