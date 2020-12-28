package main

import (
	"time"

	"github.com/inancgumus/screen"
)

func main() {

	const (
		width  = 50
		height = 10

		cellEmpty = ' '
		cellBall  = '⚾'

		maxFrames = 1000
		speed     = time.Second / 50

		// drawing buffer length
		// *2 for extra spaces
		// +1 for newlines
		bufferLen = (width*2 + 1) * height
	)

	var (
		cell rune // current cell (for caching)

		positionX, positionY int    // ball position
		velocityX, velocityY = 1, 1 // velocities
	)

	// create the board
	board := make([][]bool, width)
	for column := range board {
		board[column] = make([]bool, height)
	}

	// a drawing buffer
	buffer := make([]rune, 0, bufferLen)

	// clear the screen
	screen.Clear()

	for i := 0; i < maxFrames; i++ {
		// calculate the next ball position
		positionX += velocityX
		positionY += velocityY

		// reverse its direction if the ball hits a border
		if positionX == 0 || positionX == width-1 {
			velocityX *= -1
		}
		if positionY == 0 || positionY == height-1 {
			velocityY *= -1
		}

		// put the new ball
		board[positionX][positionY] = true

		// rewind the buffer for reuse
		buffer = buffer[:0]

		// print the board directly to the console
		for y := range board[0] {
			for x := range board {
				cell = cellEmpty
				if board[x][y] {
					cell = cellBall
					board[x][y] = false // remove the previous ball
				}
				buffer = append(buffer, cell, ' ')
			}
			buffer = append(buffer, '\n')
		}
		// print the buffer
		screen.MoveTopLeft()

		// slow down the animation
		time.Sleep(speed)
	}
}
