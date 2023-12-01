package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	lookups  = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	mappings = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"0":     0,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}
)

func day_one_part_two() {
	dat, err := os.ReadFile("input.txt")
	panicerr(err, "reading input file")

	rows := strings.Split(string(dat), "\n")
	total := 0
	for _, row := range rows {
		if row == "" {
			continue
		}

		first := len(row)
		firstValue := -1
		end := -1
		lastValue := -1
		for _, lookup := range lookups {
			if strings.Contains(row, lookup) {
				fid := strings.Index(row, lookup)
				lid := strings.LastIndex(row, lookup)
				if fid < first {
					first = fid
					firstValue = mappings[lookup]
				}
				if end < lid {
					end = lid
					lastValue = mappings[lookup]
				}
			}
		}

		value := firstValue*10 + lastValue
		total += value

		//fmt.Printf("row: %v, first: %v (firstValue: %v), last: %v (lastValue: %v), value: %v, total: %v\n", row, first, firstValue, end, lastValue, value, total)
	}

	fmt.Printf("result: %v\n", total)
}
