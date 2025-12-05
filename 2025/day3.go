package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func day3_1() {
	lines := ReadFileToLines("./data/day3_data.txt")
	totalMax := 0
	for _, line := range lines {
		numberStrings := strings.Split(line, "")
		numbers := make([]int, len(numberStrings))
		for i, s := range numberStrings {
			numbers[i], _ = strconv.Atoi(s)
		}

		maxNum := 0
		for x := 0; x < len(numbers)-1; x++ {
			for y := x + 1; y < len(numbers); y++ {
				combined := fmt.Sprintf("%d%d", numbers[x], numbers[y])
				combinedVal, _ := strconv.Atoi(combined)
				if combinedVal > maxNum {
					maxNum = combinedVal
				}
			}
		}
		log.Println(maxNum)
		totalMax += maxNum
	}
	fmt.Println(totalMax)
}

func day3_2() {
	lines := ReadFileToLines("./data/day3_data.txt")
	totalMax := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		numberStrings := strings.Split(line, "")
		numbers := make([]int, len(numberStrings))
		for i, s := range numberStrings {
			numbers[i], _ = strconv.Atoi(s)
		}
		//log.Println(numbers)

		//maxNum := 0
		digitsNeeded := 12
		currentIndexes := []int{}
		for x := len(numbers) - digitsNeeded; x <= len(numbers)-1; x++ {
			currentIndexes = append(currentIndexes, x)
		}
		currentMax := getNumberFromIndexes(currentIndexes, numbers)
		//log.Println(currentIndexes)
		//log.Println(currentMax)

		for x := 0; x < digitsNeeded; x++ {
			startIndex := currentIndexes[x]
			minIndex := 0
			if x > 0 {
				minIndex = currentIndexes[x-1] + 1
			}
			if startIndex > minIndex {
				for y := startIndex - 1; y >= minIndex; y-- {
					var copyIndexes = make([]int, len(currentIndexes))
					copy(copyIndexes, currentIndexes)
					copyIndexes[x] = y
					newMax := getNumberFromIndexes(copyIndexes, numbers)
					//log.Printf("new number: %d", newMax)
					if newMax >= currentMax {
						//log.Printf("oldMax: %d, newMax: %d, indexes: %v, newIndexes: %v", currentMax, newMax, currentIndexes, copyIndexes)
						currentMax = newMax
						currentIndexes = copyIndexes
					}
				}
			}
		}
		//log.Println("Line Max: ", currentMax)
		totalMax += currentMax
		//log.Println("------------------------------------")
	}
	fmt.Println(totalMax)
}

func getNumberFromIndexes(indexes []int, numbers []int) int {
	stringVal := ""
	for x := 0; x < len(indexes); x++ {
		stringVal += strconv.Itoa(numbers[indexes[x]])
	}
	newVal, _ := strconv.Atoi(stringVal)
	return newVal
}

func combineStringDigits(vals []string) int {
	stringVal := ""
	for i := 0; i < len(vals); i++ {
		stringVal += vals[i]
	}
	newVal, _ := strconv.Atoi(stringVal)
	return newVal
}

func combineDigits(vals []int) int {
	stringVal := ""
	for i := 0; i < len(vals); i++ {
		stringVal += strconv.Itoa(vals[i])
	}
	newVal, _ := strconv.Atoi(stringVal)
	return newVal
}
