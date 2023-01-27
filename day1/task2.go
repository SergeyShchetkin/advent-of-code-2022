package main

import (
	"advent_of_code_2022/common"
	"bufio"
	"fmt"
	"sort"
)

func getElfCalories() <-chan int {
	result := make(chan int)

	f := common.OpenFile("input_data.txt")
	sc := bufio.NewScanner(f)

	go func() {
		var elfCalories int
		defer close(result)
		defer f.Close()

		for sc.Scan() {
			partCalories := sc.Text()
			if partCalories == "" {
				result <- elfCalories
				elfCalories = 0
			} else {
				if calories, err := common.StrToInt(partCalories); err != nil {
					panic(err)
				} else {
					elfCalories += calories
				}
			}
		}
	}()

	return result
}

func main() {
	calories := make([]int, 0)
	for elfCalories := range getElfCalories() {
		calories = append(calories, elfCalories)
	}

	sort.Ints(calories)
	fmt.Println(common.SliceSum(calories[len(calories)-3:]))
}
