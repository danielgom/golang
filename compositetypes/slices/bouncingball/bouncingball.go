package main

import (
	"fmt"
	"github.com/mattn/go-runewidth"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	const (
		cellEmpty = ' '
		cellBall  = '⚾'

		maxFrames = 1200
		speed     = time.Second / 50

		// initial velocities
		ivx, ivy = 1, 1
	)

	var (
		px, py int        // ball position
		vx, vy = ivx, ivy // velocities

		cell rune // current cell (for caching)
	)

	width, height := screen.Size()
	width /= runewidth.RuneWidth(cellBall)
	height-- // there is a 1 pixel border in my terminal

	// create a drawing buffer
	// *2 for extra spaces
	// +1 for newlines
	buf := make([]rune, 0, (width*2+1)*height)

	screen.Clear()

	for i := 0; i < maxFrames; i++ {
		// calculate the next ball position
		px += vx
		py += vy

		// when the ball hits a border reverse its direction
		if px <= 0 || px >= width-1 {
			vx *= -1
		}
		if py <= 0 || py >= height-1 {
			vy *= -1
		}

		buf = buf[:0]

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				cell = cellEmpty

				if px == x && py == y {
					cell = cellBall
				}

				buf = append(buf, cell, ' ')
			}
			buf = append(buf, '\n')
		}

		screen.MoveTopLeft()
		fmt.Print(string(buf))

		time.Sleep(speed)
	}
}
