package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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
	total := 0
	for _, value := range regexp.MustCompile(`-?\d+`).FindAllString(input, -1) {
		i, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		total += i
	}

	return total
}

func part2(input string) int {
	var m interface{}
	err := json.Unmarshal([]byte(input), &m)
	if err != nil {
		log.Fatal(err)
	}

	return sumNumbers(m)
}

func sumNumbers(m interface{}) int {
	total := 0
	switch item := m.(type) {
	case int:
		total += int(item)
	case float64:
		total += int(item)
	case []interface{}:
		for _, v := range item {
			total += sumNumbers(v)
		}
	case map[string]interface{}:
		for _, v := range item {
			if str, ok := v.(string); ok && str == "red" {
				return 0
			}
			total += sumNumbers(v)
		}
	}

	return total
}
