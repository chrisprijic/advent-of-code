package main

import (
	"fmt"
	"os"
	"strings"
)

func isInt(c byte) bool {
	return c >= '0' && c <= '9'
}

func parseIntByte(v byte) int {
	return int(v - '0')
}

func day_one_part_one() {
	dat, err := os.ReadFile("input.txt")
	panicerr(err, "reading input file")

	rows := strings.Split(string(dat), "\n")
	total := 0
	for _, row := range rows {
		if len(row) == 0 {
			continue
		}
		start := 0
		end := len(row) - 1
		for !isInt(row[start]) {
			start++
		}
		for !isInt(row[end]) {
			end--
		}

		value := parseIntByte(row[start])*10 + parseIntByte(row[end])
		total += value
	}

	fmt.Printf("result: %v\n", total)
}
