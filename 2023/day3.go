package main

import (
	"log"
	"strconv"
)

type Position struct {
	X int
	Y int
}

func (p *Position) Equals(pos Position) bool {
	return p.X == pos.X && p.Y == pos.Y
}

type GearNumber struct {
	Value             int
	Positions         []Position
	TouchingPositions []Position
}

type Gear struct {
	Position Position
}

func day3_1() {
	lines := ReadFileToLines("./data/day3_data.txt")
	total := 0
	for rowIndex, row := range lines {
		newNumber := ""
		numberTouching := false
		row += "."
		numberPositions := []Position{}
		for columnIndex, letter := range row {
			if letter >= '0' && letter <= '9' {
				numberPositions = append(numberPositions, Position{columnIndex, rowIndex})
				newNumber += string(letter)
				touching, _ := TouchingSymbol(lines, columnIndex, rowIndex)
				log.Printf("x: %d y: %d touching: %v", columnIndex, rowIndex, touching)
				if touching {
					numberTouching = true
				}
			} else {
				if len(newNumber) > 0 {
					log.Printf("new number found: %s is touching %v", newNumber, numberTouching)
					if numberTouching {
						toNumber, _ := strconv.Atoi(newNumber)
						total += toNumber
					}
				}
				numberTouching = false
				newNumber = ""
			}
		}
	}
	log.Printf("total: %d", total)
}

func day3_2() {
	lines := ReadFileToLines("./data/day3_data.txt")
	total := 0
	gearNumbers := []GearNumber{}
	gears := []Gear{}
	for rowIndex, row := range lines {
		newNumber := ""
		numberTouching := false
		row += "."
		numberPositions := []Position{}
		numberTouchingPositions := []Position{}
		for columnIndex, letter := range row {
			if letter >= '0' && letter <= '9' {
				numberPositions = append(numberPositions, Position{columnIndex, rowIndex})
				newNumber += string(letter)
				touching, touchingPositions := TouchingSymbol(lines, columnIndex, rowIndex)
				numberTouchingPositions = append(numberTouchingPositions, touchingPositions...)
				if touching {
					numberTouching = true
				}
			} else {
				if len(newNumber) > 0 {
					log.Printf("new number found: %s is touching %v", newNumber, numberTouching)
					if numberTouching {
						toNumber, _ := strconv.Atoi(newNumber)
						gearNumbers = append(gearNumbers, GearNumber{
							Value:             toNumber,
							Positions:         numberPositions,
							TouchingPositions: numberTouchingPositions,
						})
					}
				}
				numberPositions = []Position{}
				numberTouchingPositions = []Position{}
				numberTouching = false
				newNumber = ""
				if letter == '*' {
					gears = append(gears, Gear{Position: Position{X: columnIndex, Y: rowIndex}})
				}
			}
		}
	}
	//log.Printf("numbers %#v", gearNumbers)
	//log.Printf("gears %#v", gears)

	for _, gear := range gears {
		connected := []int{}
		for _, number := range gearNumbers {
			for _, pos := range number.TouchingPositions {
				if pos.Equals(gear.Position) {
					connected = append(connected, number.Value)
					break
				}
			}
		}
		if len(connected) == 2 {
			total += (connected[0] * connected[1])
		}
		//log.Printf("gear x: %d y: %d connected to: %#v", gear.Position.X, gear.Position.Y, connected)
	}

	log.Printf("total %#v", total)
}

func TouchingSymbol(lines []string, x int, y int) (bool, []Position) {
	touchingPositions := []Position{}
	touching := false
	maxX := len(lines)
	maxY := len(lines[0])
	for rIndex := -1; rIndex <= 1; rIndex++ {
		for cIndex := -1; cIndex <= 1; cIndex++ {
			newX := x + rIndex
			newY := y + cIndex
			if newX >= 0 && newX < maxX && newY >= 0 && newY < maxY && (newX != x || newY != y) {
				//log.Printf("x: %d y: %d checking x: %d y: %d", x, y, newX, newY)
				letter := lines[newY][newX]
				//log.Printf("x: %d y: %d checking x: %d y: %d val: %s", x, y, newX, newY, string(letter))
				if letter == '!' || letter == '@' || letter == '#' || letter == '$' || letter == '%' || letter == '+' || letter == '^' ||
					letter == '&' || letter == '*' || letter == '(' || letter == '/' || letter == '\\' || letter == '-' || letter == '_' ||
					letter == '=' || letter == '~' || letter == '`' || letter == '|' || letter == ';' || letter == ':' || letter == ',' ||
					letter == ')' || letter == '[' || letter == ']' || letter == '{' || letter == '}' || letter == '<' || letter == '>' ||
					letter == '?' {
					touching = true
					if letter == '*' {
						touchingPositions = append(touchingPositions, Position{newX, newY})
					}
				}
			}
		}
	}

	return touching, touchingPositions
}
