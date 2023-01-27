package base

import "sync"

var (
	Wg            sync.WaitGroup
	Mx            sync.Mutex
	PrioritiesSum int
	symbols       = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func InSequence(b byte, sequence []byte) bool {
	for _, item := range sequence {
		if item == b {
			return true
		}
	}

	return false
}

func GetSymbolPriority(symbol byte) int {
	for num, item := range symbols {
		if item == symbol {
			return num + 1
		}
	}

	return 0
}
