package main

import "fmt"

type c4Piece int // enum for board pieces

const (
	Empty c4Piece = iota
	OutOfBounds
	Red
	Black
)

/*
 * prints a piece, used in c4board.print()
 */
func (p c4Piece) print() {
	switch p {
	case Empty:
		fmt.Print("  ")
	case Red:
		fmt.Print("\033[31m" + "██" + "\033[0m")
	case Black:
		fmt.Print("\033[32m" + "██" + "\033[0m")
	case OutOfBounds:
		fmt.Print("OB")
	}
}

// standard connect 4 board size
const COLSIZE int = 6
const ROWSIZE int = 7
const COLS int = ROWSIZE
const ROWS int = COLSIZE

/*
 * This struct wraps an array of connect 4 pieces
 *
 * A valid game state is ensured by using the following methods
 * newC4Board: constructor
 * playMove: plays a move on a given column
 *
 */
type c4Board struct {
	arr [ROWSIZE][COLSIZE]c4Piece
}

/*
 * initializes the board to empty
 */
func newC4Board() c4Board {
	return c4Board{arr: [ROWSIZE][COLSIZE]c4Piece{}}
}

/*
 * if a given position is out of bounds for any connect 4 board
 */
func c4OutOfBounds(col int, row int) bool {
	return row < 0 || row >= ROWSIZE || col < 0 || col >= COLSIZE
}

/*
 * prints the board to console
 */
func (b *c4Board) print() {
	fmt.Println()
	for r := ROWS - 1; r >= 0; r-- {
		for c := 0; c < COLS; c++ {
			fmt.Print("|")
			b.getPiece(c, r).print()
			fmt.Print("")
		}
		fmt.Print("|")
		fmt.Println()
		for c := 0; c < COLS; c++ {
			fmt.Print("+--")
		}
		fmt.Println("+")
	}

	fmt.Print(" ")
	for i := 1; i <= COLS; i++ {
		fmt.Print(" " + fmt.Sprint(i) + " ")
	}
	fmt.Println("  ")
}

/*
 * gets a piece from the board and returns
 */
func (b *c4Board) getPiece(col int, row int) c4Piece {
	if !c4OutOfBounds(row, col) {
		return b.arr[col][row]
	} else {
		return OutOfBounds
	}
}

/*
 * sets a piece on the board and return if it was successful
 */
func (b *c4Board) setPiece(p c4Piece, col int, row int) bool {
	if !c4OutOfBounds(row, col) {
		b.arr[col][row] = p
		return true
	} else {
		return false
	}
}

/*
 * plays a move on the board and return if it was successful
 */
func (b *c4Board) playMove(p c4Piece, col int) bool {
	row := b.getValidPos(col)
	return b.setPiece(p, col, row)
}

/*
 * returns a copy of the board with a given move played
 */
func (b c4Board) tryMove(p c4Piece, col int) c4Board {
	row := b.getValidPos(col)
	b.setPiece(p, col, row)
	return b
}

/*
 * returns the height (row number) that a piece can be played in a column
 */
func (b *c4Board) getValidPos(col int) int {
	for height := 0; true; height++ {
		switch b.getPiece(col, height) {
		case Red, Black:
			continue

		case OutOfBounds:
			return -1

		case Empty:
			return height
		}
	}
	return -1
}

/*
 * Returns the possible columns to make a move
 */
func (b *c4Board) getMoves() []int {
	moves := make([]int, 0, COLS)
	for col := 0; col < COLS; col++ {
		pos := b.getValidPos(col)
		if pos != -1 {
			moves = append(moves, col)
			continue
		}
	}
	return moves
}

/*
 * If there is 4 in a row with a Red or Black piece, return which
 * otherwise return Empty
 */
func (b *c4Board) check4Row() c4Piece {
	// horizontalCheck
	for i := 0; i < COLS; i++ {
		for j := 0; j < ROWS-3; j++ {
			piece := b.getPiece(i, j)
			if piece != Red && piece != Black {
				continue
			}
			if b.getPiece(i, j) == piece && b.getPiece(i, j+1) == piece && b.getPiece(i, j+2) == piece && b.getPiece(i, j+3) == piece {
				return piece
			}
		}
	}

	// verticalCheck
	for i := 0; i < COLS-3; i++ {
		for j := 0; j < ROWS; j++ {
			piece := b.getPiece(i, j)
			if piece != Red && piece != Black {
				continue
			}
			if b.getPiece(i, j) == piece && b.getPiece(i+1, j) == piece && b.getPiece(i+2, j) == piece && b.getPiece(i+3, j) == piece {
				return piece
			}
		}
	}

	// ascendingDiagonalCheck
	for i := 0; i < COLS-3; i++ {
		for j := 0; j < ROWS-3; j++ {
			piece := b.getPiece(i, j)
			if piece != Red && piece != Black {
				continue
			}
			if (b.getPiece(i, j) == piece) && (b.getPiece(i+1, j+1) == piece) && (b.getPiece(i+2, j+2) == piece) && (b.getPiece(i+3, j+3) == piece) {
				return piece
			}
		}
	}

	// descendingDiagonalCheck
	for i := 0; i < COLS-3; i++ {
		for j := 3; j < ROWS; j++ {
			piece := b.getPiece(i, j)
			if piece != Red && piece != Black {
				continue
			}
			if (b.getPiece(i, j) == piece) && (b.getPiece(i+1, j-1) == piece) && (b.getPiece(i+2, j-2) == piece) && (b.getPiece(i+3, j-3) == piece) {
				return piece
			}
		}
	}
	return Empty
}
