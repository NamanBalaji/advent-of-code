package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}

func part1() int {
	pairs, err := getPairsFromInput("input.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	for i, pair := range pairs {
		if diff(pair[0], pair[1]) < 0 {
			sum += i + 1
		}
	}

	return sum
}

func part2() int {
	pairs, err := getPairsFromInput("input.txt")
	if err != nil {
		panic(err)
	}

	allPackets := [][]interface{}{
		{[]interface{}{float64(2)}},
		{[]interface{}{float64(6)}},
	}

	for _, pair := range pairs {
		allPackets = append(allPackets, pair[0])
		allPackets = append(allPackets, pair[1])
	}

	sort.Slice(allPackets, func(i, j int) bool {
		left, right := allPackets[i], allPackets[j]
		diff := diff(left, right)

		return diff < 1
	})

	product := 1

	for i, p := range allPackets {
		if fmt.Sprint(p) == "[[2]]" || fmt.Sprint(p) == "[[6]]" {
			product *= i + 1
		}
	}

	return product
}

func getPairsFromInput(filename string) ([][2][]interface{}, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	input := make([][2][]interface{}, 0)

	pairs := strings.Split(string(data), "\n\n")

	for _, pair := range pairs {
		ps := strings.Split(pair, "\n")

		input = append(input, [2][]interface{}{
			getPacketFromString(ps[0]),
			getPacketFromString(ps[1]),
		})
	}

	return input, nil
}

func getPacketFromString(str string) []interface{} {
	packet := []interface{}{}
	json.Unmarshal([]byte(str), &packet)

	return packet
}

func diff(left, right interface{}) int {

	leftNum, isLeftNum := left.(float64)
	rightNum, isRightNum := right.(float64)

	if isLeftNum {
		if isRightNum {
			return int(leftNum) - int(rightNum)
		} else {
			return diff([]interface{}{left}, right)
		}
	} else {
		if isRightNum {
			return diff(left, []interface{}{right})
		}
	}

	leftList := left.([]interface{})
	rightList := right.([]interface{})

	merged := merge(leftList, rightList)

	for _, pair := range merged {
		d := diff(pair[0], pair[1])
		if d != 0 {
			return d
		}
	}

	return len(leftList) - len(rightList)
}

func merge(slices ...[]interface{}) [][]interface{} {
	minLen := len(slices[0])
	for _, slice := range slices {
		if len(slice) < minLen {
			minLen = len(slice)
		}
	}

	merged := make([][]interface{}, minLen)
	for i := 0; i < minLen; i++ {
		merged[i] = make([]interface{}, len(slices))
		for j, slice := range slices {
			merged[i][j] = slice[i]
		}
	}

	return merged
}
