package main

import (
	"fmt"
	"os"
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

func parseInts(s string) []int64 {
	nums := make([]int64, 0)
	for _, num := range strings.Split(s, " ") {
		nums = append(nums, parseInt(num))
	}
	return nums
}

func day_five_part_one() {
	dat, err := os.ReadFile("input.txt")
	panicerr(err, "reading input file")

	rows := strings.Split(string(dat), "\n")

	seeds := make([]int64, 0)
	soilMap := Map{}
	fertilizerMap := Map{}
	waterMap := Map{}
	lightMap := Map{}
	tempMap := Map{}
	humidityMap := Map{}
	locationMap := Map{}
	var currentMap Map

	RowToMap := map[string]Map{
		"seed-to-soil map:":            soilMap,
		"soil-to-fertilizer map:":      fertilizerMap,
		"fertilizer-to-water map:":     waterMap,
		"water-to-light map:":          lightMap,
		"light-to-temperature map:":    tempMap,
		"temperature-to-humidity map:": humidityMap,
		"humidity-to-location map:":    locationMap,
	}

	for _, row := range rows {
		if len(row) == 0 {
			continue
		}

		if row[0] <= '9' && row[0] >= '0' {
			values := parseInts(row)
			source := values[1]
			dest := values[0]
			currentMap[Range{source, values[2]}] = dest
		} else if strings.HasPrefix(row, "seeds") {
			seeds = parseInts(strings.Split(row, ": ")[1])
		} else {
			if newMap, ok := RowToMap[row]; ok {
				currentMap = newMap
			} else {
				panic(fmt.Sprintf("unknown map: %v", row))
			}

		}
	}

	//fmt.Printf("seeds: %v; maps: %v\n", seeds, RowToMap)

	lowest := int64(0x7fffffffffffffff)
	for _, seed := range seeds {
		v := seed
		for _, m := range []Map{soilMap, fertilizerMap, waterMap, lightMap, tempMap, humidityMap, locationMap} {
			for vRange, dest := range m {
				if vRange.Start <= v && v <= (vRange.Start+vRange.Len) {
                    fmt.Printf("value in range: %v in (%v, %v)\n", v, vRange.Start, vRange.Start+vRange.Len)
                    v = dest + v - vRange.Start
                    break
				}
			}
		}
       
        fmt.Printf("final value: %v\n", v)

		if v < lowest {
			lowest = v
		}
	}

	fmt.Printf("lowest: %v\n", lowest)
}
