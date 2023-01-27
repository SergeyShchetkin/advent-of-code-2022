package main

import (
	"advent_of_code_2022/day5/base"
	"fmt"
)

func main() {
	base.ReadAndParseData("input_data.txt")
	for _, action := range base.Actions {
		for i := 1; i <= action.CountCrate; i++ {
			newLnFromStack := len(base.DataStacks[action.FromStack]) - 1
			base.DataStacks[action.ToStack] = append(
				base.DataStacks[action.ToStack],
				base.DataStacks[action.FromStack][newLnFromStack],
			)
			base.DataStacks[action.FromStack] = base.DataStacks[action.FromStack][:newLnFromStack]
		}
	}

	fmt.Println(base.GetVertexDataStack())
}
