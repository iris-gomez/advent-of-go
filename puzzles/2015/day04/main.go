package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
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
	number := 0
	for {
		h := md5.New()
		io.WriteString(h, input+strconv.Itoa(number))
		sum := hex.EncodeToString(h.Sum(nil)[:])
		if strings.HasPrefix(sum, "00000") {
			break
		}
		number++
	}

	return number
}

func part2(input string) int {
	number := 0
	for {
		h := md5.New()
		io.WriteString(h, input+strconv.Itoa(number))
		sum := hex.EncodeToString(h.Sum(nil)[:])
		if strings.HasPrefix(sum, "000000") {
			break
		}
		number++
	}

	return number
}
