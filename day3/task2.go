package main

import (
	"advent_of_code_2022/common"
	"advent_of_code_2022/day3/base"
	"bufio"
	"fmt"
)

func calcDuplicateGroupPriority(group [3]string) {
	defer base.Wg.Done()
	for _, symbol := range []byte(group[0]) {
		if base.InSequence(symbol, []byte(group[1])) && base.InSequence(symbol, []byte(group[2])) {
			base.Mx.Lock()
			base.PrioritiesSum += base.GetSymbolPriority(symbol)
			base.Mx.Unlock()
			break
		}
	}
}

func main() {
	var (
		step  int
		group [3]string
	)

	f := common.OpenFile("input_data.txt")
	sc := bufio.NewScanner(f)
	defer f.Close()

	for sc.Scan() {
		group[step] = sc.Text()
		if step == 2 {
			base.Wg.Add(1)
			go calcDuplicateGroupPriority(group)
			step = 0
		} else {
			step++
		}
	}

	base.Wg.Wait()
	fmt.Printf("Sum group priorities = %d\n", base.PrioritiesSum)
}
