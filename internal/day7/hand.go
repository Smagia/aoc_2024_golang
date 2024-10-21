package day7

import (
	"sort"
	"strconv"
	"strings"
)

func getCharValue(card rune, jolly bool) int {
	switch card {
	case '2', '3', '4', '5', '6', '7', '8', '9':
		return int(card - '0')
	case 'T':
		return 10
	case 'J':
		if jolly {
			return 1
		}
		return 11
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		return 0
	}
}

func calculateScoreWithJ(cards string) int {
	charCount := make(map[rune]int)
	keys := make([]rune, 0)
	jCount := 0
	for _, char := range cards {
		if char == 'J' {
			jCount++
		} else {
			charCount[char]++
		}
		if !strings.Contains(string(keys), string(char)) {
			keys = append(keys, char)

		}
	}

	sort.Slice(keys, func(i, j int) bool {
		return charCount[keys[i]] > charCount[keys[j]]
	})

	fistCharSum := charCount[keys[0]] + jCount
	if fistCharSum == 5 {
		return 6
	} else if fistCharSum == 4 {
		return 5
	} else if fistCharSum == 3 && charCount[keys[1]] == 2 {
		return 4
	} else if fistCharSum == 3 {
		return 3
	} else if charCount[keys[0]] == 2 && (jCount == 2 || charCount[keys[1]] == 2) {
		return 2
	} else if charCount[keys[0]] == 2 || jCount == 1 {
		return 1
	}

	return 0
}

func calculateScore(cards string, jolly bool) int {
	if jolly {
		return calculateScoreWithJ(cards)
	}

	charCount := make(map[rune]int)

	for _, char := range cards {
		charCount[char]++
	}

	alreadyFound := 0
	for _, count := range charCount {
		switch count {
		case 5:
			return 10
		case 4:
			return 8
		case 3:
			if alreadyFound != 0 && alreadyFound == 2 {
				return 7
			} else {
				alreadyFound = 6
			}
		case 2:
			if alreadyFound != 0 {
				if alreadyFound == 2 {
					return 4
				} else if alreadyFound == 6 {
					return 7
				}
			} else {
				alreadyFound = 2
			}
		}
	}
	if alreadyFound != 0 {
		return alreadyFound
	}
	return 0
}

type Hand struct {
	Cards string
	Bid   int64
	Score int
	Rank  int
	Jolly bool
}

func NewHand(line string, jolly bool) *Hand {
	data := strings.Split(line, " ")
	if len(data) != 2 {
		panic("Wrong input!")
	}
	bid, err := strconv.ParseInt(data[1], 10, 0)
	if err != nil {
		panic("Wrong Bid!")
	}
	score := calculateScore(data[0], jolly)
	return &Hand{
		Cards: data[0],
		Bid:   bid,
		Score: score,
		Rank:  0,
		Jolly: jolly,
	}
}

type ByCards []*Hand

func (a ByCards) Len() int { return len(a) }
func (a ByCards) Less(i, j int) bool {
	if a[i].Score == a[j].Score {
		for index := 0; index < 5; index++ {
			if a[i].Cards[index] != a[j].Cards[index] {
				return getCharValue(rune(a[i].Cards[index]), a[i].Jolly) < getCharValue(rune(a[j].Cards[index]), a[j].Jolly)
			}
		}
	}
	return a[i].Score < a[j].Score
}
func (a ByCards) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
