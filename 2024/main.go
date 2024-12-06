package main

import (
	"log"
	"os"
	"strings"
)

func main() {
	day5_2()
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
