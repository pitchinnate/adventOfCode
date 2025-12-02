package main

import (
	"log"
	"strconv"
	"strings"
)

func day2_1() {
	lines := ReadFileToLines("./data/day2_data.txt")
	ranges := strings.Split(lines[0], ",")
	accumulator := 0
	for _, lineRange := range ranges {
		pieces := strings.Split(lineRange, "-")
		start, _ := strconv.Atoi(pieces[0])
		end, _ := strconv.Atoi(pieces[1])
		//log.Printf("day2_1: start=%d, end=%d", start, end)
		for i := start; i <= end; i++ {
			stringVal := strconv.Itoa(i)
			if len(stringVal)%2 == 0 {
				startPiece := stringVal[0 : len(stringVal)/2]
				endPiece := stringVal[len(stringVal)/2:]
				//log.Printf("full %s start: %s end: %s", stringVal, startPiece, endPiece)
				if startPiece == endPiece {
					accumulator += i
				}
			}
		}
	}
	log.Printf("accumulator %d\n", accumulator)
}

func day2_2() {
	lines := ReadFileToLines("./data/day2_data.txt")
	ranges := strings.Split(lines[0], ",")
	accumulator := 0
	for _, lineRange := range ranges {
		pieces := strings.Split(lineRange, "-")
		start, _ := strconv.Atoi(pieces[0])
		end, _ := strconv.Atoi(pieces[1])
		//log.Printf("day2_1: start=%d, end=%d", start, end)
		for i := start; i <= end; i++ {
			stringVal := strconv.Itoa(i)
			alreadyInvalid := false
			for split := 1; split <= len(stringVal)/2; split++ {
				if len(stringVal)%split == 0 && !alreadyInvalid {
					allPieces := []string{}
					//log.Printf("val: %d len %d split %d", i, len(stringVal), split)
					for j := 0; j < len(stringVal)/split; j += 1 {
						startIndex := j * split
						endIndex := startIndex + split
						stringPiece := stringVal[startIndex:endIndex]
						//log.Printf("start: %d end: %d stringPiece: %s", startIndex, endIndex, stringPiece)
						allPieces = append(allPieces, stringPiece)
					}
					//log.Print(allPieces)

					allMatch := true
					for j := 1; j < len(allPieces); j++ {
						if allPieces[j] != allPieces[j-1] {
							allMatch = false
							break
						}
					}
					if allMatch {
						accumulator += i
						alreadyInvalid = true
					}
				}
			}
		}
	}
	log.Printf("accumulator %d\n", accumulator)
}
