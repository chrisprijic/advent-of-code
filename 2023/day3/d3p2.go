package main

import (
	"fmt"
	"os"
	"strings"
)

func day_three_part_two() {
	dat, err := os.ReadFile("input.txt")
	panicerr(err, "reading input file")

	rows := strings.Split(string(dat), "\n")

	numbers := make([]Number, 0)
	symbols := make([]Symbol, 0)

	for y := 0; y < len(rows); y++ {
		var num *Number
		for x := 0; x < len(rows[y]); x++ {
			switch rows[y][x] {
			case '.':
				if num != nil {
					num.Finalize()
					numbers = append(numbers, *num)
					num = nil
				}
				continue
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				if num == nil {
					num = &Number{RawValue: string(rows[y][x]), BB: BB{Vert{int64(x), int64(y)}, Vert{int64(x), int64(y)}}}
				} else {
					num.BB.Max.X = int64(x)
					num.RawValue += string(rows[y][x])
				}
			default:
				if num != nil {
					num.Finalize()
					numbers = append(numbers, *num)
					num = nil
				}
				symbols = append(symbols, Symbol{string(rows[y][x]), 0, Vert{int64(x), int64(y)}})
			}
		}
		if num != nil {
			num.Finalize()
			numbers = append(numbers, *num)
			num = nil
		}
	}

	total := int64(0)

	for _, symbol := range symbols {
		adjacents := []Number{}
		for _, num := range numbers {
			if num.BB.Contains(symbol.Vert) {
				adjacents = append(adjacents, num)
			}
		}

		if len(adjacents) == 2 {
			//fmt.Printf("Symbol %v has two adjacent numbers: %v\n", symbol, adjacents)
			total += adjacents[0].Value * adjacents[1].Value
		}
	}

	fmt.Printf("Total: %v\n", total)
}
