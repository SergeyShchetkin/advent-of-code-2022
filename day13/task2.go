package main

import (
	"advent_of_code_2022/day13/base"
	"fmt"
)

func main() {
	fmt.Printf("signal = %d\n", base.ReadData("input_data.txt").GetDistressSignalCode())
}
