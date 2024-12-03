package main

import (
	"log"
	"math"
	"strconv"
	"strings"
)

type Report struct {
	Numbers []int
}

func (r Report) Safe() bool {
	increasing := true
	if r.Numbers[0] > r.Numbers[1] {
		increasing = false
	}
	for i, n := range r.Numbers {
		if i == 0 {
			continue
		}
		if n > r.Numbers[i-1] && !increasing {
			return false
		}
		if n < r.Numbers[i-1] && increasing {
			return false
		}
		diff := int(math.Abs(float64(n - r.Numbers[i-1])))
		if (diff < 1) || (diff > 3) {
			return false
		}
	}
	return true
}

func (r Report) ExpandSafe() bool {
	safe := r.Safe()
	if safe {
		return true
	}
	for x, _ := range r.Numbers {
		newReport := Report{}
		for y, n := range r.Numbers {
			if x != y {
				newReport.Numbers = append(newReport.Numbers, n)
			}
		}
		safe := newReport.Safe()
		if safe {
			return true
		}
	}

	return false
}

func day2_1() {
	lines := ReadFileToLines("./data/day2_data.txt")
	reports := getReports(lines)
	safeReports := 0
	for _, report := range reports {
		if report.Safe() {
			safeReports++
		}
	}
	log.Printf("Safe: %d", safeReports)
}

func day2_2() {
	lines := ReadFileToLines("./data/day2_data.txt")
	reports := getReports(lines)
	safeReports := 0
	for _, report := range reports {
		if report.ExpandSafe() {
			safeReports++
		}
	}
	log.Printf("Safe: %d", safeReports)
}

func getReports(lines []string) []Report {
	reports := []Report{}
	for _, line := range lines {
		if line == "" {
			continue
		}
		report := Report{}
		pieces := strings.Split(line, " ")
		for _, piece := range pieces {
			val, _ := strconv.Atoi(piece)
			report.Numbers = append(report.Numbers, val)
		}
		reports = append(reports, report)
	}
	return reports
}
