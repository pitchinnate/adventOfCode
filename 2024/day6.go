package main

import (
	"fmt"
	"log"
	"slices"
)

func convertDirection(absolute int) int {
	if absolute == 0 {
		return 0
	} else if absolute == 1 {
		return 1
	} else if absolute == 2 {
		return 0
	} else if absolute == 3 {
		return -1
	}
	return 0
}

func day6_1() {
	lines := ReadFileToLines("./data/day6_data.txt")
	currentX := 0
	currentY := 0
	changeX := 0
	changeXRunning := 0
	changeY := -1
	changeYRunning := 3

	maxX := len(lines[0])
	maxY := len(lines)

	for y, line := range lines {
		for x, letter := range line {
			if string(letter) == "^" {
				currentX = x
				currentY = y
			}
		}
	}

	spacesVisited := []string{}
	spacesVisited = append(spacesVisited, fmt.Sprintf("%d:%d", currentX, currentY))
	for {
		newX := currentX + changeX
		newY := currentY + changeY
		if newX < 0 || newX >= maxX || newY < 0 || newY >= maxY {
			break
		}

		if lines[newY][newX] == '#' {
			changeXRunning = (changeXRunning + 1) % 4
			changeX = convertDirection(changeXRunning)
			changeYRunning = (changeYRunning + 1) % 4
			changeY = convertDirection(changeYRunning)
			newX = currentX + changeX
			newY = currentY + changeY
		}

		currentX = newX
		currentY = newY

		key := fmt.Sprintf("%d:%d", currentX, currentY)
		if !slices.Contains(spacesVisited, key) {
			spacesVisited = append(spacesVisited, key)
		}
	}
	log.Printf("Spaces Visited: %d", len(spacesVisited))
}

func day6_2() {
	lines := ReadFileToLines("./data/day6_test.txt")

	possibleOptions := 0
	for y, line := range lines {
		for x, letter := range line {
			if letter == '.' {
				copyLines := lines
				copyLine := line
				modified := copyLine[:x] + "0" + copyLine[x+1:]
				copyLines[y] = modified
				possibleOptions += checkPossible(copyLines)
			}
		}
	}

	log.Printf("Options: %d", possibleOptions)
}

func checkPossible(lines []string) int {
	log.Print("checking Options:")
	for _, line := range lines {
		fmt.Print(line + "\n")
	}
	fmt.Print("\n")

	currentX := 0
	currentY := 0
	changeX := 0
	changeXRunning := 0
	changeY := -1
	changeYRunning := 3

	maxX := len(lines[0])
	maxY := len(lines)

	for y, line := range lines {
		for x, letter := range line {
			if string(letter) == "^" {
				currentX = x
				currentY = y
			}
		}
	}

	spacesVisited := []string{}
	spacesVisited = append(spacesVisited, fmt.Sprintf("%d:%d:%d:%d", currentX, currentY, changeX, changeY))
	for {
		newX := currentX + changeX
		newY := currentY + changeY
		if newX < 0 || newX >= maxX || newY < 0 || newY >= maxY {
			return 0
		}

		if lines[newY][newX] == '#' || lines[newY][newX] == 'O' {
			changeXRunning = (changeXRunning + 1) % 4
			changeX = convertDirection(changeXRunning)
			changeYRunning = (changeYRunning + 1) % 4
			changeY = convertDirection(changeYRunning)
			newX = currentX + changeX
			newY = currentY + changeY
		}

		currentX = newX
		currentY = newY

		key := fmt.Sprintf("%d:%d", currentX, currentY, changeX, changeY)
		if slices.Contains(spacesVisited, key) {
			return 1
		}
	}
	return 0
}
