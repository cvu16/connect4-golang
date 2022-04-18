package main

import (
	"fmt"
)

type c4Player interface {
	getPiece() c4Piece
	getMove(b *c4Board) int
	getName() string
}

func printMoves(moves []int) {
	fmt.Print("Moves: ")
	for _, mv := range moves {
		fmt.Print(mv + 1)
		fmt.Print(" ")
	}
	fmt.Println()
}

func containsMove(moves []int, move int) bool {
	for _, m := range moves {
		if m == move {
			return true
		}
	}
	return false
}
