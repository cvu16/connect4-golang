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
		return human{"Alice", Black}, minimaxAI{Red}

	case "2", "m", "M", "n", "N":
		// fmt.Println("Multiplayer")
		return human{"Alice", Black}, human{"Bob", Red}
	}

	return human{"Alice", Black}, human{"Bob", Red}
}

func main() {
	fmt.Println("Welcome to Connect 4 in Golang!")
	fmt.Println()

	player1, player2 := getPlayers()

	c4Game := newConnect4Game(player1, player2)

	for !c4Game.isGameOver() {
		c4Game.nextMove()
	}

	c4Game.board.print()
	if c4Game.winner != nil {
		fmt.Println(c4Game.winner.getName() + " has won!")
	} else {
		fmt.Println("It was a tie!")
	}
}
