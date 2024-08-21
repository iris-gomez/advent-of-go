package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

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
	matrix := getMatrix(lines)
	for i := 0; i < 100; i++ {
		matrix = updateState(matrix, false)
	}

	lightsOn := 0
	for _, row := range matrix {
		for _, light := range row {
			if light {
				lightsOn++
			}
		}
	}

	return lightsOn
}

func part2(lines []string) int {
	matrix := getMatrix(lines)

	matrix[0][0] = true
	matrix[0][99] = true
	matrix[99][0] = true
	matrix[99][99] = true

	for i := 0; i < 100; i++ {
		matrix = updateState(matrix, true)
	}

	lightsOn := 0
	for _, row := range matrix {
		for _, light := range row {
			if light {
				lightsOn++
			}
		}
	}

	return lightsOn
}

func getMatrix(lines []string) [100][100]bool {
	matrix := [100][100]bool{}
	for i, line := range lines {
		lights := strings.Split(line, "")
		for j, light := range lights {
			switch light {
			case "#":
				matrix[i][j] = true
			case ".":
				matrix[i][j] = false
			}
		}
	}

	return matrix
}

func updateState(matrix [100][100]bool, cornersAlwaysOn bool) [100][100]bool {
	updated := matrix
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if cornersAlwaysOn && isCorner(i, j) {
				updated[i][j] = true
				continue
			}

			onNeighbors := 0
			for k := max(0, i-1); k < min(100, i+2); k++ {
				for l := max(0, j-1); l < min(100, j+2); l++ {
					if k == i && l == j {
						continue
					}
					if matrix[k][l] {
						onNeighbors++
					}
				}
			}

			switch matrix[i][j] {
			case true:
				if onNeighbors != 2 && onNeighbors != 3 {
					updated[i][j] = false
				}
			case false:
				if onNeighbors == 3 {
					updated[i][j] = true
				}
			}
		}
	}

	return updated
}

func isCorner(x, y int) bool {
	return (x == 0 || x == 99) && (y == 0 || y == 99)
}
