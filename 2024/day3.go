package main

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

func day3_1() {
	command := ReadFileToString("./data/day3_data.txt")
	r, _ := regexp.Compile("mul\\(\\d{1,3},\\d{1,3}\\)")
	matches := r.FindAllString(command, -1)
	totalProd := 0
	for _, m := range matches {
		inside := m[4 : len(m)-1]
		pieces := strings.Split(inside, ",")
		val1, _ := strconv.Atoi(pieces[0])
		val2, _ := strconv.Atoi(pieces[1])
		totalProd += val1 * val2
	}
	log.Printf("total: %d", totalProd)
}

func day3_2() {
	command := ReadFileToString("./data/day3_data.txt")
	r, _ := regexp.Compile("mul\\(\\d{1,3},\\d{1,3}\\)|don\\'t\\(\\)|do\\(\\)")
	matches := r.FindAllString(command, -1)
	totalProd := 0
	enabled := true
	for _, m := range matches {
		if m == "don't()" {
			enabled = false
			continue
		}
		if m == "do()" {
			enabled = true
			continue
		}
		if !enabled {
			continue
		}
		inside := m[4 : len(m)-1]
		pieces := strings.Split(inside, ",")
		val1, _ := strconv.Atoi(pieces[0])
		val2, _ := strconv.Atoi(pieces[1])
		totalProd += val1 * val2
	}
	log.Printf("total: %d", totalProd)
}
