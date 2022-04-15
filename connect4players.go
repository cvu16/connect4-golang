package main

import (
	"fmt"
	"strconv"
	"strings"
)

type c4Player interface {
	getPiece() c4Piece
	getMove(b *c4Board) int
	getName() string
}

type human struct {
	name  string
	piece c4Piece
}

func (h human) getPiece() c4Piece {
	return h.piece
}

func (h human) getName() string {
	return h.name
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

// func (p *c4Player) scoreBoard(b *c4Board) {

// }

func (h human) getMove(b *c4Board) (col int) {
	b.print()
	var moves []int = b.getMoves()
	fmt.Println("Enter column to place piece")
	printMoves(moves)

	var input string
	for {
		fmt.Print(h.getName() + " > ")
		fmt.Scanln(&input)
		input = strings.TrimSpace(input)
		col, err := strconv.Atoi(input)
		col-- // 0 indexed

		if err != nil {
			println("Unrecognized Input: " + input)
			printMoves(moves)
			break
		}

		if containsMove(moves, col) {
			return col
		} else {
			println("Invalid Move")
			printMoves(moves)
			continue
		}
	}
	return
}

type perfectAI struct {
	piece c4Piece
}

func (ai perfectAI) getPiece() c4Piece {
	return ai.piece
}

func (ai perfectAI) getName() string {
	return "Perfect AI"
}

func (ai perfectAI) getMove(b *c4Board) (col int) {
	return 0
}

func (ai perfectAI) minimax(board c4Board, depth int, alpha int, beta int, maximizingPlayer bool) {
	winner := board.getWinner()
	if winner == {
		return
	}
	if maximizingPlayer {
		value := Math.MinInt32
		for _, move := range getMoves(board) {
			value := Math.max(value, alphabeta(child, depth-1, alpha, beta, FALSE))
			alpha := Math.max(alpha, value)
			if value >= beta {
				break
			}
		}
		return value
	} else {
		value := Math.MaxInt32
		for _, move := range getMoves(board) {
			value := Math.min(value, alphabeta(child, depth-1, alpha, beta, TRUE))
			beta := Math.min(beta, value)
			if value <= alpha {
				break
			}
			return value
		}
	}
}
