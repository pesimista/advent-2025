package main

import (
	"fmt"
	"os"
	"strings"
)

// https://adventofcode.com/2025/day/1 `

func main() {
	fileContent, err := os.ReadFile("./day02/input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	stringContent := strings.ReplaceAll(string(fileContent), "\n", "")
	lines := strings.Split(stringContent, ",")

	part1(lines)
	part2(lines)
}
