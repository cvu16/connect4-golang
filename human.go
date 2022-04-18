package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
