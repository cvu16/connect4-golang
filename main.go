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
		var input2 string
		fmt.Println("Write your name: ")
		fmt.Print("> ")
		fmt.Scanln(&input2)
		input2 = strings.TrimSpace(input2)

		fmt.Println("")
		return human{input2, Black}, minimaxAI{Red}

	case "2", "m", "M", "n", "N":
		// fmt.Println("Multiplayer")
		var input2 string
		fmt.Println("Write your name (player1): ")
		fmt.Print("> ")
		fmt.Scanln(&input2)
		input2 = strings.TrimSpace(input2)

		var input3 string
		fmt.Println("Write your name (player2): ")
		fmt.Print("> ")
		fmt.Scanln(&input3)
		input3 = strings.TrimSpace(input3)

		return human{input2, Black}, human{input3, Red}
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
