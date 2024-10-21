package day6_test

import (
	"aoc/smagia/internal/day6"
	"testing"
)

func TestRun(t *testing.T) {
	testCase := []struct {
		receTime       int
		receRecord     int
		expectedResult int
	}{
		{7, 9, 4},
		{15, 40, 8},
		{30, 200, 9},
	}

	for _, tc := range testCase {
		race := day6.NewRace(tc.receTime, tc.receRecord)
		boat := day6.NewBoat("pippo")
		boat.RunBoat(race)

		if len(boat.TimesToWin) != tc.expectedResult {
			t.Fatalf("Error uncorrect result %d", len(boat.TimesToWin))
			for _, element := range boat.TimesToWin {
				t.Fatalf("Seconds found %d", element)
			}
		}
	}

}
