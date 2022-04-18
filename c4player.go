package main

import (
	"fmt"
	"strings"
)

func getPlayers() (p1 c4Player, p2 c4Player) {

	fmt.Println("Select: ")
	fmt.Println(" (1) Singleplayer")
	fmt.Println(" (2) Multiplayer")
	fmt.Print("> ")

	var input string
	fmt.Scanln(&input)
	input = strings.TrimSpace(input)
	switch input {
	case "1", "s", "S", "y", "Y", "":
		fmt.Println("No bitches :-|")
		return human{"Alice", Black}, randomPlayer{Red}
		//return human{"Alice", Black}, perfectAI{"Bob", Red}

	case "2", "m", "M", "n", "N":
		// fmt.Println("Multiplayer")
		return human{"Alice", Black}, human{"Bob", Red}
	default:
		// fmt.Println("Singleplayer")
		return human{"Alice", Black}, perfectAI{Red}
	}
}

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
