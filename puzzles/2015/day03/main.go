package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Santa struct {
	X int
	Y int
}

func (s *Santa) Move(direction rune) {
	switch direction {
	case '^':
		s.Y += 1
	case 'v':
		s.Y -= 1
	case '>':
		s.X += 1
	case '<':
		s.X -= 1
	}
}

func main() {
	input := getInput()
	fmt.Println("Part 1 solution:", part1(input))
	fmt.Println("Part 2 solution:", part2(input))
}

func getInput() string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	return strings.TrimSpace(string(input))
}

func part1(input string) int {
	santa := Santa{0, 0}

	houses := make(map[string]int)
	houses[fmt.Sprintf("x%dy%d", santa.X, santa.Y)] += 1

	for _, char := range input {
		santa.Move(char)
		houses[fmt.Sprintf("x%dy%d", santa.X, santa.Y)] += 1
	}

	return len(houses)
}

func part2(input string) int {
	santa, robo := Santa{0, 0}, Santa{0, 0}

	houses := make(map[string]int)
	houses[fmt.Sprintf("x%dy%d", santa.X, santa.Y)] += 1
	houses[fmt.Sprintf("x%dy%d", robo.X, robo.Y)] += 1

	for i, char := range input {
		if i%2 == 0 {
			santa.Move(char)
			houses[fmt.Sprintf("x%dy%d", santa.X, santa.Y)] += 1
		} else {
			robo.Move(char)
			houses[fmt.Sprintf("x%dy%d", robo.X, robo.Y)] += 1
		}
	}

	return len(houses)
}
