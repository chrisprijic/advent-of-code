package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day_two_part_two() {
	dat, err := os.ReadFile("input.txt")
	panicerr(err, "reading input file")

	rows := strings.Split(string(dat), "\n")

	total := int64(0)
	for _, row := range rows {
		if len(row) == 0 {
			continue
		}

		values := Values{0, 0, 0}

		parts := strings.Split(row, "; ")
		for _, part := range parts {
			items := strings.Split(part, ", ")
			for _, item := range items {
				match := ColorRegexp.FindStringSubmatch(item)
				if len(match) == 0 {
					continue
				}

				num, _ := strconv.ParseInt(match[1], 10, 32)
				color := match[2]

				switch color {
				case "red":
					if values.Red < num {
						values.Red = num
					}
				case "green":
					if values.Green < num {
						values.Green = num
					}
				case "blue":
					if values.Blue < num {
						values.Blue = num
					}
				}
			}
		}

		total += values.Red * values.Green * values.Blue
	}

	fmt.Printf("result: %v\n", total)
}
