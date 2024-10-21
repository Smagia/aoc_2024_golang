package day6

import "fmt"

func Run1() {
	testCase := []struct {
		receTime   int
		receRecord int
	}{
		{47, 207},
		{84, 1394},
		{74, 1209},
		{67, 1014},
	}

	boatList := make([]*Boat, 0)

	for _, tc := range testCase {
		race := NewRace(tc.receTime, tc.receRecord)
		boat := NewBoat("pippo")
		boat.RunBoat(race)
		boatList = append(boatList, boat)
	}

	result := 1
	for _, boat := range boatList {
		result = result * len(boat.TimesToWin)
	}

	fmt.Printf("Final result : %d\n", result)
}

func Run2() {
	testCase := []struct {
		receTime   int
		receRecord int
	}{
		{47847467, 207139412091014},
	}

	boatList := make([]*Boat, 0)

	for _, tc := range testCase {
		race := NewRace(tc.receTime, tc.receRecord)
		boat := NewBoat("pippo")
		boat.RunBoat(race)
		boatList = append(boatList, boat)
	}

	result := int64(1)
	for _, boat := range boatList {
		result = result * int64(len(boat.TimesToWin))
	}

	fmt.Printf("Final result : %d\n", result)
}
