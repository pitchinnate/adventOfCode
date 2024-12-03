package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	day3_2()
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
