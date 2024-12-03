package main

import (
	"fmt"
	"log"
	"strconv"
)

var digits = map[string]int32{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func day1_1() {
	lines := ReadFileToLines("./data/day1_test.txt")
	readLines(lines)
}

func day1_2() {
	lines := ReadFileToLines("./data/day1_data.txt")
	readLines(lines)
}

func isNumber(letter uint8) bool {
	return letter >= '1' && letter <= '9'
}

func readLines(lines []string) {
	total := 0
	for _, line := range lines {
		//log.Print("Running line: ", line)
		firstNumber := uint8(0)
		lastNumber := uint8(0)
		for i := 0; i < len(line); i++ {
			letter := line[i]
			if isNumber(letter) {
				firstNumber = letter
			} else {
				for key, val := range digits {
					digitLen := len(key)
					if len(line) >= i+digitLen {
						piece := line[i:(i + digitLen)]
						//log.Print("piece: ", piece, " key: ", key)
						if piece == key {
							firstNumber = uint8(val)
						}
					}
				}
			}
			if firstNumber > 0 {
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			letter := line[i]
			if isNumber(letter) {
				lastNumber = letter
			} else {
				for key, val := range digits {
					digitLen := len(key)
					if i-digitLen >= 0 {
						piece := line[(i - digitLen + 1) : i+1]
						//log.Print("piece: ", piece, " key: ", key)
						if piece == key {
							lastNumber = uint8(val)
						}
					}
				}
			}
			if lastNumber > 0 {
				break
			}
		}

		newNumber, _ := strconv.Atoi(fmt.Sprintf("%s%s", string(firstNumber), string(lastNumber)))
		total += newNumber
		//log.Print(line, "->", numberString, "->", newNumber, " total: ", total)
		//log.Printf("line: %s first: %s last %s newNumber: %d total: %d", line, string(firstNumber), string(lastNumber), newNumber, total)
	}
	log.Printf("\n------------------\nTotal: %d\n------------------\n", total)
}
