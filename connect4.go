package main

type connect4Game struct {
	board  c4Board
	p1, p2 c4Player
	p1turn bool
	winner *c4Player
}

func NewConnect4Game(p1 c4Player, p2 c4Player) *connect4Game {
	return &connect4Game{
		board:  newC4Board(),
		p1:     p1,
		p2:     p2,
		p1turn: true,
		winner: nil,
	}
}

func (game *connect4Game) calcWinner() {
	game.winner = nil
}

func (game *connect4Game) gameOver() bool {
	if game.winner == nil {
		game.calcWinner()
		if game.winner != nil {
			game.board.print()

			return true
		}
	}
	return false
}

func (game *connect4Game) nextMove() {
	if game.p1turn {
		game.board.playMove(game.p1.getPiece(), game.p1.getMove(&game.board))
	} else {
		game.board.playMove(game.p2.getPiece(), game.p2.getMove(&game.board))
	}
	game.p1turn = !game.p1turn
}
