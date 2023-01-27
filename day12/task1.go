package main

import (
	"advent_of_code_2022/day12/base"
	"fmt"
)

func main() {
	fmt.Printf("count steps = %d\n", base.ReadData("input_data.txt").CalcPart1Steps())
}
