package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(strings.ReplaceAll(strings.TrimSpace(string(input)), ".", ""), "\n")

	fmt.Printf("Part 1 solution: %d\n", part1(lines))
	fmt.Printf("Part 2 solution: %d\n", part2(lines))
}

func part1(lines []string) int {
	return maxHappiness(getAtendeesAndRelationships(lines))
}

func part2(lines []string) int {
	atendees, relationships := getAtendeesAndRelationships(lines)

	relationships["me"] = map[string]int{}
	for _, atendee := range atendees {
		relationships["me"][atendee] = 0
	}
	atendees = append(atendees, "me")

	return maxHappiness(atendees, relationships)
}

func getAtendeesAndRelationships(lines []string) ([]string, map[string]map[string]int) {
	relationships := map[string]map[string]int{}
	atendees := []string{}
	for _, line := range lines {
		split := strings.Split(line, " ")

		if _, ok := relationships[split[0]]; !ok {
			relationships[split[0]] = map[string]int{}
			atendees = append(atendees, split[0])
		}

		happiness, err := strconv.Atoi(split[3])
		if err != nil {
			log.Fatal(err)
		}

		switch split[2] {
		case "gain":
			relationships[split[0]][split[10]] = happiness
		case "lose":
			relationships[split[0]][split[10]] = -happiness
		}
	}

	return atendees, relationships
}

func maxHappiness(atendees []string, relationships map[string]map[string]int) int {
	happiness := 0
	for _, permutation := range permutations(atendees) {
		currentHappiness := 0

		for i, atendee := range permutation {
			switch i {
			case 0:
				currentHappiness += relationships[atendee][permutation[len(permutation)-1]]
				currentHappiness += relationships[atendee][permutation[i+1]]
			case len(permutation) - 1:
				currentHappiness += relationships[atendee][permutation[i-1]]
				currentHappiness += relationships[atendee][permutation[0]]
			default:
				currentHappiness += relationships[atendee][permutation[i-1]]
				currentHappiness += relationships[atendee][permutation[i+1]]
			}
		}

		happiness = max(happiness, currentHappiness)
	}

	return happiness
}

func permutations(a []string) [][]string {
	var heap func(int, []string)
	result := [][]string{}

	heap = func(k int, a []string) {
		if k == 1 {
			temp := make([]string, len(a))
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
