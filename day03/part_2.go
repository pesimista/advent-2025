package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Digit struct {
	value    int
	position int
}

func part2(entries []string) {
	total := 0

	length := 12
	digits := make([]Digit, length)
	for _, entry := range entries {
		if len(entry) == 0 {
			continue
		}

		str := strings.Builder{}

		for i := 0; i < length; i++ {
			digits[i] = Digit{0, 0}

			endIndex := len(entry) - length + i + 1
			startIndex := 0
			if i > 0 {
				startIndex = digits[i-1].position + 1
			}

			for j := startIndex; j < endIndex; j++ {
				char := entry[j]
				value, _ := strconv.Atoi(string(char))

				if digits[i].value < value {
					digits[i].value = value
					digits[i].position = j
				}
			}

			str.WriteString(strconv.Itoa(digits[i].value))
		}

		voltage, _ := strconv.Atoi(str.String())
		// fmt.Println(voltage)

		total += voltage
	}

	fmt.Printf("Total: %v\n", total)
}
