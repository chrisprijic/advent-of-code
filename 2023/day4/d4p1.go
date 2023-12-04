package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Card struct {
	Winners []string
	Values  []string
}

func string_to_int(v string) int {
	i, _ := strconv.ParseInt(v, 10, 32)
	return int(i)
}

func day_four_part_one() {
	dat, err := os.ReadFile("input.txt")
	panicerr(err, "reading input file")

	rows := strings.Split(string(dat), "\n")
	cards := []Card{}

	for _, row := range rows {
		if len(row) == 0 {
			continue
		}
		info := strings.Split(row, ": ")[1]
        info = regexp.MustCompile("\\s+").ReplaceAllString(info, " ")
		parts := strings.Split(info, " | ")

		cards = append(cards, Card{
			Winners: strings.Split(parts[0], " "),
			Values:  strings.Split(parts[1], " "),
		})
	}

	total := 0

	for _, card := range cards {
		score := 0
		for _, winner := range card.Winners {
			for _, value := range card.Values {
				if winner == value {
					if score == 0 {
						score = 1
					} else {
						score *= 2
					}
				}
			}
		}

		total += score
	}

	fmt.Printf("Total: %v\n", total)
}
