package day8

import (
	"regexp"
)

type Point struct {
	ID    string
	Left  string
	Right string
}

func NewPoint(point string) *Point {
	pattern := `(\w+)\s*=\s*\((\w+),\s*(\w+)\)`

	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(point)
	return &Point{
		ID:    matches[1],
		Left:  matches[2],
		Right: matches[3],
	}
}

type Road struct {
	Steps  []rune
	Points map[string]*Point
}

func NewRoad() *Road {
	return &Road{
		Steps:  []rune{},
		Points: make(map[string]*Point),
	}
}

func (road *Road) SetStep(steps string) {
	road.Steps = []rune(steps)
}

func (road *Road) AddPoint(point *Point) {
	road.Points[point.ID] = point
}

func (road *Road) Navigate(position string, step int) int {
	currentStep := ""
	char := road.Steps[step%len(road.Steps)]

	if char == 'L' {
		currentStep = road.Points[position].Left
	} else if char == 'R' {
		currentStep = road.Points[position].Right
	}

	if currentStep == "ZZZ" {
		return step + 1
	} else {
		return road.Navigate(currentStep, step+1)
	}
}
