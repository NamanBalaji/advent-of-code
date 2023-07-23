package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	datastream, err := getDatastreamInput("input.txt")
	if err != nil {
		panic(err)
	}

	indexPacket := getMarkerIndex(datastream, 4)
	indexMessage := getMarkerIndex(datastream, 14)

	fmt.Println(indexPacket)
	fmt.Println(indexMessage)
}

func getDatastreamInput(fileName string) ([]string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), ""), nil
}

func getMarkerIndex(data []string, distict int) int {
	i := 0
	j := distict - 1
	for j < len(data) {
		found := false
		set := make(map[string]interface{})
		for k := i; k <= j; k++ {
			_, ok := set[data[k]]
			if ok {
				found = true
				break
			}

			set[data[k]] = 1
		}

		if found {
			i++
			j++

			continue
		}

		return j + 1
	}

	panic("unreachable")
}
