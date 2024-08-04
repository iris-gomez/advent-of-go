package main

import (
	"fmt"
	"log"
	"os"
)

type Santa struct {
	X int
	Y int
}

func (s *Santa) Move(direction byte) {
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
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Part 1 solution: %d\n", part1(input))
	fmt.Printf("Part 2 solution: %d\n", part2(input))
}

func part1(input []byte) int {
	santa := Santa{0, 0}

	houses := make(map[string]int)
	houses[fmt.Sprintf("x%dy%d", santa.X, santa.Y)] += 1

	for _, char := range input {
		santa.Move(char)
		houses[fmt.Sprintf("x%dy%d", santa.X, santa.Y)] += 1
	}

	return len(houses)
}

func part2(input []byte) int {
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
