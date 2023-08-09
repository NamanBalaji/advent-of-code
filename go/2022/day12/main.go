package main

import (
	"fmt"
	"os"
	"strings"
)

type node struct {
	hops     int
	position pos
}

type pos struct {
	row, col int
}

type queue []node

var coords = [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func main() {
	grid, err := getGridFromInput("input.txt")
	if err != nil {
		panic(err)
	}

	var (
		sr, er, sc, ec int
	)

	for r, row := range grid {
		for c := range row {
			if grid[r][c] == "S" {
				grid[r][c] = "a"
				sr, sc = r, c
			}

			if grid[r][c] == "E" {
				grid[r][c] = "z"
				er, ec = r, c
			}
		}
	}

	fmt.Println(findHops(grid, sr, sc, false, er, ec, ""))
	fmt.Println(findHops(grid, er, ec, true, -1, -1, "a"))
}

func findHops(grid [][]string, sr, sc int, unknownEnd bool, er, ec int, endPeak string) int {
	visited := make(map[pos]bool, 0)

	position := pos{
		row: sr,
		col: sc,
	}

	visited[position] = true

	q := make(queue, 0)
	q.push(node{
		hops:     0,
		position: position,
	})

	for !q.isEmpty() {
		n := q.pop()
		for _, coord := range coords {
			nr := n.position.row + coord[0]
			nc := n.position.col + coord[1]

			if nr < 0 || nr >= len(grid) || nc < 0 || nc >= len(grid[0]) {
				continue
			}

			position := pos{
				row: nr, col: nc,
			}

			if visited[position] {
				continue
			}

			if !validHop(grid[nr][nc], grid[n.position.row][n.position.col], unknownEnd) {
				continue
			}

			if unknownEnd {
				if grid[nr][nc] == endPeak {
					return n.hops + 1
				}
			} else {
				if nr == er && nc == ec {
					return n.hops + 1
				}
			}

			visited[position] = true
			q.push(node{
				hops:     n.hops + 1,
				position: position,
			})

		}
	}

	return -1
}

func getGridFromInput(filename string) ([][]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	grid := [][]string{}
	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		row := []string{}
		chars := strings.Split(line, "")
		row = append(row, chars...)
		grid = append(grid, row)
	}

	return grid, nil
}

func (q *queue) push(n node) {
	*q = append(*q, n)
}

func (q *queue) pop() node {
	n := (*q)[0]
	*q = (*q)[1:]

	return n
}

func (q *queue) isEmpty() bool {
	return len(*q) == 0
}

func validHop(s1, s2 string, unknown bool) bool {
	char1 := []rune(s1)
	char2 := []rune(s2)

	if unknown {
		return int(char1[0])-int(char2[0]) >= -1
	}

	return int(char1[0])-int(char2[0]) <= 1
}
