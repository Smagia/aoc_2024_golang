package day7

import (
	"sort"
)

type CamelPoker interface {
	AddHand(*Hand)
	TotalWinningPrice() int
}

type Game struct {
	Hands []*Hand
}

func NewCamelPokerGame() *Game {
	return &Game{
		Hands: make([]*Hand, 0),
	}
}

func (gm *Game) AddHand(hand *Hand) {
	gm.Hands = append(gm.Hands, hand)
}

func (gm *Game) TotalWinningPrice() int {
	result := 0
	sort.Sort(ByCards(gm.Hands))
	for index, hand := range gm.Hands {
		result += int(hand.Bid) * (index + 1)
	}

	return result
}
