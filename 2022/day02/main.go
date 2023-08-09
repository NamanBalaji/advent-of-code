package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	rock = iota + 1
	paper
	scissor
)

type round struct {
	opp int
	me  int
}

type roundNew struct {
	opp    int
	result string
}

func main() {

	data, err := readFile("input.txt")
	if err != nil {
		panic(err)
	}

	rounds := getRoundsForGame1(string(data))
	roundsNew := getNewRoundsForGame2(string(data))

	var totalScoreR1 int
	var totalScoreR2 int

	for _, r := range rounds {
		totalScoreR1 += r.play()
	}

	for _, r := range roundsNew {
		totalScoreR2 += r.play()
	}

	fmt.Printf("total score after playing 1st game is %d \n", totalScoreR1)
	fmt.Printf("total score after playing 2nd game is %d \n", totalScoreR2)
}

func readFile(fileName string) ([]byte, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func getRoundsForGame1(data string) []round {
	rounds := make([]round, 0)

	roundsString := strings.Split(string(data), "\n")
	for _, rs := range roundsString {
		choices := strings.Split(string(rs), " ")

		rounds = append(rounds, round{
			opp: getValueFromChoice(choices[0]),
			me:  getValueFromChoice(choices[1]),
		})
	}

	return rounds
}

func getNewRoundsForGame2(data string) []roundNew {
	rounds := make([]roundNew, 0)

	roundsString := strings.Split(string(data), "\n")
	for _, rs := range roundsString {
		choices := strings.Split(string(rs), " ")

		rounds = append(rounds, roundNew{
			opp:    getValueFromChoice(choices[0]),
			result: choices[1],
		})
	}

	return rounds
}

func getValueFromChoice(letter string) int {
	switch letter {
	case "A":
		return 1
	case "B":
		return 2
	case "C":
		return 3
	case "X":
		return 1
	case "Y":
		return 2
	case "Z":
		return 3
	}

	return -1
}

func (r round) play() int {
	if r.opp == rock {
		if r.me == rock {
			return r.me + 3
		} else if r.me == paper {
			return r.me + 6
		} else {
			return r.me
		}
	} else if r.opp == paper {
		if r.me == rock {
			return r.me
		} else if r.me == paper {
			return r.me + 3
		} else {
			return r.me + 6
		}
	} else {
		if r.me == rock {
			return r.me + 6
		} else if r.me == paper {
			return r.me
		} else {
			return r.me + 3
		}
	}
}

func (r roundNew) play() int {
	if r.opp == rock {
		if r.result == "X" {
			return scissor
		} else if r.result == "Y" {
			return rock + 3
		} else {
			return paper + 6
		}
	} else if r.opp == paper {
		if r.result == "X" {
			return rock
		} else if r.result == "Y" {
			return paper + 3
		} else {
			return scissor + 6
		}
	} else {
		if r.result == "X" {
			return paper
		} else if r.result == "Y" {
			return scissor + 3
		} else {
			return rock + 6
		}
	}
}
