package day7

import (
	"bufio"
	"fmt"
	"os"
)

func initFile(fileName string) *bufio.Scanner {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanLines)
	return fileScanner
}

func Run1() {
	fileScanner := initFile("./internal/day7/input.txt")

	game := NewCamelPokerGame()
	for fileScanner.Scan() {
		line := fileScanner.Text()
		hand := NewHand(line, false)
		game.AddHand(hand)
	}
	fmt.Printf("Result 1 : %d\n", game.TotalWinningPrice())
}

func Run2() {
	fileScanner := initFile("./internal/day7/input.txt")

	game := NewCamelPokerGame()
	for fileScanner.Scan() {
		line := fileScanner.Text()
		hand := NewHand(line, true)
		game.AddHand(hand)
	}
	fmt.Printf("Result 2 : %d\n", game.TotalWinningPrice())
}
