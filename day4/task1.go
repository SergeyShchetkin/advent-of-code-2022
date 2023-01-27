package main

import (
	"advent_of_code_2022/day4/base"
	"fmt"
	"strings"
)

func main() {
	for line := range base.DataReader.Read("input_data.txt") {
		base.Wg.Add(1)
		line := line
		go func() {
			defer base.Wg.Done()
			pairs := strings.Split(line, ",")
			pairOne := base.InitPair(pairs[0])
			pairTwo := base.InitPair(pairs[1])

			if base.IsNestedPairs(pairOne, pairTwo) {
				base.Mx.Lock()
				base.CountNestedPairs++
				base.Mx.Unlock()
			}
		}()
	}

	base.Wg.Wait()
	fmt.Printf("Count nested pairs = %d\n", base.CountNestedPairs)
}
