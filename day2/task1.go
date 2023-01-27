package main

import (
	"advent_of_code_2022/common"
	"advent_of_code_2022/day2/base"
	"bufio"
	"fmt"
	"strings"
)

func getMatchPoints(opponentCard, yourCard base.Card) int {
	if base.Superiority[yourCard] == opponentCard {
		return base.MatchPoints[base.WinType]
	}

	if base.Superiority[opponentCard] == yourCard {
		return base.MatchPoints[base.DrawType]
	}

	return base.MatchPoints[base.LossType]
}

func main() {
	var totalPoints int

	f := common.OpenFile("input_data.txt")
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		match := sc.Text()
		abbreviations := strings.Split(match, " ")
		yourCard := base.GetCardByAbbreviation(abbreviations[1])
		opponentCard := base.GetCardByAbbreviation(abbreviations[0])
		totalPoints += getMatchPoints(opponentCard, yourCard) + yourCard.Points
	}

	fmt.Printf("total points = %d\n", totalPoints)
}
