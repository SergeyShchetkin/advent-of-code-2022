package base

import "errors"

const (
	WinType  = "win"
	DrawType = "draw"
	LossType = "loss"
)

type Card struct {
	Points int
}

var (
	Stone    = Card{Points: 1}
	Paper    = Card{Points: 2}
	Scissors = Card{Points: 3}

	Cards = map[string]Card{
		"A": Stone, "X": Stone,
		"B": Paper, "Y": Paper,
		"C": Scissors, "Z": Scissors,
	}

	Superiority = map[Card]Card{
		Stone:    Scissors,
		Scissors: Paper,
		Paper:    Stone,
	}

	MatchTarget = map[string]string{
		"X": DrawType,
		"Y": LossType,
		"Z": WinType,
	}

	MatchPoints = map[string]int{
		WinType:  6,
		DrawType: 0,
		LossType: 3,
	}
)

func GetCardByAbbreviation(s string) Card {
	return Cards[s]
}

func GetSuperiorityByVal(c Card) (Card, error) {
	for key, val := range Superiority {
		if val == c {
			return key, nil
		}
	}

	return Card{}, errors.New("card not found")
}
