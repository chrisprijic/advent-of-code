package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	jstrengths = map[rune]int{
		'A': 13,
		'K': 12,
		'Q': 11,
		'T': 10,
		'9': 9,
		'8': 8,
		'7': 7,
		'6': 6,
		'5': 5,
		'4': 4,
		'3': 3,
		'2': 2,
		'J': 1,
	}

	jstrengthList = []rune{
		'A',
		'K',
		'Q',
		'T',
		'9',
		'8',
		'7',
		'6',
		'5',
		'4',
		'3',
		'2',
	    'J',
    }
)

type JCard struct {
	Value       string
	SortedValue string
	Bid         int
	rank        int
	strength    int
}

func sortJHand(v string) string {
	ra := []rune(v)
	sort.Slice(ra, func(i, j int) bool {
		return jstrengths[ra[i]] > jstrengths[ra[j]]
	})

	return string(ra)
}

func calcJHand(s string) int {
	v := []rune(s)
	// five of a kind
	if v[0] == v[1] && v[1] == v[2] && v[2] == v[3] && v[3] == v[4] {
		return 7
	}

	// four of a kind
	if v[0] == v[1] && v[1] == v[2] && v[2] == v[3] {
		return 6
	}

	if v[1] == v[2] && v[2] == v[3] && v[3] == v[4] {
		return 6
	}

	// full house
	if v[0] == v[1] && v[1] == v[2] && v[3] == v[4] {
		return 5
	}

	if v[0] == v[1] && v[2] == v[3] && v[3] == v[4] {
		return 5
	}

	// three of a kind
	if v[0] == v[1] && v[1] == v[2] {
		return 4
	}

	if v[1] == v[2] && v[2] == v[3] {
		return 4
	}

	if v[2] == v[3] && v[3] == v[4] {
		return 4
	}

	// two pair
	if v[0] == v[1] && v[2] == v[3] {
		return 3
	}

	if v[0] == v[1] && v[3] == v[4] {
		return 3
	}

	if v[1] == v[2] && v[3] == v[4] {
		return 3
	}

	// one pair
	if v[0] == v[1] || v[1] == v[2] || v[2] == v[3] || v[3] == v[4] {
		return 2
	}

	return 1
}

func calcJ(s string, rank int) int {
    v := []rune(s)
    nj := 0
    for i := len(v) - 1; i >= 0; i-- {
        if v[i] == 'J' {
            nj++
        } else {
            break
        }
    }

    fmt.Printf("'%v' %v\n", s, nj)

    if nj == 0 {
        return rank
    }

    switch rank {
    case 7:
        return 7
    case 6:
        return 7
    case 5:
        return 7
    case 4:
        return 6
    case 3:
        if nj == 1 {
            return 5
        }
        return 6
    case 2:
        return 4
    case 1:
        return 2
    }

    return rank
}

func calcJStrength(s string) int {
	v := []rune(s)
	str := float64(0)
	for i, r := range v {
		str += float64(jstrengths[r]) * math.Pow(100, float64(len(v)-i))
	}

	return int(str)
}

func NewJCard(value string, bid int) JCard {
	c := JCard{
		Value:       value,
		SortedValue: sortJHand(value),
		Bid:         bid,
		rank:        0,
		strength:    0,
	}

	c.rank = calcJ(c.SortedValue, calcJHand(c.SortedValue))
	c.strength = calcJStrength(c.Value)

	return c
}

func day_seven_part_two() {
	dat, err := os.ReadFile("input.txt")
	panicerr(err, "reading input file")

	rows := strings.Split(string(dat), "\n")
	cards := []JCard{}
	for _, row := range rows {
		if len(row) == 0 {
			continue
		}

		parts := strings.Split(row, " ")
		value := parts[0]
		bid, _ := strconv.Atoi(parts[1])
		cards = append(cards, NewJCard(value, bid))
	}

	sort.Slice(cards, func(i, j int) bool {
		return cards[i].rank < cards[j].rank || (cards[i].rank == cards[j].rank && cards[i].strength < cards[j].strength)
	})

	total := 0
	for i, c := range cards {
		fmt.Printf("%v %v (%v) (%v)\n", c.Value, c.SortedValue, c.rank, c.strength)
		total += c.Bid * (i + 1)
	}

	fmt.Printf("total: %d\n", total)
}
