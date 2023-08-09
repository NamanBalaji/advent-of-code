package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	elves, err := parseInputFromFile("input.txt")
	if err != nil {
		panic(err)
	}

	claoriesByElf, err := getTotalCaloriesPerElf(elves)
	if err != nil {
		panic(err)
	}

	maxCalorie := math.MinInt
	secondMaxCalorie := math.MinInt
	thirdMaxCalorie := math.MinInt

	for _, cal := range claoriesByElf {
		if cal > maxCalorie {
			maxCalorie = cal

			continue
		}

		if cal > secondMaxCalorie {
			secondMaxCalorie = cal

			continue
		}

		if cal > thirdMaxCalorie {
			thirdMaxCalorie = cal
		}
	}

	fmt.Printf(" %d calories is the maximum being crried. \n", maxCalorie)
	fmt.Printf(" %d calories is sum of top 3 maximum being crried. \n", maxCalorie+secondMaxCalorie+thirdMaxCalorie)
}

func parseInputFromFile(fileName string) ([]string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n\n"), nil
}

func getTotalCaloriesPerElf(elves []string) (map[int]int, error) {
	claoriesByElf := make(map[int]int)

	for i, cals := range elves {
		totalCalorieByElf := 0
		calories := strings.Split(cals, "\n")
		for _, c := range calories {
			calInt, err := strconv.Atoi(c)
			if err != nil {
				return nil, err
			}

			totalCalorieByElf += calInt
		}
		claoriesByElf[i+1] = totalCalorieByElf
	}

	return claoriesByElf, nil
}
