package main

import (
	"fmt"
	"os"
	"strings"
)

// https://adventofcode.com/2025/day/3

func main() {
	fileContent, err := os.ReadFile("./day03/input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(fileContent), "\n")

	part1(lines)
	part2(lines)
}
