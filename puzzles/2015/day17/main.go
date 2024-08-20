package main

import (
	"fmt"
	"log"
	"os"
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
	containers := getContainers(lines)
	combinations := 0
	for i := 1; i < 1<<len(containers); i++ {
		storedEggnog := 0
		for j := 0; j < len(containers); j++ {
			if i&(1<<j) != 0 {
				storedEggnog += containers[j]
			}
		}

		if storedEggnog == 150 {
			combinations++
		}
	}

	return combinations
}

func part2(lines []string) int {
	containers := getContainers(lines)
	usage := make([]int, len(containers))
	for i := 1; i < 1<<len(containers); i++ {
		storedEggnog, containersUsed := 0, 0
		for j := 0; j < len(containers); j++ {
			if i&(1<<j) != 0 {
				storedEggnog += containers[j]
				containersUsed++
			}
		}

		if storedEggnog == 150 {
			usage[containersUsed]++
		}
	}

	var variations int
	for i := 0; i < len(usage); i++ {
		if usage[i] != 0 {
			variations = usage[i]
			break
		}
	}

	return variations
}

func getContainers(lines []string) []int {
	containers := []int{}
	for _, line := range lines {
		container, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		containers = append(containers, container)
	}
	return containers
}
