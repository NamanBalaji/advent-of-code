package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	cmd string
	val int
}

func main() {
	instructions, err := getInstructionsFromInput("input.txt")
	if err != nil {
		panic(err)
	}

	cycle, cycleMap := getCycleToValueMap(instructions)

	sum := 0
	for c := 20; c <= 240; c += 40 {
		sum += c * cycleMap[c]
	}

	fmt.Println(sum)

	crt := getImage(cycle, cycleMap)

	for row := 0; row < 6; row++ {
		for column := 0; column < 40; column++ {
			fmt.Print(crt[row][column], " ")
		}
		fmt.Print("\n")
	}
}

func getInstructionsFromInput(fileName string) ([]instruction, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")

	instructions := []instruction{}

	for _, line := range lines {
		lineContent := strings.Split(line, " ")

		if lineContent[0] == "addx" {
			value, err := strconv.Atoi(lineContent[1])
			if err != nil {
				return nil, err
			}

			instructions = append(instructions, instruction{
				cmd: lineContent[0],
				val: value,
			})

			continue
		}

		instructions = append(instructions, instruction{
			cmd: lineContent[0],
		})
	}

	return instructions, nil
}

func getCycleToValueMap(instructions []instruction) (int, map[int]int) {
	x := 1
	cycle := 0
	last := 1
	cycleToX := make(map[int]int)

	cycleToX[cycle] = x

	for _, instruction := range instructions {
		if instruction.cmd == "addx" {
			cycleToX[cycle+1] = x
			cycleToX[cycle+2] = x
			x += instruction.val
			cycle += 2
			last = 2
		} else {
			cycleToX[cycle+1] = x
			cycle++
			last = 1
		}

	}

	return cycle - last, cycleToX
}

func getImage(cycle int, cycleToX map[int]int) [][]string {

	crt := make([][]string, 6)
	for i := range crt {
		crt[i] = make([]string, 40)
	}

	pos := 0
	sprite := 1
	for i := 1; i <= 240; i++ {
		sprite = cycleToX[i]
		c, r := pos%40, pos/40
		if c >= sprite-1 && c <= sprite+1 {
			crt[r][c] = "#"
		} else {
			crt[r][c] = "."
		}

		pos++
	}

	return crt
}
