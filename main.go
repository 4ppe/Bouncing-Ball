package main

import (
	"fmt"
)

func main() {

	// var xVelocity, yVelocity int

	const (
		width  = 50
		height = 10

		cellEmpty = ' '
		cellBall  = '⚾'
	)

	var cell rune // Current cell (for caching)

	// Create the board
	board := make([][]bool, width)
	for column := range board {
		board[column] = make([]bool, height)
	}

	// A drawing buffer
	buffer := make([]rune, 0, width*height)

	board[12][2] = true // Random test location

	buffer = buffer[:0] // reuse buffer

	// Print the board directly to the console
	for y := range board[0] {
		for x := range board {
			cell = cellEmpty
			if board[x][y] {
				cell = cellBall
			}
			//fmt.Print(string(cell), " ")
			buffer = append(buffer, cell, ' ')
		}
		buffer = append(buffer, '\n')
	}
	fmt.Print(string(buffer))

	// for {
	// 	screen.Clear()
	// 	time.Sleep(time.Second / 20)
	// 	screen.MoveTopLeft()
	// 	//...
	// }
}
