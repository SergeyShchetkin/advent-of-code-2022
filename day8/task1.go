package main

import (
	"advent_of_code_2022/day8/base"
	"fmt"
)

func main() {
	fmt.Printf(
		"Count visible trees = %d\n",
		base.Read("input_data.txt").GetVisibleTrees(),
	)
}
