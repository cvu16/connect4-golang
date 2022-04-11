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

// 6 rows, 7 columns
type c4Board struct {
	arr [6][7]c4Piece
}

func newC4Board() c4Board {
	return c4Board{arr: [6][7]c4Piece{}}
}

func c4OutOfBounds(row int, col int) bool {
	return !(row >= 0 && row < 6 && col >= 0 && col < 7)
}

func printPiece(p c4Piece) {
	switch p {
	case Empty:
		fmt.Print("_")
	case Red:
		fmt.Print("\033[31m" + "█" + "\033[0m")
	case Black:
		fmt.Print("\033[30m" + "█" + "\033[0m")
	case OutOfBounds:
		fmt.Print("X")
	}
}

func (b *c4Board) print() {
	fmt.Println()
	for r := 5; r >= 0; r-- {
		for c := 0; c < 7; c++ {
			fmt.Print("| ")
			printPiece(b.getPiece(r, c))
			fmt.Print(" ")
		}
		fmt.Print("|")
		fmt.Println()
	}
	fmt.Println("  1   2   3   4   5   6   7  ")
	fmt.Println()
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
	moves := make([]int, 0, 7)
	for col := 0; col < 7; col++ {
		pos := b.getValidPos(col)
		if pos != -1 {
			moves = append(moves, col)
			continue
		}
	}
	return moves
}
