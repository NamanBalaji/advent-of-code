package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

type pair struct {
	coord1, coord2 *coord
}

var highestCol = -1
var highestRow = -1
var lowestCol = 100000000

func main() {
	grid, err := getGridFromInput("input.txt")
	if err != nil {
		panic(err)
	}

	ans := 0

	for !dropSand(grid) {
		ans++
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			fmt.Printf("%s ", grid[i][j])
		}
		fmt.Println()
	}

	fmt.Println(ans)
}

func getGridFromInput(filename string) ([][]string, error) {

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	coordinates := []*coord{}
	pairs := []pair{}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		coords := strings.Split(line, " -> ")
		c1 := strings.Split(coords[0], ",")

		for i := 1; i < len(coords); i++ {
			coord1 := evalCoordinateString(c1)
			coordinates = append(coordinates, coord1)

			c2 := strings.Split(coords[i], ",")

			coord2 := evalCoordinateString(c2)
			coordinates = append(coordinates, coord2)

			pairs = append(pairs, pair{coord1: coord1, coord2: coord2})

			c1 = c2
		}

	}

	// each row's last index will be highestCol - lowestCol
	colNum := highestCol - lowestCol + 1
	rowNum := highestRow + 1
	// get adjusted coordinates
	for _, c := range coordinates {
		c.y = c.y - lowestCol
	}

	grid := make([][]string, rowNum)

	for i := 0; i < rowNum; i++ {
		grid[i] = make([]string, colNum)

		for j := 0; j < colNum; j++ {
			if 500-lowestCol == j && i == 0 {
				grid[i][j] = "+"
			} else {
				grid[i][j] = "."
			}
		}
	}

	fillRocks(grid, pairs)

	return grid, nil
}

func getInt(s string) int {
	num, _ := strconv.Atoi(s)

	return num
}

func evalCoordinateString(str []string) *coord {
	x := getInt(str[1])
	y := getInt(str[0])

	if x > highestRow {
		highestRow = x
	}

	if y > highestCol {
		highestCol = y
	}

	if y < lowestCol {
		lowestCol = y
	}

	return &coord{x, y}
}

func fillRocks(grid [][]string, pairs []pair) {
	for _, p := range pairs {
		if p.coord1.x == p.coord2.x {
			s, b := getSortedCoordinate(p.coord1.y, p.coord2.y)
			for i := s; i <= b; i++ {
				grid[p.coord1.x][i] = "#"
			}
		} else if p.coord1.y == p.coord2.y {
			s, b := getSortedCoordinate(p.coord1.x, p.coord2.x)
			for i := s; i <= b; i++ {
				grid[i][p.coord1.y] = "#"
			}
		}
	}
}

func getSortedCoordinate(a, b int) (int, int) {
	if a-b > 0 {
		return b, a
	}

	return a, b
}

func dropSand(grid [][]string) bool {
	r, c := 0, 500-lowestCol

	for r < len(grid)-1 {
		rowBelow := r + 1
		colRight := c + 1
		colLeft := c - 1

		if grid[rowBelow][c] == "." {
			r++
		} else if colLeft >= 0 && grid[rowBelow][colLeft] == "." {
			r++
			c--
		} else if colRight < len(grid[0]) && grid[rowBelow][colRight] == "." {
			r++
			c++
		} else if colLeft < 0 {
			return true
		} else if colRight >= len(grid[0]) {
			return true
		} else {
			grid[r][c] = "o"

			return false
		}
	}

	return true
}
