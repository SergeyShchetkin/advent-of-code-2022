package main

import (
	"advent_of_code_2022/day5/base"
	"fmt"
)

func main() {
	base.ReadAndParseData("input_data.txt")
	for _, action := range base.Actions {
		lnFromStack := len(base.DataStacks[action.FromStack])
		for i := action.CountCrate; i > 0; i-- {
			base.DataStacks[action.ToStack] = append(
				base.DataStacks[action.ToStack],
				base.DataStacks[action.FromStack][lnFromStack-i],
			)
		}
		base.DataStacks[action.FromStack] = base.DataStacks[action.FromStack][:lnFromStack-action.CountCrate]
	}

	fmt.Println(base.GetVertexDataStack())
}
