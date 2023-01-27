package base

import (
	"advent_of_code_2022/common"
)

const (
	startSymbol = "S"
	endSymbol   = "E"
)

type position struct {
	x, y int
}

type place [][]rune

func (p place) getStartEndPositions(allStarts bool) (position, position, []position) {
	var (
		s  = position{}
		e  = position{}
		as = make([]position, 0)
	)

	for y := 0; y < len(p); y++ {
		for x := 0; x < len(p[y]); x++ {
			if string(p[y][x]) == startSymbol || (string(p[y][x]) == "a" && allStarts) {
				s.x = x
				s.y = y
				p[y][x] = 'a'
				as = append(as, s)
			} else if string(p[y][x]) == endSymbol {
				e.x = x
				e.y = y
				p[y][x] = 'z'
			}
		}
	}

	return s, e, as
}

func (p place) CalcPart2Steps() int {
	var (
		height         = len(p)
		width          = len(p[0])
		max            = width * height
		_, end, starts = p.getStartEndPositions(true)
	)

	for _, start := range starts {
		current := p.calcSteps(start, end)
		if current > 0 && current < max {
			max = current
		}
	}

	return max
}

func (p place) CalcPart1Steps() int {
	start, end, _ := p.getStartEndPositions(false)
	return p.calcSteps(start, end)
}

func (p place) calcSteps(start, end position) int {
	var (
		height    = len(p)
		width     = len(p[0])
		positions = []position{start}
		steps     = make([][]int, height)
	)

	for y := 0; y < height; y++ {
		steps[y] = make([]int, width)
		for x := 0; x < width; x++ {
			steps[y][x] = -1
		}
	}

	steps[start.y][start.x] = 0
	for len(positions) > 0 {
		current := positions[0]
		positions = positions[1:]

		if current == end {
			return steps[current.y][current.x]
		}

		possiblePoints := []position{
			{x: current.x, y: current.y - 1},
			{x: current.x, y: current.y + 1},
			{x: current.x - 1, y: current.y},
			{x: current.x + 1, y: current.y},
		}

		for _, next := range possiblePoints {
			if next.y < 0 || next.y == height || next.x < 0 ||
				next.x == width || steps[next.y][next.x] != -1 ||
				p[next.y][next.x]-p[current.y][current.x] > 1 {
				continue
			}

			steps[next.y][next.x] = steps[current.y][current.x] + 1
			positions = append(positions, next)
		}
	}

	return 0
}

func ReadData(filePath string) place {
	var (
		p          place
		dataReader = &common.DataReader{}
	)

	for line := range dataReader.Read(filePath) {
		p = append(p, []rune(line))
	}

	return p
}
