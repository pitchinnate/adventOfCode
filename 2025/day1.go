package main

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func day1_1() {
	lines := ReadFileToLines("./data/day1_data.txt")
	count := 0
	currentPoint := 50
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) < 2 {
			continue
		}
		remaining := trimmedLine[1:]
		changeAmount, _ := strconv.Atoi(remaining)

		changeAmount = changeAmount % 100
		if changeAmount < 100 {
			if strings.HasPrefix(trimmedLine, "R") {
				currentPoint += changeAmount
			} else {
				currentPoint -= changeAmount
			}
		}
		if currentPoint > 99 {
			currentPoint -= 100
		} else if currentPoint < -0 {
			currentPoint += 100
		}
		log.Printf("%s -> %d -> %d", trimmedLine, changeAmount, currentPoint)
		if currentPoint == 0 {
			count += 1
		}
	}
	log.Printf("Day 1 - Total Points: %d", count)
}

func day1_2() {
	lines := ReadFileToLines("./data/day1_data.txt")
	count := 0
	currentPoint := 50
	previousPoint := 50
	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if len(trimmedLine) < 2 {
			continue
		}
		remaining := trimmedLine[1:]
		changeAmount, _ := strconv.Atoi(remaining)

		fullRotations := math.Floor(float64(changeAmount) / 100)
		count += int(fullRotations)

		changeAmount = changeAmount % 100
		if changeAmount < 100 {
			if strings.HasPrefix(trimmedLine, "R") {
				currentPoint += changeAmount
			} else {
				currentPoint -= changeAmount
			}
		}
		if currentPoint > 99 {
			currentPoint -= 100
			if previousPoint != 0 {
				count += 1
			}
		} else if currentPoint < 0 {
			currentPoint += 100
			if previousPoint != 0 {
				count += 1
			}
		} else if currentPoint == 0 {
			count += 1
		}
		previousPoint = currentPoint
		log.Printf("%s -> %d -> %d = %d", trimmedLine, changeAmount, currentPoint, count)
	}
	log.Printf("Day 1 - Total Points: %d", count)
}
