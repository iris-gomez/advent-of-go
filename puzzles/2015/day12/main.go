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
	input, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	document := strings.TrimSpace(string(input))

	fmt.Printf("Part 1 solution: %v\n", part1(document))
	fmt.Printf("Part 2 solution: %v\n", part2(document))
}

func part1(document string) int {
	total := 0
	for _, value := range regexp.MustCompile(`-?\d+`).FindAllString(document, -1) {
		i, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		total += i
	}

	return total
}

func part2(document string) int {
	var m interface{}
	err := json.Unmarshal([]byte(document), &m)
	if err != nil {
		log.Fatal(err)
	}

	return walkJSON(m)
}

func walkJSON(m interface{}) int {
	total := 0
	switch item := m.(type) {
	case int:
		total += int(item)
	case float64:
		total += int(item)
	case []interface{}:
		for _, v := range item {
			total += walkJSON(v)
		}
	case map[string]interface{}:
		for _, v := range item {
			if str, ok := v.(string); ok && str == "red" {
				return 0
			}
			total += walkJSON(v)
		}
	}

	return total
}
