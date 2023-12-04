package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func day_four_part_two() {
	dat, err := os.ReadFile("input.txt")
	panicerr(err, "reading input file")

	rows := strings.Split(string(dat), "\n")
	cards := []Card{}

    queue := []int{}
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

        queue = append(queue, len(cards) - 1)
	}

    total := 0

    for len(queue) > 0 {
        index := queue[0]
        queue = queue[1:]

        score := 0
        for _, winner := range cards[index].Winners {
            for _, value := range cards[index].Values {
                if winner == value {
                    score += 1
                }
            }
        }

        for i := 0; i < score; i++ {
            queue = append(queue, index + 1 + i)
        }
        total += 1
    }
	
	fmt.Printf("Total: %v\n", total)
}
