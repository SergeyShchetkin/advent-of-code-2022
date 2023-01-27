package main

import (
	"advent_of_code_2022/common"
	"advent_of_code_2022/day13/base"
	"fmt"
)

func main() {
	fmt.Printf("sum = %d\n", common.SliceSum(base.ReadData("input_data.txt").GetCorrectOrders()))
}
