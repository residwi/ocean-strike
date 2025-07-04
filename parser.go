package main

import (
	"strconv"
	"strings"
)

type Input struct {
	BoardSize            int
	TotalShips           int
	Player1ShipPositions [][]int
	Player1ShotPositions [][]int
	TotalMissile         int
	Player2ShipPositions [][]int
	Player2ShotPositions [][]int
}

func ParseInput(data string) (Input, error) {
	lines := strings.Split(data, "\n")

	boardSize, err := strconv.ParseInt(lines[0], 10, 64)
	if err != nil {
		return Input{}, err
	}

	totalShips, err := strconv.ParseInt(lines[1], 10, 64)
	if err != nil {
		return Input{}, err
	}

	player1ShipPositions, err := parsePosition(lines[2])
	if err != nil {
		return Input{}, err
	}

	player1ShotPositions, err := parsePosition(lines[3])
	if err != nil {
		return Input{}, err
	}

	totalMissile, err := strconv.ParseInt(lines[4], 10, 64)
	if err != nil {
		return Input{}, err
	}

	player2ShipPositions, err := parsePosition(lines[5])
	if err != nil {
		return Input{}, err
	}

	player2ShotPositions, err := parsePosition(lines[6])
	if err != nil {
		return Input{}, err
	}

	input := Input{
		BoardSize:            int(boardSize),
		TotalShips:           int(totalShips),
		Player1ShipPositions: player1ShipPositions,
		Player1ShotPositions: player1ShotPositions,
		TotalMissile:         int(totalMissile),
		Player2ShipPositions: player2ShipPositions,
		Player2ShotPositions: player2ShotPositions,
	}
	return input, nil
}

func parsePosition(strPositions string) ([][]int, error) {
	var intArrPosition [][]int
	for position := range strings.SplitSeq(strPositions, ":") {
		xyPosition := strings.Split(strings.TrimSpace(position), ",")
		x, err := strconv.ParseInt(xyPosition[0], 10, 64)
		if err != nil {
			return nil, err
		}

		y, err := strconv.ParseInt(xyPosition[1], 10, 64)
		if err != nil {
			return nil, err
		}

		intArrPosition = append(intArrPosition, []int{int(x), int(y)})
	}

	return intArrPosition, nil
}
