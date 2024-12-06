package main

import (
	"log"
	"math"
	"strconv"
	"strings"
)

type Rule struct {
	Left  int
	Right int
}
type Update struct {
	PageNumbers []int
}

func (update *Update) getMiddle() int {
	count := len(update.PageNumbers)
	middleIndex := int(math.Ceil(float64(count)/2.0)) - 1
	return update.PageNumbers[middleIndex]
}
func IntSliceSearch(numbers []int, search int) (int, bool) {
	for i, val := range numbers {
		if val == search {
			return i, true
		}
	}

	return -1, false
}
func (update *Update) Validate(rules []Rule) int {
	isValid := false
	for !isValid {
		for _, rule := range rules {
			leftIndex, ok := IntSliceSearch(update.PageNumbers, rule.Left)
			rightIndex, ok2 := IntSliceSearch(update.PageNumbers, rule.Right)
			if !ok || !ok2 {
				continue
			}
			if leftIndex > rightIndex {
				temp := update.PageNumbers[leftIndex]
				update.PageNumbers[leftIndex] = update.PageNumbers[rightIndex]
				update.PageNumbers[rightIndex] = temp
			}
		}
		isValid, _ = update.IsValid(rules)
	}
	return update.getMiddle()
}
func (update *Update) IsValid(rules []Rule) (bool, int) {
	isValid := true
	for _, rule := range rules {
		leftIndex, ok := IntSliceSearch(update.PageNumbers, rule.Left)
		rightIndex, ok2 := IntSliceSearch(update.PageNumbers, rule.Right)

		//log.Printf("Rule: %d | %d", rule.Left, rule.Right)
		//log.Print(leftIndex, ok)
		//log.Print(rightIndex, ok2)

		if !ok || !ok2 {
			continue
		}
		if leftIndex > rightIndex {
			isValid = false
			break
		}
	}

	//if isValid {
	//	log.Print("valid update", update)
	//	log.Print("valid update middle", update.getMiddle())
	//}

	return isValid, update.getMiddle()
}

func day5_1() {
	lines := ReadFileToLines("./data/day5_data.txt")
	rules := []string{}
	section2 := []string{}
	useSection2 := false
	for _, line := range lines {
		if line == "" {
			useSection2 = true
			continue
		}
		if !useSection2 {
			rules = append(rules, line)
		} else {
			section2 = append(section2, line)
		}
	}

	allRules := []Rule{}
	for _, rule := range rules {
		pieces := strings.Split(rule, "|")
		newRule := Rule{}
		newRule.Left, _ = strconv.Atoi(pieces[0])
		newRule.Right, _ = strconv.Atoi(pieces[1])
		allRules = append(allRules, newRule)
	}

	allUpdates := []Update{}
	for _, section := range section2 {
		newUpdate := Update{}
		pieces := strings.Split(section, ",")
		for _, piece := range pieces {
			number, _ := strconv.Atoi(piece)
			newUpdate.PageNumbers = append(newUpdate.PageNumbers, number)
		}
		allUpdates = append(allUpdates, newUpdate)
	}

	totalSum := 0
	for _, update := range allUpdates {
		valid, val := update.IsValid(allRules)
		if valid {
			totalSum += val
		}
	}
	log.Printf("Total Sum is : %d", totalSum)
}

func day5_2() {
	lines := ReadFileToLines("./data/day5_data.txt")
	rules := []string{}
	section2 := []string{}
	useSection2 := false
	for _, line := range lines {
		if line == "" {
			useSection2 = true
			continue
		}
		if !useSection2 {
			rules = append(rules, line)
		} else {
			section2 = append(section2, line)
		}
	}

	allRules := []Rule{}
	for _, rule := range rules {
		pieces := strings.Split(rule, "|")
		newRule := Rule{}
		newRule.Left, _ = strconv.Atoi(pieces[0])
		newRule.Right, _ = strconv.Atoi(pieces[1])
		allRules = append(allRules, newRule)
	}

	allUpdates := []Update{}
	for _, section := range section2 {
		newUpdate := Update{}
		pieces := strings.Split(section, ",")
		for _, piece := range pieces {
			number, _ := strconv.Atoi(piece)
			newUpdate.PageNumbers = append(newUpdate.PageNumbers, number)
		}
		allUpdates = append(allUpdates, newUpdate)
	}

	invalidUpdates := []Update{}
	for _, update := range allUpdates {
		valid, _ := update.IsValid(allRules)
		if !valid {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	totalSum := 0
	for _, update := range invalidUpdates {
		totalSum += update.Validate(allRules)
	}

	log.Printf("Total Sum is : %d", totalSum)
}
