package main

import (
	"fmt"
	"os"
	"strings"
)

type Vert struct {
	X int64
	Y int64
}

type BB struct {
	Min Vert
	Max Vert
}

type Number struct {
	RawValue string
	Value    int64
	BB
}

type Symbol struct {
	Value         string
	AdjacentCount int64
	Vert
}

func (n *Number) Finalize() {
	n.Value = 0
	for i := 0; i < len(n.RawValue); i++ {
		n.Value *= 10
		n.Value += int64(n.RawValue[i] - '0')
	}
}

func (bb BB) Expand() BB {
	return BB{Vert{bb.Min.X - 1, bb.Min.Y - 1}, Vert{bb.Max.X + 1, bb.Max.Y + 1}}
}

func (bb BB) Contains(v Vert) bool {
	bb = bb.Expand()
	return v.X >= bb.Min.X && v.X <= bb.Max.X && v.Y >= bb.Min.Y && v.Y <= bb.Max.Y
}

func day_three_part_one() {
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

	for _, num := range numbers {
		for _, symbol := range symbols {
			if num.BB.Contains(symbol.Vert) {
				//fmt.Printf("Number %v contains symbol %v\n", num, symbol)
				total += num.Value
				break
			}
		}
	}

	fmt.Printf("Total: %v\n", total)
}
