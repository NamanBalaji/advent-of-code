package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	lines, err := getLinesFromInput("input.txt")
	if err != nil {
		panic(err)
	}

	pairGroups, err := getPairGroupsFromLine(lines)
	if err != nil {
		panic(err)
	}

	count := count(pairGroups)
	overlaps := countOverlaps(pairGroups)

	fmt.Println(count)
	fmt.Println(overlaps)
}

func getLinesFromInput(fileName string) ([]string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}

func getPairGroupsFromLine(lines []string) ([][4]int, error) {
	var groups [][4]int

	for _, line := range lines {
		var r [4]int
		_, err := fmt.Sscanf(strings.TrimSpace(line), "%d-%d,%d-%d", &r[0], &r[1], &r[2], &r[3])
		if err != nil {
			return nil, err
		}

		groups = append(groups, r)
	}

	return groups, nil
}

func count(groups [][4]int) int {
	var count int
	for _, pairs := range groups {
		if (pairs[0] <= pairs[2] && pairs[1] >= pairs[3]) || (pairs[2] <= pairs[0] && pairs[3] >= pairs[1]) {
			count++
		}
	}

	return count
}

func countOverlaps(groups [][4]int) int {
	var count int
	for _, pairs := range groups {
		if (pairs[2] <= pairs[0] && pairs[0] <= pairs[3]) ||
			(pairs[2] <= pairs[1] && pairs[1] <= pairs[3]) ||
			(pairs[0] <= pairs[2] && pairs[2] <= pairs[1]) ||
			(pairs[0] <= pairs[3] && pairs[3] <= pairs[1]) {
			count++
		}
	}

	return count
}
