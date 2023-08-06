package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	rounds1    = 20
	rounds2    = 10000
	decreaseBy = 3
)

type monkeyOp struct {
	items     []int
	operation func(int) int
	divisor   int
	ifTrue    int
	ifFalse   int
}

func main() {
	monkeyOps, err := getMoneyOpsFromInput("input.txt")
	if err != nil {
		panic(err)
	}

	//fmt.Println(getMonkeyBuisness(monkeyOps, rounds1, false))
	fmt.Println(getMonkeyBuisness(monkeyOps, rounds2, true))
}

func getMonkeyBuisness(monkeyOps []monkeyOp, rounds int, worrying bool) int {
	productOfDivisors := 1
	for i := range monkeyOps {
		productOfDivisors *= monkeyOps[i].divisor
	}

	buisness := make([]int, len(monkeyOps))

	for i := 0; i < rounds; i++ {
		for i, monkey := range monkeyOps {
			buisness[i] += len(monkey.items)

			for _, item := range monkeyOps[i].items {
				worry := monkey.operation(item)

				if !worrying {
					worry /= decreaseBy
				}

				worry %= productOfDivisors

				next := monkeyOps[i].ifFalse
				if worry%monkeyOps[i].divisor == 0 {
					next = monkeyOps[i].ifTrue
				}

				monkeyOps[next].items = append(monkeyOps[next].items, worry)
			}

			monkeyOps[i].items = monkeyOps[i].items[:0]
		}

	}

	sort.Sort(sort.Reverse(sort.IntSlice(buisness)))

	return buisness[0] * buisness[1]
}

func getMoneyOpsFromInput(filename string) ([]monkeyOp, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	result := []monkeyOp{}

	ops := strings.Split(string(data), "\n\n")

	for _, op := range ops {
		var ifTrue, ifFalse, divisible int
		var operation, operand string

		lines := strings.Split(op, "\n")

		_, err = fmt.Sscanf(lines[2], "   Operation: new = old %s %s", &operation, &operand)
		if err != nil {
			return nil, err
		}

		_, err = fmt.Sscanf(lines[3], "  Test: divisible by %d", &divisible)
		if err != nil {
			return nil, err
		}

		_, err = fmt.Sscanf(lines[4], "    If true: throw to monkey %d", &ifTrue)
		if err != nil {
			return nil, err
		}

		_, err = fmt.Sscanf(lines[5], "    If false: throw to monkey %d", &ifFalse)
		if err != nil {
			return nil, err
		}

		startingItemsString := strings.Split(lines[1], ": ")
		sIs := strings.Split(startingItemsString[1], ", ")

		startingItems := []int{}

		for _, i := range sIs {
			item, _ := strconv.Atoi(i)
			startingItems = append(startingItems, item)
		}

		result = append(result, monkeyOp{
			items:     startingItems,
			divisor:   divisible,
			ifTrue:    ifTrue,
			ifFalse:   ifFalse,
			operation: getOperationFunction(operand, operation),
		})
	}

	return result, nil
}

func getOperationFunction(operand, operation string) func(int) int {
	var num int

	if operand != "old" {
		num, _ = strconv.Atoi(operand)
	}

	if operation == "*" {
		if operand == "old" {
			return func(old int) int { return old * old }
		}

		return func(old int) int { return num * old }
	}

	if operation == "+" {
		if operand == "old" {
			return func(old int) int { return old + old }
		}
	}

	return func(old int) int { return num + old }
}
