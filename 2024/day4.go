package main

import "log"

func day4_1() {
	searchFor := "XMAS"
	lines := ReadFileToLines("./data/day4_data.txt")
	totalFound := 0
	for y, line := range lines {
		for x, letter := range line {
			if letter == 'X' {
				found := 0
				for changeX := -1; changeX <= 1; changeX++ {
					for changeY := -1; changeY <= 1; changeY++ {
						found += checkForString(lines, searchFor, x, y, changeX, changeY)
					}
				}
				totalFound += found
			}
		}
	}
	log.Printf("Total found: %d", totalFound)
}
func day4_2() {
	searchFor := "MAS"
	lines := ReadFileToLines("./data/day4_data.txt")
	totalFound := 0
	for y, line := range lines {
		for x, letter := range line {
			if letter == 'A' {
				found := checkCorners(lines, searchFor, x, y)
				if found {
					totalFound += 1
				}
			}
		}
	}
	log.Printf("Total found: %d", totalFound)
}

func checkCorners(lines []string, search string, startX int, startY int) bool {
	maxX := len(lines[0]) - 1
	maxY := len(lines) - 1

	numFound := 0

	for y := -1; y <= 1; y += 2 {
		for x := -1; x <= 1; x += 2 {
			newX := startX + x
			newY := startY + y
			oppositeX := startX - x
			oppositeY := startY - y
			if newX < 0 || newY < 0 || newX > maxX || newY > maxY {
				continue
			}
			if oppositeX < 0 || oppositeY < 0 || oppositeX > maxX || oppositeY > maxY {
				continue
			}
			buildString := string(lines[newY][newX]) + "A" + string(lines[oppositeY][oppositeX])
			if buildString == search {
				numFound += 1
			}
		}
	}

	return numFound == 2
}

func checkForString(lines []string, search string, startX int, startY int, changeX int, changeY int) int {
	maxX := len(lines[0]) - 1
	maxY := len(lines) - 1

	buildString := ""
	for i := 0; i < len(search); i++ {
		newX := startX + (changeX * i)
		newY := startY + (changeY * i)
		if newX < 0 || newY < 0 || newX > maxX || newY > maxY {
			return 0
		}
		buildString += string(lines[newY][newX])
	}
	if buildString == search {
		return 1
	}

	return 0
}
