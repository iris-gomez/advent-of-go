package main

import (
	"fmt"
	"log"
	"math"
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
	return shortestPath(makeMatrix(lines))
}

func part2(lines []string) int {
	return longestPath(makeMatrix(lines))
}

func makeMatrix(lines []string) [][]int {
	locations := make(map[string]int)
	for _, line := range lines {
		split := strings.Split(line, " ")
		if _, ok := locations[split[0]]; !ok {
			locations[split[0]] = len(locations)
		}
		if _, ok := locations[split[2]]; !ok {
			locations[split[2]] = len(locations)
		}
	}

	matrix := make([][]int, len(locations))
	for i := range matrix {
		matrix[i] = make([]int, len(locations))
	}

	for _, line := range lines {
		split := strings.Split(line, " ")

		weight, err := strconv.Atoi(split[4])
		if err != nil {
			log.Fatal(err)
		}

		matrix[locations[split[0]]][locations[split[2]]] = weight
		matrix[locations[split[2]]][locations[split[0]]] = weight
	}

	return matrix
}

func shortestPath(matrix [][]int) int {
	minWeight := math.MaxInt

	for source := 0; source < len(matrix); source++ {
		var vertices []int
		for i := 0; i < len(matrix); i++ {
			if i != source {
				vertices = append(vertices, i)
			}
		}

		for _, permutation := range getPermutations(vertices) {
			currentWeight := 0

			j := source
			for _, i := range permutation {
				currentWeight += matrix[j][i]
				j = i
			}

			minWeight = min(minWeight, currentWeight)
		}
	}

	return minWeight
}

func longestPath(matrix [][]int) int {
	maxWeight := 0

	for source := 0; source < len(matrix); source++ {
		var vertices []int
		for i := 0; i < len(matrix); i++ {
			if i != source {
				vertices = append(vertices, i)
			}
		}

		for _, permutation := range getPermutations(vertices) {
			currentWeight := 0

			j := source
			for _, i := range permutation {
				currentWeight += matrix[j][i]
				j = i
			}

			maxWeight = max(maxWeight, currentWeight)
		}
	}

	return maxWeight
}

func getPermutations(a []int) [][]int {
	var heap func(int, []int)
	result := [][]int{}

	heap = func(k int, a []int) {
		if k == 1 {
			temp := make([]int, len(a))
			copy(temp, a)
			result = append(result, temp)
		} else {
			heap(k-1, a)

			for i := 0; i < k-1; i++ {
				if k%2 == 0 {
					a[i], a[k-1] = a[k-1], a[i]
				} else {
					a[0], a[k-1] = a[k-1], a[0]
				}
				heap(k-1, a)
			}
		}
	}
	heap(len(a), a)

	return result
}
