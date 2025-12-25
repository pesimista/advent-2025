package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func part2(entries []string) {
	total := 0

	for _, entry := range entries {
		if len(entry) == 0 {
			continue
		}

		ids := strings.Split(entry, "-")
		startValue, _ := strconv.Atoi(ids[0])
		endValue, _ := strconv.Atoi(ids[1])

		//fmt.Println(entry)

		for i := startValue; i <= endValue; i++ {
			stringValue := strconv.Itoa(i)
			length := len(stringValue)
			halfLength := length / 2

			firstHalf := stringValue[:halfLength]
			secondHalf := stringValue[halfLength:]

			if length%2 == 0 && firstHalf == secondHalf {
				//fmt.Println("  invalid", i, firstHalf)
				total += i
				continue
			}

			for j := 0; j < halfLength; j++ {
				index := j + 1
				exp := regexp.MustCompile(stringValue[:index])

				matches := exp.FindAll([]byte(stringValue), -1)
				if len(matches) == length/index && length%index == 0 {
					//fmt.Println("  invalid", i, stringValue[:index])
					total += i
					break
				}
			}
		}
	}

	fmt.Printf("Total: %v\n", total)
}
