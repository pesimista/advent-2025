package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part2() {
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
		prev := current
		steps := value % maxValue
		extraSteps := value / maxValue
		password += extraSteps
		//change := 0

		if direction == "L" {
			current = current - steps
		}

		if direction == "R" {
			current = current + steps
		}

		if current >= maxValue {
			current = current % maxValue
			if current != 0 {
				// change++
				password++
			}
		}

		if current < minValue {
			current = maxValue + current
			if prev != 0 {
				// change++
				password++
			}
		}

		if current == 0 {
			// change++
			password++
		}

		//fmt.Printf("prev: %v \tline: %v \tcurrent: %v \textraSteps: %v \tchange: %v \tpassword: %v\n", prev, line, current, extraSteps, change, password)
	}

	fmt.Printf("Final value: %v, Password: %v", current, password)
}
