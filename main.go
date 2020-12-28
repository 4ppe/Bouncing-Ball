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

	board[12][2] = true // Random test location

	// Print the board directly to the console
	for y := range board[0] {
		for x := range board {
			cell = cellEmpty
			if board[x][y] {
				cell = cellBall
			}
			fmt.Print(string(cell), " ")
		}
		fmt.Println()
	}

	// for {
	// 	screen.Clear()
	// 	time.Sleep(time.Second / 20)
	// 	screen.MoveTopLeft()
	// 	//...
	// }
}
