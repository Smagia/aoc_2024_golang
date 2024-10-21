package day8_test

import (
	"aoc/smagia/internal/day8"
	"testing"
)

func TestPoint(t *testing.T) {
	point := day8.NewPoint("XFB = (KTX, QQG)")
	if point.ID != "XFB" || point.Left != "KTX" || point.Right != "QQG" {
		t.Fatalf("Point is not correct : %s %s %s ", point.ID, point.Left, point.Right)
	}
}

func TestRoad(t *testing.T) {
	testCase := []string{
		"AAA = (BBB, BBB)",
		"BBB = (AAA, ZZZ)",
		"ZZZ = (ZZZ, ZZZ)",
	}

	road := day8.NewRoad()
	road.SetStep("LLR")
	for _, tc := range testCase {
		p := day8.NewPoint(tc)
		road.AddPoint(p)
	}

	result := road.Navigate("AAA", 0)
	if result != 6 {
		t.Fatalf("Uncorrect result: %d", result)
	}
}
