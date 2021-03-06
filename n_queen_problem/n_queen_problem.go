// Given a chess board having N×N cells
// place N queens on the board in such a way that no queen attacks any other queen.

// There can be a number of possible solutions for a given length of board.
// This implementation prints only one valid solution, it can be extended to print all possible valid solutions.

package main

import "fmt"

func main() {
	boardSize := 8 // size of chess board (8 x 8)

	// make board
	board := make([][]int, boardSize)
	for i := range board {
		board[i] = make([]int, boardSize)
	}

	if PlaceQueen(&board, boardSize) {
		for i := 0; i < boardSize; i++ {
			for j := 0; j < boardSize; j++ {
				fmt.Printf("%d ", board[i][j])
			}
			fmt.Print("\n")
		}
	} else {
		fmt.Println("Not possible")
	}
}

func isAttacked(board *[][]int, x, y, boardSize int) bool {
	for i := 0; i < boardSize; i++ {
		if (*board)[x][i] == 1 {
			return true
		}
		if (*board)[i][y] == 1 {
			return true
		}
		for j := 0; j < boardSize; j++ {
			if (i-j == x-y) || (i+j == x+y) {
				if (*board)[i][j] == 1 {
					return true
				}
			}
		}
	}
	return false
}

// PlaceQueen place queens on the board and will return true if it is possible to place all the queens else false
func PlaceQueen(board *[][]int, boardSize int) bool {
	if boardSize == 0 {
		return true
	}
	for i := 0; i < len(*board); i++ {
		for j := 0; j < len(*board); j++ {
			if !(isAttacked(board, i, j, len(*board))) {
				(*board)[i][j] = 1
				if PlaceQueen(board, boardSize-1) {
					return true
				}
				(*board)[i][j] = 0
			}
		}
	}
	return false
}
