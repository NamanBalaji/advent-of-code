package main

import (
	"fmt"
	"os"
	"strings"
)

type Stack []string

type move struct {
	amount int
	from   int
	to     int
}

func main() {
	lines, err := getLinesFromInput("input.txt")
	if err != nil {
		panic(err)
	}
	stacks := getStacks(lines[0])

	moves, err := getMoves(lines[1])
	if err != nil {
		panic(err)
	}

	for _, m := range moves {
		m.performMove(stacks)
	}

	answer := ""

	for _, s := range stacks {
		element, ok := s.Peek()
		if !ok {
			answer += ""
		}

		answer += element
	}

	fmt.Println(answer)

	partTwo()
}

func partTwo() {
	lines, err := getLinesFromInput("input.txt")
	if err != nil {
		panic(err)
	}
	stacks := getStacks(lines[0])

	moves, err := getMoves(lines[1])
	if err != nil {
		panic(err)
	}

	for _, m := range moves {
		m.performBetterMove(stacks)
	}

	answer := ""

	for _, s := range stacks {
		element, ok := s.Peek()
		if !ok {
			answer += ""
		}

		answer += element
	}

	fmt.Println(answer)
}

func getLinesFromInput(fileName string) ([]string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n\n"), nil
}

func getStacks(data string) []Stack {
	lines := strings.Split(string(data), "\n")
	lines = lines[:len(lines)-1]

	stacks := make([]Stack, 0)

	for _, line := range lines {
		split := make([]string, 0)
		for i := 0; i < len(line); i += 4 {
			if string(line[i:i+3]) == "   " {
				split = append(split, "#")
			} else {
				split = append(split, line[i:i+3])
			}
		}

		if len(stacks) == 0 {
			for i := 0; i < len(split); i++ {
				stack := make([]string, 0)
				stacks = append(stacks, stack)
			}
		}

		for i, v := range split {
			if v == "#" {
				continue
			}
			stacks[i] = append(stacks[i], string([]rune(v)[1]))
		}
	}

	for _, s := range stacks {
		s.reverse()
	}

	return stacks
}

func getMoves(data string) ([]move, error) {
	var moves []move
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		m := move{}
		_, err := fmt.Sscanf(strings.TrimSpace(line), "move %d from %d to %d", &m.amount, &m.from, &m.to)
		if err != nil {
			return nil, err
		}

		moves = append(moves, m)
	}

	return moves, nil
}

func (m move) performBetterMove(stacks []Stack) {
	var elements []string

	for i := 0; i < m.amount; i++ {
		element, ok := stacks[m.from-1].Pop()
		if !ok {
			panic("empty stack")
		}

		elements = append(elements, element)
	}

	for i := len(elements) - 1; i >= 0; i-- {
		stacks[m.to-1].Push(elements[i])
	}
}

func (m move) performMove(stacks []Stack) {
	for i := 0; i < m.amount; i++ {
		element, ok := stacks[m.from-1].Pop()
		if !ok {
			panic("empty stack")
		}

		stacks[m.to-1].Push(element)
	}
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(char string) {
	*s = append(*s, char)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}
	index := len(*s) - 1
	element := (*s)[index]
	*s = (*s)[:index]

	return element, true
}

func (s *Stack) Peek() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	index := len(*s) - 1
	return (*s)[index], true
}

func (s *Stack) reverse() {
	l := 0
	r := len(*s) - 1

	for l < r {
		(*s)[l], (*s)[r] = (*s)[r], (*s)[l]
		l++
		r--
	}
}
