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
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")

	fmt.Printf("Part 1 solution: %d\n", part1(lines))
	fmt.Printf("Part 2 solution: %d\n", part2(lines))
}

func part1(instructions []string) int {
	grid := [1000][1000]bool{}
	for _, instruction := range instructions {
		directive, from, to := parseInstruction(instruction)
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

func part2(instructions []string) int {
	grid := [1000][1000]int{}
	for _, instruction := range instructions {
		directive, from, to := parseInstruction(instruction)
		for x := from.X; x <= to.X; x++ {
			for y := from.Y; y <= to.Y; y++ {
				switch directive {
				case "toggle":
					grid[x][y] += 2
				case "on":
					grid[x][y]++
				case "off":
					grid[x][y]--
					if grid[x][y] < 0 {
						grid[x][y] = 0
					}
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

func parseInstruction(instruction string) (string, Coordinates, Coordinates) {
	var directive string
	var from, to Coordinates
	var x, y int

	parts := strings.Split(instruction, " ")
	if parts[0] == "toggle" {
		directive = parts[0]
		x, y = splitCoordinates(parts[1])
		from = Coordinates{x, y}
		x, y = splitCoordinates(parts[3])
		to = Coordinates{x, y}
	} else {
		directive = parts[1]
		x, y = splitCoordinates(parts[2])
		from = Coordinates{x, y}
		x, y = splitCoordinates(parts[4])
		to = Coordinates{x, y}
	}

	return directive, from, to
}

func splitCoordinates(coordinates string) (int, int) {
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
