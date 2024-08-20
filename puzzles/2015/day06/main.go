package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Coordinates struct {
	X, Y int
}

func main() {
	lines := strings.Split(getInput(), "\n")
	fmt.Println("Part 1 solution:", part1(lines))
	fmt.Println("Part 2 solution:", part2(lines))
}

func getInput() string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(input))
}

func part1(lines []string) int {
	grid := [1000][1000]bool{}
	for _, line := range lines {
		directive, from, to := parseLine(line)
		for x := from.X; x <= to.X; x++ {
			for y := from.Y; y <= to.Y; y++ {
				switch directive {
				case "toggle":
					grid[x][y] = !grid[x][y]
				case "on":
					grid[x][y] = true
				case "off":
					grid[x][y] = false
				}
			}
		}
	}

	onCounter := 0
	for _, row := range grid {
		for _, light := range row {
			if light {
				onCounter++
			}
		}
	}

	return onCounter
}

func part2(lines []string) int {
	grid := [1000][1000]int{}
	for _, line := range lines {
		directive, from, to := parseLine(line)
		for x := from.X; x <= to.X; x++ {
			for y := from.Y; y <= to.Y; y++ {
				switch directive {
				case "toggle":
					grid[x][y] += 2
				case "on":
					grid[x][y]++
				case "off":
					grid[x][y] = max(0, grid[x][y]-1)
				}
			}
		}
	}

	totalBrightness := 0
	for _, row := range grid {
		for _, brightness := range row {
			totalBrightness += brightness
		}
	}

	return totalBrightness
}

func parseLine(line string) (string, Coordinates, Coordinates) {
	var directive string
	var from, to Coordinates
	var x, y int

	split := strings.Split(line, " ")
	if split[0] == "toggle" {
		directive = split[0]
		x, y = getCoordinates(split[1])
		from = Coordinates{x, y}
		x, y = getCoordinates(split[3])
		to = Coordinates{x, y}
	} else {
		directive = split[1]
		x, y = getCoordinates(split[2])
		from = Coordinates{x, y}
		x, y = getCoordinates(split[4])
		to = Coordinates{x, y}
	}

	return directive, from, to
}

func getCoordinates(coordinates string) (int, int) {
	xy := strings.Split(coordinates, ",")

	x, err := strconv.Atoi(xy[0])
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(xy[1])
	if err != nil {
		log.Fatal(err)
	}

	return x, y
}
