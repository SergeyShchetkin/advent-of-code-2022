package main

import (
	"advent_of_code_2022/common"
	"advent_of_code_2022/day3/base"
	"bufio"
	"fmt"
)

func calcDuplicatePriority(seq1, seq2 []byte) {
	defer base.Wg.Done()
	for _, symbol := range seq1 {
		if base.InSequence(symbol, seq2) {
			base.Mx.Lock()
			base.PrioritiesSum += base.GetSymbolPriority(symbol)
			base.Mx.Unlock()
			break
		}
	}
}

func main() {
	f := common.OpenFile("input_data.txt")
	sc := bufio.NewScanner(f)
	defer f.Close()

	for sc.Scan() {
		pack := []byte(sc.Text())
		part1, part2 := pack[:len(pack)/2], pack[len(pack)/2:]
		base.Wg.Add(1)
		go calcDuplicatePriority(part1, part2)
	}

	base.Wg.Wait()
	fmt.Printf("Sum priorities = %d\n", base.PrioritiesSum)
}
