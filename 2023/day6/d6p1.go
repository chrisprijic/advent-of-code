package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Range struct {
	Start int64
	Len   int64
}

type Map map[Range]int64

func parseInt(s string) int64 {
	num, err := strconv.ParseInt(s, 10, 64)
	panicerr(err, "parsing int")
	return num
}

func calc(T int64, D int64) (int64, int64) {
	ft := float64(T)
	fd := float64(D)
	inner := math.Sqrt(ft*ft - 4*fd)
	min := (-ft - inner) / 2.0
	max := (-ft + inner) / 2.0

	return int64(math.Ceil(min)), int64(math.Floor(max))
}

func parseInts(s string) []int64 {
	nums := make([]int64, 0)
	re := regexp.MustCompile("\\s+")
	for _, num := range re.Split(s, -1) {
		nums = append(nums, parseInt(num))
	}
	return nums
}

func day_six_part_one() {
	dat, err := os.ReadFile("input.txt")
	panicerr(err, "reading input file")

	rows := strings.Split(string(dat), "\n")

	times := []int64{}
	distances := []int64{}

	for _, row := range rows {
		if len(row) == 0 {
			continue
		}

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
