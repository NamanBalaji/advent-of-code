package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func main() {
	lines, err := getCommandLogFromInput("input.txt")
	if err != nil {
		panic(err)
	}

	dirsizes := getDirectorySizes(lines)

	var sum int
	for _, sz := range dirsizes {
		if sz <= 100000 {
			sum += sz
		}
	}

	fmt.Println(sum)

	toFree := dirsizes["/"] - 40000000
	best := dirsizes["/"]
	for _, sz := range dirsizes {
		if sz >= toFree && sz < best {
			best = sz
		}
	}

	fmt.Println(best)
}

func getCommandLogFromInput(fileName string) ([]string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}

func getDirectorySizes(lines []string) map[string]int {
	dirsizes := make(map[string]int)

	var dir string
	for i := 0; i < len(lines); {
		fields := strings.Fields(lines[i])
		switch fields[1] {
		case "cd":
			dir = path.Clean(path.Join(dir, fields[2]))
			i++
		case "ls":
			for i++; i < len(lines) && lines[i][0] != '$'; i++ {
				fields := strings.Fields(lines[i])
				if fields[0] != "dir" {
					for d := dir; ; d = path.Dir(d) {
						size, _ := strconv.Atoi(fields[0])
						dirsizes[d] += size
						if d == "/" {
							break
						}
					}
				}
			}
		}
	}

	return dirsizes
}
