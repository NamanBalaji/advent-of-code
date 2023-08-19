package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type coordinate struct {
	x, y int
}

type interval struct {
	start, end int
}

var infoString = "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d"

const Y = 10

func main() {
	sensors, beacons, err := getSensorsAndBeaconsFromInput("input.txt")
	if err != nil {
		panic(err)
	}

	dist := getDistances(sensors, beacons)

	fmt.Println(part1(sensors, beacons, dist))
	fmt.Println(part2(sensors, beacons, dist))
}

func part1(sensors, beacons []coordinate, dist []int) int {
	minX := math.MaxInt
	maxX := math.MinInt

	intervals := make([]interval, 0)

	for i, sensor := range sensors {
		dx := dist[i] - abs(Y-sensor.y)

		if dx <= 0 {
			continue
		}

		start := sensor.x - dx
		if start < minX {
			minX = start
		}

		end := sensor.x + dx
		if end > maxX {
			maxX = end
		}

		intervals = append(intervals, interval{start, end})
	}

	allowedX := make(map[int]bool)
	for _, beacon := range beacons {
		if beacon.y == Y {
			allowedX[beacon.x] = true
		}
	}

	ans := 0

	for i := minX; i <= maxX; i++ {
		if allowedX[i] {
			continue
		}

		for _, interval := range intervals {
			if interval.start <= i && interval.end >= i {
				ans++

				break
			}
		}
	}

	return ans
}

func part2(sensors, beacons []coordinate, dist []int) int {
	posLines := make([]int, 0)
	negLines := make([]int, 0)

	pos := 0
	neg := 0

	n := len(sensors)

	for i, sensor := range sensors {
		d := dist[i]

		negLines = append(negLines, sensor.x+sensor.y-d, sensor.x+sensor.y+d)
		posLines = append(posLines, sensor.x-sensor.y-d, sensor.x-sensor.y+d)
	}

	for i := 0; i < 2*n; i++ {
		for j := i + 1; j < 2*n; j++ {
			a, b := posLines[i], posLines[j]
			if abs(a-b) == 2 {
				pos = min(a, b) + 1
			}
			a, b = negLines[i], negLines[j]
			if abs(a-b) == 2 {
				neg = min(a, b) + 1
			}
		}
	}

	x, y := (pos+neg)/2, (neg-pos)/2
	ans := x*4000000 + y

	return ans
}

func getSensorsAndBeaconsFromInput(filename string) ([]coordinate, []coordinate, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}

	sensors := make([]coordinate, 0)
	beacons := make([]coordinate, 0)

	for _, line := range strings.Split(string(data), "\n") {
		var sx, sy, bx, by int
		_, err := fmt.Sscanf(line, infoString, &sx, &sy, &bx, &by)
		if err != nil {
			return nil, nil, err
		}

		sensors = append(sensors, coordinate{x: sx, y: sy})
		beacons = append(beacons, coordinate{x: bx, y: by})
	}

	return sensors, beacons, nil
}

func getDistances(sensors, beacons []coordinate) []int {
	dist := make([]int, 0)

	for i, sensor := range sensors {
		dist = append(dist, getDistance(sensor, beacons[i]))
	}

	return dist
}

func getDistance(c1, c2 coordinate) int {
	x := abs(c1.x - c2.x)
	y := abs(c1.y - c2.y)

	return x + y
}

func abs(x int) int {
	if x < 0 {
		x *= -1
	}

	return x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
