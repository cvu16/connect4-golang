package main

import (
	"math/rand"
	"sync"
)

type minimaxAI struct {
	piece c4Piece
}

func (ai minimaxAI) getName() string {
	return "Computer"
}

func (ai minimaxAI) getPiece() c4Piece {
	return ai.piece
}

func (ai minimaxAI) getOpponentPiece() c4Piece {
	if ai.getPiece() == Black {
		return Red
	} else if ai.getPiece() == Red {
		return Black
	}
	return Empty
}

// func printPiece(p c4Piece) {
// 	switch p {
// 	case Empty:
// 		fmt.Println("empty")
// 	case Red:
// 		fmt.Println("red")
// 	case Black:
// 		fmt.Println("black")
// 	default:
// 		fmt.Println("Unknown Piece")
// 	}
// }

func (ai minimaxAI) scoreBoard(board c4Board) int {
	piece := board.check4Row()

	// printPiece(piece)
	// printPiece(ai.getPiece())
	// printPiece(ai.getOpponentPiece())
	// fmt.Println()

	if piece == ai.getPiece() {
		return 100
	} else if piece == ai.getOpponentPiece() {
		return -100
	} else {
		return 0
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (ai minimaxAI) minimax(board c4Board, depth int, alpha int, beta int, maximizingPlayer bool) int {
	score := ai.scoreBoard(board)
	if depth == 10 || score == 100 || score == -100 {
		return score - depth
	}

	var bestVal int
	if maximizingPlayer {
		bestVal = -101
		for _, move := range board.getMoves() {
			nextBoard := board.tryMove(ai.getPiece(), move)
			//nextBoard.print()
			value := ai.minimax(nextBoard, depth+1, alpha, beta, false)
			bestVal = max(bestVal, value)
			alpha = max(alpha, bestVal)
			if beta <= alpha {
				break
			}
		}
		return bestVal
	} else {
		bestVal = 101
		for _, move := range board.getMoves() {
			nextBoard := board.tryMove(ai.getOpponentPiece(), move)
			//nextBoard.print()
			value := ai.minimax(nextBoard, depth+1, alpha, beta, true)
			bestVal = min(bestVal, value)
			beta = min(beta, bestVal)
			if beta <= alpha {
				break
			}
		}
		return bestVal
	}
}

func pickMove(moves, scores []int) int {
	var max_score int = scores[0]
	var best_moves []int
	for _, score := range scores {
		if score > max_score {
			max_score = score
		}
	}

	for i, move := range moves {
		if scores[i] == max_score {
			best_moves = append(best_moves, move)
		}
	}

	return best_moves[rand.Intn(len(best_moves))]
}

func (ai minimaxAI) getMove(b *c4Board) (col int) {
	moves := b.getMoves()
	// fmt.Print("moves: ")
	// fmt.Println(moves)
	var scores = make([]int, len(moves))
	//Adding condition variables
	var wg sync.WaitGroup
	wg.Add(len(moves))
	//Implement multiple threads here for to shorten computational time
	for i := 0; i < len(moves); i++ {

		go func(i int) {
			defer wg.Done()
			tryBoard := b.tryMove(ai.getPiece(), moves[i])
			scores[i] = ai.minimax(tryBoard, 0, -101, 101, false)
		}(i)

	}
	wg.Wait()
	/*
		for i := 0; i < len(moves); i++ {
			tryBoard := b.tryMove(ai.getPiece(), moves[i])
			scores[i] = ai.minimax(tryBoard, 0, -101, 101, false)
		}
	*/

	return pickMove(moves, scores)
}
