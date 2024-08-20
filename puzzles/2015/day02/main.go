package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
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
	totalSquareFeet := 0
	for _, line := range lines {
		l, w, h := parseLine(line)

		boxSquareFeet := surface(l, w, h)
		areas := []int{area(l, w), area(w, h), area(l, h)}

		totalSquareFeet += boxSquareFeet + slices.Min(areas)
	}

	return totalSquareFeet
}

func part2(lines []string) int {
	totalRibbonFeet := 0
	for _, line := range lines {
		l, w, h := parseLine(line)

		volume := volume(l, w, h)
		perimeters := []int{perimeter(l, w), perimeter(w, h), perimeter(l, h)}

		totalRibbonFeet += volume + slices.Min(perimeters)
	}

	return totalRibbonFeet
}

func parseLine(line string) (int, int, int) {
	dimensions := strings.Split(line, "x")

	l, err := strconv.Atoi(dimensions[0])
	if err != nil {
		log.Fatal(err)
	}

	w, err := strconv.Atoi(dimensions[1])
	if err != nil {
		log.Fatal(err)
	}

	h, err := strconv.Atoi(dimensions[2])
	if err != nil {
		log.Fatal(err)
	}

	return l, w, h
}

func surface(l, w, h int) int {
	return (2 * l * w) + (2 * w * h) + (2 * h * l)
}

func area(x, y int) int {
	return x * y
}

func volume(l, w, h int) int {
	return l * w * h
}

func perimeter(x, y int) int {
	return (x * 2) + (y * 2)
}
