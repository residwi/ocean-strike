package main

import (
	"fmt"
)

type Player struct {
	Ships   []Coordinate
	Shots   []Coordinate
	ShipMap [][]string
	Hits    int
}

type Coordinate struct {
	X, Y int
}

func NewPlayer(shipPositions [][]int, shotPositions [][]int) *Player {
	shipCoordinates := toCoordinate(shipPositions)
	shotCoordinates := toCoordinate(shotPositions)

	return &Player{Ships: shipCoordinates, Shots: shotCoordinates}
}

func (player *Player) DisplayBoard() string {
	board := ""
	for rowIndex := range player.ShipMap {
		for columnIndex := range player.ShipMap[rowIndex] {
			board = fmt.Sprintf("%v%v", board, player.ShipMap[rowIndex][columnIndex])
		}
		board = fmt.Sprintf("%v\n", board)
	}
	return board
}

func toCoordinate(positions [][]int) []Coordinate {
	var coordinate []Coordinate
	for _, posisiton := range positions {
		coordinate = append(coordinate, Coordinate{posisiton[0], posisiton[1]})
	}
	return coordinate
}

// Questions:
// is the player possible to have difference shot with the total Missile?
// should we validate it?
// what happens if the missile possition has more than the total missile?

// how about the ship posisiton?

// the number of total missile player have (0 < T < 100) but the board size is M * M which M is (0 < M < 10) which mean max M is 9, so the total 81. What happens if the total missile is 99? it is more than the board size, should we validate?
