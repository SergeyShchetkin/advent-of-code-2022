package base

import (
	"advent_of_code_2022/common"
	"sync"
)

type place map[int][]int

func (p place) GetMaxTreeScore() int {
	var (
		max    int
		scores []int
		mx     sync.Mutex
		wg     sync.WaitGroup
	)

	pLn := len(p)
	for num, trees := range p {
		trees := trees
		wg.Add(1)

		go func(pos int) {
			defer wg.Done()
			for n, t := range trees {
				var (
					score int
					tmp   int
				)

				tLn := len(trees)
				if n != 0 && n != tLn-1 && pos != 0 && pos != pLn-1 {
					for i := n - 1; i >= 0; i-- {
						tmp++
						if trees[i] >= t || i == 0 {
							score = tmp
							tmp = 0
							break
						}
					}
					for i := n + 1; i < tLn; i++ {
						tmp++
						if trees[i] >= t || i == tLn-1 {
							score *= tmp
							tmp = 0
							break
						}
					}
					for j := pos - 1; j >= 0; j-- {
						tmp++
						if p[j][n] >= t || j == 0 {
							score *= tmp
							tmp = 0
							break
						}
					}
					for j := pos + 1; j < pLn; j++ {
						tmp++
						if p[j][n] >= t || j == pLn-1 {
							score *= tmp
							tmp = 0
							break
						}
					}

					mx.Lock()
					scores = append(scores, score)
					mx.Unlock()
				}
			}
		}(num)
	}

	wg.Wait()
	for _, score := range scores {
		if score > max {
			max = score
		}
	}

	return max
}

func (p place) GetVisibleTrees() int {
	var (
		countVisibleTrees int
		mx                sync.Mutex
		wg                sync.WaitGroup
	)

	pLn := len(p)
	for num, trees := range p {
		trees := trees
		if num == 0 || num == pLn-1 {
			countVisibleTrees += len(trees)
			continue
		}

		wg.Add(1)
		go func(pos int) {
			var inc int
			defer wg.Done()

			for n, t := range trees {
				var b = true
				tLn := len(trees)
				if n == 0 || n == tLn-1 {
					inc++
					b = false
				} else if b {
					for i := 0; i < n; i++ {
						b = b && trees[i] < t
					}
					if !b {
						b = true
						for i := n + 1; i < tLn; i++ {
							b = b && trees[i] < t
						}
					}
					if !b {
						b = true
						for j := 0; j < pos; j++ {
							b = b && p[j][n] < t
						}
					}
					if !b {
						b = true
						for j := pos + 1; j < pLn; j++ {
							b = b && p[j][n] < t
						}
					}
				}

				if b {
					inc++
				}
			}

			if inc > 0 {
				mx.Lock()
				countVisibleTrees += inc
				mx.Unlock()
			}
		}(num)
	}

	wg.Wait()
	return countVisibleTrees
}

func Read(filePath string) place {
	var (
		data       = make(place)
		dataReader = &common.DataReader{}
		step       int
	)

	for line := range dataReader.Read(filePath) {
		for _, tree := range line {
			height, _ := common.StrToInt(string(tree))
			data[step] = append(data[step], height)
		}
		step++
	}

	return data
}
