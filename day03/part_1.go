package main

import (
	"fmt"
	"strconv"
)

func part1(entries []string) {
	total := 0

	for _, entry := range entries {
		if len(entry) == 0 {
			continue
		}

		maxIndex := 0
		maxFirst := 0
		// find the max battery value
		for index, char := range entry {
			if index == len(entry)-1 {
				break
			}

			value, _ := strconv.Atoi(string(char))

			if maxFirst < value {
				maxIndex = index
				maxFirst = value
			}
		}

		maxSecond := 0
		for i := maxIndex + 1; i < len(entry); i++ {
			char := entry[i]
			value, _ := strconv.Atoi(string(char))
			if maxSecond < value {
				maxSecond = value
			}
		}

		voltage, _ := strconv.Atoi(fmt.Sprintf("%d%d", maxFirst, maxSecond))
		fmt.Println(voltage)

		total += voltage
	}

	fmt.Printf("Total: %v\n", total)
}
