package main

import (
	"fmt"
	"time"
)

/*
 * This struct is the main abstraction of a gamestate of connect4
 *
 * It is designed to be accessed through the following methods
 * NewConnect4Game: constructor
 * isGameOver: 		used in the game loop to determine if another move will be played
 * nextMove: 		requests the next move from the player who's turn it is
 */
type connect4Game struct {
	board    c4Board
	p1, p2   c4Player
	p1turn   bool
	winner   c4Player
	gameOver bool
}

/*
 * Returns a pointer to a fresh game
 */
func newConnect4Game(p1 c4Player, p2 c4Player) *connect4Game {
	return &connect4Game{
		board:    newC4Board(),
		p1:       p1,
		p2:       p2,
		p1turn:   true,
		winner:   nil,
		gameOver: false,
	}
}

/*
 * checks if the game was determined to be over (game.gameOver),
 * if there are 4 in a row, or if there are no moves left
 * also sets the winner appropriately or to nil if tie
 */
func (game *connect4Game) isGameOver() bool {
	if game.gameOver { // game already over
		return true
	}

	piece := game.board.check4Row()
	if piece != Empty { // check for 4 in a row
		game.gameOver = true
		if piece == game.p1.getPiece() {
			game.winner = game.p1
		} else if piece == game.p2.getPiece() {
			game.winner = game.p2
		} else { // just in case
			game.winner = nil
		}
	} else if len(game.board.getMoves()) == 0 { // no moves left
		game.gameOver = true
	}

	return game.gameOver
}

func (game *connect4Game) nextMove() {
	var player c4Player
	if game.p1turn {
		player = game.p1
	} else {
		player = game.p2
	}

	for i := 0; i < 10; i++ {
		start := time.Now()
		if game.board.playMove(player.getPiece(), player.getMove(&game.board)) {
			end := time.Since(start)
			fmt.Printf("%s %s \n", "Move took: ", end)
			game.p1turn = !game.p1turn
			return
		} else {
			fmt.Println("Invalid move")
		}
	}

	game.endGame()
	fmt.Println("Too many invalid moves, ending game")
}

func (game *connect4Game) endGame() {
	game.gameOver = true
}
