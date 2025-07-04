package main

import (
	"errors"
	"fmt"
)

const (
	ALIVE  = "B"
	DEAD   = "X"
	MISSED = "O"
)

type Game struct {
	BoardSize int
	Player1   *Player
	Player2   *Player
}

func NewGame(boardSize int) *Game {
	return &Game{BoardSize: boardSize}
}

func (game *Game) AddPlayer1(player *Player) {
	game.Player1 = player
}

func (game *Game) AddPlayer2(player *Player) {
	game.Player2 = player
}

func (game *Game) Start() error {
	if game.BoardSize < 1 || game.BoardSize > 9 {
		return errors.New("Board size should be between 1 to 9")
	}

	board := initializeBoard(game.BoardSize)
	game.Player1.ShipMap = board
	boardPlayer2 := initializeBoard(game.BoardSize)
	game.Player2.ShipMap = boardPlayer2

	for _, shipPosition := range game.Player1.Ships {
		game.Player1.ShipMap[shipPosition.X][shipPosition.Y] = ALIVE
	}

	for _, shipPosition := range game.Player2.Ships {
		game.Player2.ShipMap[shipPosition.X][shipPosition.Y] = ALIVE
	}

	// start shooting
	for _, shotPosition := range game.Player1.Shots {
		status := game.Player2.ShipMap[shotPosition.X][shotPosition.Y]
		switch status {
		case ALIVE:
			game.Player2.ShipMap[shotPosition.X][shotPosition.Y] = DEAD
			game.Player1.Hits += 1
		case DEAD:
			game.Player2.ShipMap[shotPosition.X][shotPosition.Y] = DEAD
		default:
			game.Player2.ShipMap[shotPosition.X][shotPosition.Y] = MISSED
		}
	}

	for _, shotPosition := range game.Player2.Shots {
		status := game.Player1.ShipMap[shotPosition.X][shotPosition.Y]
		switch status {
		case ALIVE:
			game.Player1.ShipMap[shotPosition.X][shotPosition.Y] = DEAD
			game.Player2.Hits += 1
		case DEAD:
			game.Player1.ShipMap[shotPosition.X][shotPosition.Y] = DEAD
		default:
			game.Player1.ShipMap[shotPosition.X][shotPosition.Y] = MISSED
		}
	}

	return nil
}

func (game *Game) PrintResult() string {
	return fmt.Sprintf(
		"Player1\n%v\n\nPlayer2\n%v\n\nP1:%v\nP2:%v\n%v",
		game.Player1.DisplayBoard(),
		game.Player2.DisplayBoard(),
		game.Player1.Hits,
		game.Player2.Hits,
		game.scoreResult(),
	)
}

func (game *Game) scoreResult() string {
	if game.Player1.Hits > game.Player2.Hits {
		return "Player 1 wins"
	} else if game.Player1.Hits < game.Player2.Hits {
		return "Player 2 wins"
	} else {
		return "It is a draw"
	}
}

func initializeBoard(boardSize int) [][]string {
	board := make([][]string, boardSize)

	for rowIndex := range boardSize {
		board[rowIndex] = make([]string, boardSize)
		for columnIndex := range boardSize {
			board[rowIndex][columnIndex] = "_"
		}
	}

	return board
}
