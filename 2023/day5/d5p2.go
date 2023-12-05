package main

import (
	"fmt"
	"os"
	"strings"
)

func min(a, b int64) int64 {
    if a < b {
        return a
    }

    return b
}

func max(a, b int64) int64 {
    if a > b {
        return a
    }

    return b
}

func overlap(r1, r2 Range) *Range {
    if r1.Start > r2.Start {
        r1, r2 = r2, r1
    }

    if r2.Start > r1.Start + r1.Len {
        return nil
    }

    start := max(r1.Start, r2.Start)
    end := min(r1.Start + r1.Len, r2.Start + r2.Len)

    return &Range{start, end - start}
}

func nonOverlap(r1, r2 Range) []Range {
    overlap := overlap(r1, r2)
    if overlap == nil {
        return []Range{r1}
    }

    ranges := []Range{}

    if r1.Start < overlap.Start {
        ranges = append(ranges, Range{r1.Start, overlap.Start - r1.Start})
    }

    if r1.Start + r1.Len > overlap.Start + overlap.Len {
        ranges = append(ranges, Range{overlap.Start + overlap.Len, r1.Start + r1.Len - overlap.Start - overlap.Len})
    }

    return ranges
}

func day_five_part_two() {
	dat, err := os.ReadFile("input.txt")
	panicerr(err, "reading input file")

	rows := strings.Split(string(dat), "\n")

	seeds := make([]Range, 0)
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
			values := parseInts(strings.Split(row, ": ")[1])
			for i := 0; i < len(values); i += 2 {
				seeds = append(seeds, Range{values[i], values[i+1]})
			}
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
	for _, seedR := range seeds {
	    rangesToCheck := []Range{seedR}
        finalRanges := []Range{}
        for _, m := range []Map{soilMap, fertilizerMap, waterMap, lightMap, tempMap, humidityMap, locationMap} {
            for vRange, dest := range m {
                newRanges := []Range{}
    
                for _, r := range rangesToCheck {
                    if ol := overlap(r, vRange); ol != nil {
                        finalRanges = append(finalRanges, Range{dest+(ol.Start-vRange.Start), ol.Len})
                        newRanges = append(newRanges, nonOverlap(r, vRange)...)
                    } else {
                        newRanges = append(newRanges, r)
                    }
                }

                rangesToCheck = newRanges
            }

            if len(rangesToCheck) != 0 {
                finalRanges = append(finalRanges, rangesToCheck...)
            }
            rangesToCheck = finalRanges
            finalRanges = []Range{}
        }

        for _, r := range rangesToCheck {
            if r.Start < lowest {
                lowest = r.Start
            }
        }
    }

	fmt.Printf("lowest: %v\n", lowest)
}
