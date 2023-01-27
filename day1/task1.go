package main

import (
	"advent_of_code_2022/common"
	"bufio"
	"fmt"
)

func main() {
	var (
		elfCalories    int
		maxElfCalories int
	)

	f := common.OpenFile("input_data.txt")
	sc := bufio.NewScanner(f)
	defer f.Close()

	for sc.Scan() {
		partCalories := sc.Text()
		if partCalories == "" {
			maxElfCalories = common.GetMax(elfCalories, maxElfCalories)
			elfCalories = 0
		} else {
			if calories, err := common.StrToInt(partCalories); err != nil {
				panic(err)
			} else {
				elfCalories += calories
			}
		}
	}

	fmt.Println(maxElfCalories)
}
