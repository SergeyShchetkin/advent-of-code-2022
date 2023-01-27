package base

import (
	"advent_of_code_2022/common"
	"strings"
	"sync"
)

type Pair struct {
	PartStart int
	PartEnd   int
}

var (
	Wg                  sync.WaitGroup
	Mx                  sync.Mutex
	CountNestedPairs    int
	CountIntersectPairs int
	DataReader          = &common.DataReader{}
)

func InitPair(parts string) Pair {
	pair := Pair{}
	partsSl := strings.Split(parts, "-")
	pair.PartStart, _ = common.StrToInt(partsSl[0])
	pair.PartEnd, _ = common.StrToInt(partsSl[1])
	return pair
}

func IsNestedPairs(pairOne, pairTwo Pair) bool {
	return pairOne.PartStart <= pairTwo.PartStart && pairOne.PartEnd >= pairTwo.PartEnd ||
		pairTwo.PartStart <= pairOne.PartStart && pairTwo.PartEnd >= pairOne.PartEnd
}

func IsIntersectPairs(pairOne, pairTwo Pair) bool {
	return pairOne.PartStart >= pairTwo.PartStart && pairOne.PartStart <= pairTwo.PartEnd ||
		pairOne.PartEnd >= pairTwo.PartStart && pairOne.PartEnd <= pairTwo.PartEnd ||
		pairTwo.PartStart >= pairOne.PartStart && pairTwo.PartStart <= pairOne.PartEnd ||
		pairTwo.PartEnd >= pairOne.PartStart && pairTwo.PartEnd <= pairOne.PartEnd
}
