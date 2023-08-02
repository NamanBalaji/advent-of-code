package main

import (
	"fmt"
	"os"
	"strings"
)

type motion struct {
	coordinates coordinates
	count       int
}

type coordinates struct {
	x, y int
}

func main() {
	moves, err := getMovesFromInput("input.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println(traceMotion(2, moves))
	fmt.Println(traceMotion(10, moves))

}

func getMovesFromInput(filename string) ([]motion, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")
	motions := make([]motion, 0)

	for _, l := range lines {
		var coord coordinates
		var dir string
		var count int
		_, err := fmt.Sscanf(l, "%s %d", &dir, &count)
		if err != nil {
			return nil, err
		}

		switch dir {
		case "U":
			coord = coordinates{x: 0, y: 1}
		case "D":
			coord = coordinates{x: 0, y: -1}
		case "R":
			coord = coordinates{x: 1, y: 0}
		case "L":
			coord = coordinates{x: -1, y: 0}
		default:
			return nil, err
		}

		motions = append(motions, motion{
			coordinates: coord,
			count:       count,
		})
	}

	return motions, nil
}

func traceMotion(knotNumber int, moves []motion) int {
	knots := make([]coordinates, knotNumber)
	for i := 0; i < knotNumber; i++ {
		knots[i] = coordinates{x: 0, y: 0}
	}

	unique := make(map[coordinates]int, 0)

	for _, m := range moves {
		for i := 0; i < m.count; i++ {
			moveHead(&knots[0], m.coordinates)
			for i := 1; i < knotNumber; i++ {
				moveTail(&knots[i-1], &knots[i])
			}
			unique[knots[knotNumber-1]] = 1
		}
	}

	return len(unique)
}

func moveTail(head, tail *coordinates) {
	diffX := head.x - tail.x
	diffY := head.y - tail.y

	if diffY == 0 && abs(diffX) > 1 {
		op := -1
		if diffX > 0 {
			op = 1
		}
		tail.x = tail.x + op

		return

	} else if diffX == 0 && abs(diffY) > 1 {
		op := -1
		if diffY > 0 {
			op = 1
		}
		tail.y = tail.y + op

		return

	} else if abs(diffX) > 1 || abs(diffY) > 1 {
		opX := -1
		if diffX > 0 {
			opX = 1
		}

		opY := -1
		if diffY > 0 {
			opY = 1
		}

		tail.x = tail.x + opX
		tail.y = tail.y + opY

	}

}

func moveHead(head *coordinates, move coordinates) {
	head.x = head.x + move.x
	head.y = head.y + move.y
}

func abs(n int) int {
	if n < 0 {
		return ((-1) * n)
	}

	return n
}
