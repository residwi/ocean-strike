package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameStart(t *testing.T) {
	player1ShipPositions := [][]int{
		{0, 0},
		{1, 1},
		{2, 2},
	}

	player1ShotPositions := [][]int{
		{0, 2},
		{1, 1},
		{2, 2},
	}

	player2ShipPositions := [][]int{
		{0, 2},
		{1, 1},
		{2, 0},
	}

	player2ShotPositions := [][]int{
		{0, 0},
		{1, 1},
		{2, 0},
	}

	t.Run("should render player1 board result", func(t *testing.T) {
		player1 := NewPlayer(player1ShipPositions, player1ShotPositions)
		player2 := NewPlayer(player2ShipPositions, player2ShotPositions)

		game := NewGame(3)
		game.AddPlayer1(player1)
		game.AddPlayer2(player2)
		game.Start()

		expectedPlayer1ShipMap := [][]string{
			{DEAD, "_", "_"},
			{"_", DEAD, "_"},
			{MISSED, "_", ALIVE},
		}
		assert.Equal(t, expectedPlayer1ShipMap, player1.ShipMap)
		assert.Equal(t, 2, player1.Hits)
	})

	t.Run("should render player2 board result", func(t *testing.T) {
		player1 := NewPlayer(player1ShipPositions, player1ShotPositions)
		player2 := NewPlayer(player2ShipPositions, player2ShotPositions)

		game := NewGame(3)
		game.AddPlayer1(player1)
		game.AddPlayer2(player2)
		game.Start()

		expectedPlayer2ShipMap := [][]string{
			{"_", "_", DEAD},
			{"_", DEAD, "_"},
			{ALIVE, "_", MISSED},
		}
		assert.Equal(t, expectedPlayer2ShipMap, player2.ShipMap)
		assert.Equal(t, 2, player2.Hits)
	})

	t.Run("should error when board size is not within 1 to 9", func(t *testing.T) {
		game := NewGame(0)
		errMin := game.Start()
		assert.EqualError(t, errMin, "Board size should be between 1 to 9")

		testGame := NewGame(10)
		errMax := testGame.Start()
		assert.EqualError(t, errMax, "Board size should be between 1 to 9")
	})
}

func TestScoreResult(t *testing.T) {
	t.Run("Player 1 wins", func(t *testing.T) {
		player1 := &Player{Hits: 2}
		player2 := &Player{Hits: 1}

		game := NewGame(0)
		game.AddPlayer1(player1)
		game.AddPlayer2(player2)

		assert.Equal(t, "Player 1 wins", game.scoreResult())
	})

	t.Run("Player 2 wins", func(t *testing.T) {
		player1 := &Player{Hits: 0}
		player2 := &Player{Hits: 1}

		game := NewGame(0)
		game.AddPlayer1(player1)
		game.AddPlayer2(player2)

		assert.Equal(t, "Player 2 wins", game.scoreResult())
	})

	t.Run("It is a draw", func(t *testing.T) {
		player1 := &Player{Hits: 1}
		player2 := &Player{Hits: 1}

		game := NewGame(0)
		game.AddPlayer1(player1)
		game.AddPlayer2(player2)

		assert.Equal(t, "It is a draw", game.scoreResult())
	})
}

func TestInitializeBoard(t *testing.T) {
	expected := [][]string{
		{"_", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_"},
		{"_", "_", "_", "_", "_"},
	}
	assert.Equal(t, expected, initializeBoard(5))
}
