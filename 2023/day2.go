package main

import (
	"log"
	"strconv"
	"strings"
)

var colorMaxes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

type GameRound struct {
	Blue  int
	Red   int
	Green int
}

type Game struct {
	Id       int
	Rounds   []GameRound
	Original string
}

func (g *Game) Valid() bool {
	for _, round := range g.Rounds {
		if round.Blue > colorMaxes["blue"] {
			return false
		}
		if round.Red > colorMaxes["red"] {
			return false
		}
		if round.Green > colorMaxes["green"] {
			return false
		}
	}
	return true
}

func (g *Game) Power() int {
	minRed := 0
	minBlue := 0
	minGreen := 0

	for _, round := range g.Rounds {
		if round.Red > minRed {
			minRed = round.Red
		}
		if round.Green > minGreen {
			minGreen = round.Green
		}
		if round.Blue > minBlue {
			minBlue = round.Blue
		}
	}

	return minGreen * minBlue * minRed
}

func day2_1() {
	lines := ReadFileToLines("./data/day2_data.txt")
	games := parseGames(lines)

	total := 0
	for _, game := range games {
		if game.Valid() {
			total += game.Id
		}
	}

	log.Print(total)
}

func day2_2() {
	lines := ReadFileToLines("./data/day2_data.txt")
	games := parseGames(lines)

	total := 0
	for _, game := range games {
		total += game.Power()
	}

	log.Print(total)
}

func parseGames(lines []string) []Game {
	games := []Game{}
	for index, line := range lines {
		pieces := strings.Split(line, ":")
		gamePieces := strings.Split(pieces[1], ";")
		newGame := Game{Id: index + 1, Original: line}
		for _, game := range gamePieces {
			individualPieces := strings.Split(strings.Trim(game, " "), " ")
			//log.Printf("%#v", individualPieces)
			round := GameRound{}
			for i := 0; i < len(individualPieces); i += 2 {
				color := strings.TrimSuffix(individualPieces[i+1], ",")
				count, _ := strconv.Atoi(individualPieces[i])
				//log.Print(color, count)
				switch color {
				case "blue":
					round.Blue = count
				case "red":
					round.Red = count
				case "green":
					round.Green = count
				}
			}
			//log.Printf("%#v", round)
			newGame.Rounds = append(newGame.Rounds, round)
		}
		games = append(games, newGame)
	}
	return games
}
