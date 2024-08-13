package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Ingredient struct {
	name       string
	capacity   int
	durability int
	flavor     int
	texture    int
	calories   int
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

func part1(lines []string) int {
	ingredients := parseLines(lines)
	maxScore := 0
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100-i; j++ {
			for k := 0; k <= 100-i-j; k++ {
				l := 100 - i - j - k
				capacity := max(0, ingredients[0].capacity*i+ingredients[1].capacity*j+ingredients[2].capacity*k+ingredients[3].capacity*l)
				durability := max(0, ingredients[0].durability*i+ingredients[1].durability*j+ingredients[2].durability*k+ingredients[3].durability*l)
				flavor := max(0, ingredients[0].flavor*i+ingredients[1].flavor*j+ingredients[2].flavor*k+ingredients[3].flavor*l)
				texture := max(0, ingredients[0].texture*i+ingredients[1].texture*j+ingredients[2].texture*k+ingredients[3].texture*l)
				maxScore = max(maxScore, capacity*durability*flavor*texture)
			}
		}
	}
	return maxScore
}

func part2(lines []string) int {
	ingredients := parseLines(lines)
	maxScore := 0
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100-i; j++ {
			for k := 0; k <= 100-i-j; k++ {
				l := 100 - i - j - k
				capacity := max(0, ingredients[0].capacity*i+ingredients[1].capacity*j+ingredients[2].capacity*k+ingredients[3].capacity*l)
				durability := max(0, ingredients[0].durability*i+ingredients[1].durability*j+ingredients[2].durability*k+ingredients[3].durability*l)
				flavor := max(0, ingredients[0].flavor*i+ingredients[1].flavor*j+ingredients[2].flavor*k+ingredients[3].flavor*l)
				texture := max(0, ingredients[0].texture*i+ingredients[1].texture*j+ingredients[2].texture*k+ingredients[3].texture*l)
				calories := ingredients[0].calories*i + ingredients[1].calories*j + ingredients[2].calories*k + ingredients[3].calories*l
				if calories == 500 {
					maxScore = max(maxScore, capacity*durability*flavor*texture)
				}
			}
		}
	}
	return maxScore
}

func parseLines(lines []string) []Ingredient {
	ingredients := []Ingredient{}
	for _, line := range lines {
		split := strings.Split(strings.ReplaceAll(strings.ReplaceAll(line, ",", ""), ":", ""), " ")

		name := split[0]

		capacity, err := strconv.Atoi(split[2])
		if err != nil {
			log.Fatal(err)
		}

		durability, err := strconv.Atoi(split[4])
		if err != nil {
			log.Fatal(err)
		}

		flavor, err := strconv.Atoi(split[6])
		if err != nil {
			log.Fatal(err)
		}

		texture, err := strconv.Atoi(split[8])
		if err != nil {
			log.Fatal(err)
		}

		calories, err := strconv.Atoi(split[10])
		if err != nil {
			log.Fatal(err)
		}

		ingredients = append(ingredients, Ingredient{name, capacity, durability, flavor, texture, calories})
	}

	return ingredients
}
