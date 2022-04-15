package main

import "fmt"

type c4Piece int // enum for board pieces

const (
	Empty c4Piece = iota
	OutOfBounds
	Red
	Black
)

// const (
// 	p1Piece c4Piece = Black
// 	p2Piece         = Red
// )

const ROWS int = 6
const COLS int = 7

// 6 rows, 7 columns
type c4Board struct {
	arr [ROWS][COLS]c4Piece
}

func newC4Board() c4Board {
	return c4Board{arr: [ROWS][COLS]c4Piece{}}
}

func c4OutOfBounds(row int, col int) bool {
	return !(row >= 0 && row < ROWS && col >= 0 && col < COLS)
}

func printPiece(p c4Piece) {
	switch p {
	case Empty:
		fmt.Print("__")
	case Red:
		fmt.Print("\033[31m" + "██" + "\033[0m")
	case Black:
		fmt.Print("\033[30m" + "██" + "\033[0m")
	case OutOfBounds:
		fmt.Print("X")
	}
}

func (b *c4Board) print() {
	fmt.Println()
	for r := ROWS - 1; r >= 0; r-- {
		for c := 0; c < COLS; c++ {
			fmt.Print("|")
			printPiece(b.getPiece(r, c))
			fmt.Print("")
		}
		fmt.Print("|")
		fmt.Println()
	}

	fmt.Print(" ")
	for i := 1; i <= COLS; i++ {
		fmt.Print(" " + fmt.Sprint(i) + " ")
	}
	fmt.Println("  ")
}

func (b *c4Board) getPiece(row int, col int) c4Piece {
	if c4OutOfBounds(row, col) {
		return OutOfBounds
	} else {
		return b.arr[row][col]
	}
}

func (b *c4Board) setPiece(p c4Piece, row int, col int) bool {
	if c4OutOfBounds(row, col) {
		return false
	} else {
		b.arr[row][col] = p
		return true
	}
}

func (b *c4Board) playMove(p c4Piece, col int) bool {
	row := b.getValidPos(col)

	if row == -1 {
		return false
	} else {
		return b.setPiece(p, row, col)
	}
}

func (b *c4Board) getValidPos(col int) int {
	for height := 0; true; height++ {
		switch b.getPiece(height, col) {
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

func (b *c4Board) getWinningPiece() c4Piece {
	// horizontalCheck
	for j := 0; j < ROWS-3; j++ {
		for i := 0; i < COLS; i++ {
			piece := b.getPiece(i, j)
			if piece != Red && piece != Black {
				goto VERT
			}
			if b.getPiece(i, j) == piece && b.getPiece(i, j+1) == piece && b.getPiece(i, j+2) == piece && b.getPiece(i, j+3) == piece {
				return piece
			}
		}
	}
VERT:
	// verticalCheck
	for i := 0; i < ROWS-3; i++ {
		for j := 0; j < COLS; j++ {
			piece := b.getPiece(i, j)
			if piece != Red && piece != Black {
				goto ADIAG
			}
			if b.getPiece(i, j) == piece && b.getPiece(i+1, j) == piece && b.getPiece(i+2, j) == piece && b.getPiece(i+3, j) == piece {
				return piece
			}
		}
	}

ADIAG:
	// ascendingDiagonalCheck
	for i := 0; i < COLS-3; i++ {
		for j := 0; j < ROWS-3; j++ {
			piece := b.getPiece(i, j)
			if piece != Red && piece != Black {
				goto DDIAG
			}
			if (b.getPiece(i, j) == piece) && (b.getPiece(i+1, j+1) == piece) && (b.getPiece(i+2, j+2) == piece) && (b.getPiece(i+3, j+3) == piece) {
				return piece
			}
		}
	}
DDIAG:
	// descendingDiagonalCheck
	for i := 0; i < COLS; i++ {
		for j := 3; j < ROWS; j++ {
			piece := b.getPiece(i, j)
			if piece != Red && piece != Black {
				return Empty
			}
			if (b.getPiece(i, j) == piece) && (b.getPiece(i+1, j-1) == piece) && (b.getPiece(i+2, j-2) == piece) && (b.getPiece(i+3, j-3) == piece) {
				return piece
			}
		}
	}
	return Empty
}
