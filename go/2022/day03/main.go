package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	rustacks, err := getRucksackFromInput("input.txt")
	if err != nil {
		panic(err)
	}

	var sum1 int

	for _, rustack := range rustacks {
		common := findCommon(rustack[:len(rustack)/2], rustack[len(rustack)/2:])
		for _, c := range common {
			sum1 += getPriority(c)
		}
	}

	var sum2 int
	for i := 0; i < len(rustacks); i = i + 3 {
		common := findCommon(rustacks[i], rustacks[i+1], rustacks[i+2])
		for _, c := range common {
			sum2 += getPriority(c)
		}
	}

	fmt.Println(sum1)
	fmt.Println(sum2)
}

func getRucksackFromInput(fileName string) ([]string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}

func getPriority(char rune) int {
	if char >= 'a' {
		return int(char-'a') + 1
	}

	return int(char-'A') + 27
}

func findCommon(ss ...string) []rune {
	maps := make([]map[rune]bool, len(ss))
	var common []rune

	for i := 0; i < len(ss); i++ {
		maps[i] = make(map[rune]bool)
	}

	for i, s := range ss {
		for _, c := range s {
			maps[i][c] = true
		}
	}

	for c := range maps[0] {
		exist := true

		for i := 1; i < len(ss); i++ {
			exist = maps[i][c] && exist
		}

		if exist {
			common = append(common, c)
		}
	}

	return common
}
