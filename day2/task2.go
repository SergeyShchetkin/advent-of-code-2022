package main

import (
	"advent_of_code_2022/common"
	"advent_of_code_2022/day2/base"
	"bufio"
	"fmt"
	"strings"
)

func getYourCard(opponentCard base.Card, target string) base.Card {
	if target == base.WinType {
		card, err := base.GetSuperiorityByVal(opponentCard)
		if err != nil {
			panic(err)
		}
		return card
	}

	if target == base.DrawType {
		return base.Superiority[opponentCard]
	}

	return opponentCard
}

func main() {
	var totalPoints int

	f := common.OpenFile("input_data.txt")
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		match := strings.Split(sc.Text(), " ")
		opponentCard := base.GetCardByAbbreviation(match[0])
		matchTarget := base.MatchTarget[match[1]]
		totalPoints += base.MatchPoints[matchTarget] + getYourCard(opponentCard, matchTarget).Points
	}

	fmt.Printf("total points = %d\n", totalPoints)
}
