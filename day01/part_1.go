package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	content, err := os.ReadFile("./day01/input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	minValue := 0
	maxValue := 100
	//
	current := 50
	password := 0

	lines := strings.Split(string(content), "\n")

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		direction := string(line[0])
		value, _ := strconv.Atoi(line[1:])
		//prev := current
		steps := value % maxValue

		if direction == "L" {
			current = current - steps
		}

		if direction == "R" {
			current = current + steps
		}

		if current >= maxValue {
			current = current % maxValue
		}

		if current < minValue {
			current = maxValue + current
		}

		if current == 0 {
			password++
		}

		//fmt.Println(prev, line, current)
	}

	fmt.Printf("Final value: %v, Password: %v", current, password)
}
