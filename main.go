package main

import "fmt"

func main() {
	fmt.Println("Welcome to Connect 4 in Golang!")
	fmt.Println()

	player1, player2 := getPlayers()

	c4Game := newConnect4Game(player1, player2)

	for !c4Game.isGameOver() {
		c4Game.nextMove()
	}

	fmt.Println(c4Game.winner.getName() + " has won!")
}
