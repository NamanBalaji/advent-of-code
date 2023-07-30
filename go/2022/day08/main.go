package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	forest, err := getForestFromInput("input.txt")
	if err != nil {
		panic(err)
	}

	visible := makeVisibleMatrix(forest)

	calculateVisible(forest, visible)

	count := 0

	for _, row := range visible {
		for _, r := range row {
			if r {
				count++
			}
		}
	}

	fmt.Println(count)
	fmt.Println(bestScenicScore(forest))
}

func getForestFromInput(filename string) ([][]int, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")

	forest := make([][]int, len(lines))

	for i := 0; i < len(lines); i++ {
		row := strings.Split(lines[i], "")
		forest[i] = make([]int, len(row))

		for j, r := range row {
			forest[i][j], _ = strconv.Atoi(r)
		}
	}

	return forest, nil
}

func calculateVisible(forest [][]int, visible [][]bool) {
	// row wise r->l
	for i := 1; i < len(forest)-1; i++ {
		max := forest[i][0]
		for j := 1; j < len(forest[i])-1; j++ {
			if forest[i][j] > max {
				visible[i][j] = true
				max = forest[i][j]
			}
		}
	}

	// row wise l->r
	for i := 1; i < len(forest)-1; i++ {
		max := forest[i][len(forest[i])-1]
		for j := len(forest[i]) - 2; j > 0; j-- {
			if forest[i][j] > max {
				visible[i][j] = true
				max = forest[i][j]
			}
		}
	}

	//column wise t->b
	for j := 1; j < len(forest[0])-1; j++ {
		max := forest[0][j]
		for i := 1; i < len(forest)-1; i++ {
			if forest[i][j] > max {
				visible[i][j] = true
				max = forest[i][j]
			}
		}
	}

	//column wise b->t
	for j := 1; j < len(forest[0])-2; j++ {
		max := forest[len(forest[0])-1][j]
		for i := len(forest) - 2; i > 0; i-- {
			if forest[i][j] > max {
				visible[i][j] = true
				max = forest[i][j]
			}
		}
	}
}

func makeVisibleMatrix(forest [][]int) [][]bool {
	visible := make([][]bool, len(forest))

	for i := range forest {
		visible[i] = make([]bool, len(forest[i]))
		for j := range forest[i] {
			if i == 0 || i == len(forest)-1 || j == 0 || j == len(forest[0])-1 {
				visible[i][j] = true
			}
		}
	}

	return visible
}

func bestScenicScore(forest [][]int) int {
	bestScore := 0

	for row := range forest {
		for col := range forest[row] {
			score := scenicScore(forest, row, col)
			if score > bestScore {
				bestScore = score
			}
		}
	}

	return bestScore
}

func scenicScore(forest [][]int, row, col int) int {
	return visibleToRight(forest, row, col) *
		visibleToLeft(forest, row, col) *
		visibleToBottom(forest, row, col) *
		visibleToTop(forest, row, col)
}

func visibleToRight(forest [][]int, row, col int) int {
	count := 0
	maxHeight := forest[row][col]

	for col++; col < len(forest[row]); col++ {
		count++
		if forest[row][col] >= maxHeight {
			break
		}
	}

	return count
}

func visibleToLeft(forest [][]int, row, col int) int {
	count := 0
	maxHeight := forest[row][col]

	for col--; col >= 0; col-- {
		count++
		if forest[row][col] >= maxHeight {
			break
		}
	}

	return count
}

func visibleToBottom(forest [][]int, row, col int) int {
	count := 0
	maxHeight := forest[row][col]

	for row++; row < len(forest); row++ {
		count++
		if forest[row][col] >= maxHeight {
			break
		}
	}

	return count
}

func visibleToTop(forest [][]int, row, col int) int {
	count := 0
	maxHeight := forest[row][col]

	for row--; row >= 0; row-- {
		count++
		if forest[row][col] >= maxHeight {
			break
		}
	}

	return count
}
