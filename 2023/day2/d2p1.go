package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Values struct {
	Red   int64
	Green int64
	Blue  int64
}

var (
	GameRegexp  = regexp.MustCompile("Game\\s(\\d+):\\s")
	ColorRegexp = regexp.MustCompile("(\\d+)\\s(\\w+)")
)

func day_two_part_one() {
	dat, err := os.ReadFile("input.txt")
	panicerr(err, "reading input file")

    values := Values{12, 13, 14}

	rows := strings.Split(string(dat), "\n")

	total := int64(0)
	for _, row := range rows {
		if len(row) == 0 {
			continue
		}

		game, _ := strconv.ParseInt(GameRegexp.FindStringSubmatch(row)[1], 10, 32)
		winner := true
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
						winner = false
					}
				case "green":
					if values.Green < num {
						winner = false
					}
				case "blue":
					if values.Blue < num {
						winner = false
					}
				}
			}
		}

		if winner {
            total += game
		}
	}

	fmt.Printf("result: %v\n", total)
}
