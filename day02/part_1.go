package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part1(entries []string) {
	total := 0

	for _, entry := range entries {
		if len(entry) == 0 {
			continue
		}

		ids := strings.Split(entry, "-")
		startValue, _ := strconv.Atoi(ids[0])
		endValue, _ := strconv.Atoi(ids[1])

		for i := startValue; i <= endValue; i++ {
			stringValue := strconv.Itoa(i)
			length := len(stringValue)

			// valid can't be sequence
			if length%2 != 0 {
				continue
			}

			firstHalf := stringValue[:length/2]
			secondHalf := stringValue[length/2:]

			// invalid
			if firstHalf == secondHalf {
				total += i
			}
		}
	}

	fmt.Printf("Total: %v\n", total)
}
