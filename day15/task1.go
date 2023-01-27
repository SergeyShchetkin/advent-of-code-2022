package main

import (
	"advent_of_code_2022/day15/base"
	"fmt"
)

func main() {
	fmt.Printf("count = %d\n", base.ReadData("input_data.txt").CalcLineNotBeaconPositions())
}
