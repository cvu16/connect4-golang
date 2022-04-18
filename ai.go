package main

type perfectAI struct {
	piece c4Piece
}

func (ai perfectAI) getName() string {
	return "Perfect AI"
}

func (ai perfectAI) getPiece() c4Piece {
	return ai.piece
}

func (ai perfectAI) getOpponentPiece() c4Piece {
	if ai.getPiece() == Black {
		return Red
	} else if ai.getPiece() == Red {
		return Black
	}
	return Empty
}

func (ai perfectAI) scoreBoard(board c4Board) int {
	piece := board.check4Row()
	if piece == ai.getPiece() {
		return 1
	} else if piece == ai.getOpponentPiece() {
		return -1
	} else {
		return 0
	}
}

func max(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func (ai perfectAI) minimax(board c4Board, depth int, alpha int, beta int, maximizingPlayer bool) int {
	// board.print()
	score := ai.scoreBoard(board)
	if depth == 0 || score != 0 {
		return score
	}

	if maximizingPlayer {
		value := -100
		for _, move := range board.getMoves() {
			value = max(value, ai.minimax(board.tryMove(ai.getOpponentPiece(), move), depth-1, alpha, beta, false))
			alpha = max(alpha, value)
			if value >= beta {
				break
			}
		}
	} else {
		value := 100
		for _, move := range board.getMoves() {
			value = min(value, ai.minimax(board.tryMove(ai.getPiece(), move), depth-1, alpha, beta, true))
			beta = min(beta, value)
			if value <= alpha {
				break
			}
		}
	}
	return score
}

func (ai perfectAI) getMove(b *c4Board) (col int) {
	moves := b.getMoves()
	best_move := moves[0]
	best_score := ai.minimax(b.tryMove(ai.getPiece(), best_move), 10, 2, -2, true)
	for i, move := range moves {
		if i == 0 {
			continue
		}
		score := ai.minimax(b.tryMove(ai.getPiece(), best_move), 10, 2, -2, true)
		if score > best_score {
			best_score = score
			best_move = move
		}
	}

	return best_move
}
