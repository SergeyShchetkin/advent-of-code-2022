package base

import (
	"advent_of_code_2022/common"
	"fmt"
	"math"
	"strings"
)

const (
	right = "R"
	up    = "U"
	left  = "L"
	down  = "D"
)

type position struct {
	horizon int
	height  int
}

func (p *position) toString() string {
	return fmt.Sprintf("%d:%d", p.horizon, p.height)
}

func (t *position) move(h *position) {
	var (
		horizon = 1
		height  = 1
	)

	if math.Abs(float64(t.horizon-h.horizon)) <= 1 &&
		math.Abs(float64(t.height-h.height)) <= 1 {
		return
	}

	if h.height < t.height {
		height = -1
	}
	if h.horizon < t.horizon {
		horizon = -1
	}
	if h.horizon == t.horizon {
		t.height += 1 * height
		return
	}
	if h.height == t.height {
		t.horizon += 1 * horizon
		return
	}

	t.horizon += 1 * horizon
	t.height += 1 * height
}

func CountTailPosition(filePath string, tailLength int) int {
	var (
		dataReader = &common.DataReader{}
		positions  = make([]position, 10)
	)

	tailPositions := map[string]bool{}
	tailPositions[positions[tailLength].toString()] = true
	for line := range dataReader.Read(filePath) {
		parts := strings.Split(line, " ")
		direction := parts[0]
		steps, _ := common.StrToInt(parts[1])
		for i := 0; i < steps; i++ {
			switch direction {
			case right:
				positions[0].horizon++
			case left:
				positions[0].horizon--
			case up:
				positions[0].height++
			case down:
				positions[0].height--
			}
			for j := 1; j < len(positions); j++ {
				positions[j].move(&positions[j-1])
			}

			tailPositions[positions[tailLength].toString()] = true
		}
	}

	return len(tailPositions)
}
