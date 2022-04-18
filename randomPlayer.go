package main

import "math/rand"

type randomPlayer struct {
	piece c4Piece
}

func (r randomPlayer) getPiece() c4Piece {
	return r.piece
}

func (r randomPlayer) getName() string {
	return "Random Choice"
}

func (r randomPlayer) getMove(b *c4Board) (col int) {
	var moves []int = b.getMoves()
	randIdx := rand.Intn(len(moves))
	return moves[randIdx]
}
