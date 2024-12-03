package main

import (
	"log"
	"math"
	"slices"
	"strconv"
	"strings"
)

func day1_1() {
	lines := ReadFileToLines("./data/day1_data.txt")
	col1, col2 := processLines(lines)
	slices.Sort(col1)
	slices.Sort(col2)

	diff := 0

	for i, val := range col1 {
		diff += int(math.Abs(float64(val - col2[i])))
	}
	log.Printf("total diff: %d", diff)
}

func day1_2() {
	lines := ReadFileToLines("./data/day1_data.txt")
	col1, col2 := processLines(lines)
	slices.Sort(col1)
	slices.Sort(col2)

	diff := 0

	for _, val := range col1 {
		count := 0
		for _, val2 := range col2 {
			if val2 == val {
				count++
			}
		}

		diff += (count * val)
	}
	log.Printf("total diff: %d", diff)
}

func processLines(lines []string) (col1 []int, col2 []int) {
	for _, line := range lines {
		pieces := strings.Split(line, " ")
		foundCol1 := false
		for _, piece := range pieces {
			if piece != "" {
				val, _ := strconv.Atoi(piece)
				if !foundCol1 {
					col1 = append(col1, val)
					foundCol1 = true
				} else {
					col2 = append(col2, val)
				}
			}
		}
	}
	return
}
