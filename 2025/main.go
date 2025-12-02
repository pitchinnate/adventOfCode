package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	day2_2()
}

func ReadFileToLines(name string) []string {
	data, err := os.ReadFile(name)
	if err != nil {
		log.Fatal("error: ", err)
	}
	dataString := string(data)
	pieces := strings.Split(dataString, "\n")
	return pieces
}

func ReadFileToString(name string) string {
	data, err := os.ReadFile(name)
	if err != nil {
		log.Fatal("error: ", err)
	}
	dataString := string(data)
	return dataString
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
