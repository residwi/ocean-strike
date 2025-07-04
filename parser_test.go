package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParserInput(t *testing.T) {
	battleShipData := "5\n5\n1,1 : 2,0 : 2,3 : 3,4 : 4,3\n0,1 : 2,3 : 3,0 : 3,4 : 4,1\n5\n0,1 : 4,3 : 2,3 : 3,1 : 4,1\n0,1 : 0,0 : 1,1 : 2,3 : 4,3"
	input, err := ParseInput(battleShipData)
	assert.NoError(t, err)

	expectedInput := Input{
		BoardSize:  5,
		TotalShips: 5,
		Player1ShipPositions: [][]int{
			{1, 1},
			{2, 0},
			{2, 3},
			{3, 4},
			{4, 3},
		},
		Player1ShotPositions: [][]int{
			{0, 1},
			{2, 3},
			{3, 0},
			{3, 4},
			{4, 1},
		},
		TotalMissile: 5,
		Player2ShipPositions: [][]int{
			{0, 1},
			{4, 3},
			{2, 3},
			{3, 1},
			{4, 1},
		},
		Player2ShotPositions: [][]int{
			{0, 1},
			{0, 0},
			{1, 1},
			{2, 3},
			{4, 3},
		},
	}

	assert.Equal(t, expectedInput, input)
}
