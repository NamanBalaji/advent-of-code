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

	for !dropSand(grid, lowestCol, highestCol, highestRow) {
		ans++
	}

	fmt.Println(ans)

	resetGrid(grid)
	ans2 := 0
	for !dropSand(grid, 0, 999, highestRow+2) {
		ans2++
		if grid[0][500] == "o" {
			break
		}
	}

	fmt.Println(ans2)
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

	// as part 2 column tends to infinity
	colNum := 1000
	rowNum := highestRow + 3

	grid := make([][]string, rowNum)

	for i := 0; i < rowNum; i++ {
		grid[i] = make([]string, colNum)

		for j := 0; j < colNum; j++ {
			if 500 == j && i == 0 {
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

func resetGrid(grid [][]string) {
	for i, row := range grid {
		for j := range row {
			if grid[i][j] != "#" {
				grid[i][j] = "."
			}
		}
	}
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

	for j := range grid[len(grid)-1] {
		grid[len(grid)-1][j] = "#"
	}
}

func getSortedCoordinate(a, b int) (int, int) {
	if a-b > 0 {
		return b, a
	}

	return a, b
}

func dropSand(grid [][]string, lc, hc, hr int) bool {
	r, c := 0, 500

	for r < hr {
		rowBelow := r + 1
		colRight := c + 1
		colLeft := c - 1

		if grid[rowBelow][c] == "." {
			r++
		} else if colLeft >= lc && grid[rowBelow][colLeft] == "." {
			r++
			c--
		} else if colRight <= hc && grid[rowBelow][colRight] == "." {
			r++
			c++
		} else if colLeft < lc {
			return true
		} else if colRight > hc {
			return true
		} else {
			grid[r][c] = "o"
			return false
		}
	}

	return true
}
