package day7_test

import (
	"aoc/smagia/internal/day7"
	"sort"
	"testing"
)

func TestHandLogic(t *testing.T) {
	testCase := []struct {
		hand          string
		expectedScore int
	}{
		{"32T3K 765", 2},
		{"T55J5 684", 6},
		{"KK677 28", 4},
		{"KTJJT 220", 4},
		{"QQQJA 483", 6},
		{"QQQQQ 483", 10},
		{"AAQQQ 483", 7},
		{"QQQQA 483", 8},
	}

	expectedOrder := []string{
		"32T3K",
		"KTJJT",
		"KK677",
		"T55J5",
		"QQQJA",
		"AAQQQ",
		"QQQQA",
		"QQQQQ",
	}

	hands := make([]*day7.Hand, 0)
	for _, tc := range testCase {
		hand := day7.NewHand(tc.hand, false)
		hands = append(hands, hand)
		if hand.Score != tc.expectedScore {
			t.Fatalf("UncorrectComputation %s has %d not %d", tc.hand, tc.expectedScore, hand.Score)
		}
	}

	sort.Sort(day7.ByCards(hands))

	for index, expectedCard := range expectedOrder {
		if expectedCard != hands[index].Cards {
			t.Fatalf("Uncorrect Sorting obtained %s expected %s", hands[index].Cards, expectedCard)
		}
	}
}

func TestSecondHandLogic(t *testing.T) {
	testCase := []struct {
		hand          string
		expectedScore int
	}{
		{"JAAAA 483", 6},
		{"JJAAA 972", 6},
		{"JJJAA 972", 6},
		{"JJJJA 483", 6},
		{"JJJJJ 972", 6},
		{"AAAA4 972", 5},

		{"AAKKJ 972", 4},
		{"AA1JJ 972", 5},
		{"AA23J 972", 3},
		{"A123J 972", 1},
	}

	for _, tc := range testCase {
		hand := day7.NewHand(tc.hand, true)
		if hand.Score != tc.expectedScore {
			t.Fatalf("UncorrectComputation %s has %d not %d", tc.hand, tc.expectedScore, hand.Score)
		}
	}
}

func TestGame(t *testing.T) {
	testCase := [5]string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}
	expectedResult := 6440
	game := day7.NewCamelPokerGame()
	for _, tc := range testCase {
		hand := day7.NewHand(tc, false)
		game.AddHand(hand)
	}

	if expectedResult != game.TotalWinningPrice() {
		t.Fatalf("Uncorrect Result obtained %d expected %d", game.TotalWinningPrice(), expectedResult)
	}
}

func TestSecondGame(t *testing.T) {
	testCase := [5]string{
		"32T3K 765",
		"T55J5 684",
		"KK677 28",
		"KTJJT 220",
		"QQQJA 483",
	}
	expectedResult := 5905
	game := day7.NewCamelPokerGame()
	for _, tc := range testCase {
		hand := day7.NewHand(tc, true)
		game.AddHand(hand)
	}

	if expectedResult != game.TotalWinningPrice() {
		t.Fatalf("Uncorrect Result obtained %d expected %d", game.TotalWinningPrice(), expectedResult)
	}
}
