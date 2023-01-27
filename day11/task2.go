package main

import (
	"advent_of_code_2022/day11/base"
	"fmt"
)

func main() {
	fmt.Printf("Business level = %d\n", base.ReadData("input_data.txt").GetBusinessLevel(10000, true))
}
