package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	input, err := ParseInput(string(data))
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	player1 := NewPlayer(input.Player1ShipPositions, input.Player1ShotPositions)
	player2 := NewPlayer(input.Player2ShipPositions, input.Player2ShotPositions)

	game := NewGame(input.BoardSize)
	game.AddPlayer1(player1)
	game.AddPlayer2(player2)
	err = game.Start()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	result := game.PrintResult()

	os.WriteFile("output.txt", []byte(result), 0644)
}
