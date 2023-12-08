package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func day_six_part_two() {
	dat, err := os.ReadFile("input.txt")
	panicerr(err, "reading input file")

	rows := strings.Split(string(dat), "\n")

	times := []int64{}
	distances := []int64{}

	for _, row := range rows {
		if len(row) == 0 {
			continue
		}
		re := regexp.MustCompile("\\s")
		row = re.ReplaceAllString(row, "")
		if strings.HasPrefix(row, "Time:") {
			times = parseInts(strings.TrimLeft(strings.Split(row, ":")[1], " "))
		} else if strings.HasPrefix(row, "Distance:") {
			distances = parseInts(strings.TrimLeft(strings.Split(row, ":")[1], " "))
		} else {
			panic("Unknown line: " + row)
		}
	}

	//fmt.Printf("Times: %v\n", times)
	//fmt.Printf("Distances: %v\n", distances)
	total := 1
	for i := 0; i < len(times); i++ {
		time := times[i]
		distance := distances[i]
		totalWins := 0

		for t := int64(0); t < time; t++ {
			buttonTime := t
			goTime := time - t
			d := buttonTime * goTime

			fmt.Printf("Time: %v, Button: %v, Go: %v, d: %v, distance: %v, totalWins: %v\n", t, buttonTime, goTime, d, distance, totalWins)

			if d > distance {
				totalWins++
			}
		}
		fmt.Printf("Total wins: %v\n", totalWins)
		total *= totalWins
	}

	fmt.Printf("Total: %v\n", total)
}
