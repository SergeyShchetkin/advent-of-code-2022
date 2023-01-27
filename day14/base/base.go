package base

import (
	"advent_of_code_2022/common"
	"sort"
	"strings"
)

const (
	stoneDelimiter     = "->"
	pointDelimiter     = ","
	stoneId            = "#"
	airId              = "."
	sandId             = "o"
	sandStartPositionX = 500
)

var toTopPoint bool

type stone struct {
	x int
	y int
}

type sand struct {
	x, y          int
	allowDiagonal bool
}

func newStone(point string) stone {
	p := strings.Split(point, pointDelimiter)
	x, _ := common.StrToInt(p[0])
	y, _ := common.StrToInt(p[1])
	return stone{
		x: x,
		y: y,
	}
}

type place map[int][]string

func (pl place) align(addY int) place {
	var maxLnY int
	for _, y := range pl {
		if maxLnY < len(y) {
			maxLnY = len(y)
		}
	}

	maxLnY += addY
	toTopPoint = addY > 0
	for x, y := range pl {
		if len(y) <= maxLnY {
			for i := len(y); i < maxLnY; i++ {
				if toTopPoint && i == maxLnY-1 {
					pl[x] = append(pl[x], stoneId)
				} else {
					pl[x] = append(pl[x], airId)
				}
			}
		}
	}

	return pl
}

func (pl place) GetSandCount() int {
	var (
		ctn        int
		maxX, minX int
		plKeys     []int
		startSand  = sand{
			x: sandStartPositionX,
		}
	)

	for key := range pl {
		plKeys = append(plKeys, key)
	}

	sort.Ints(plKeys)
	maxX = plKeys[len(plKeys)-1]
	minX = plKeys[0]

	currentSand := startSand
	checkColumns := func() {
		addColumn := func(x int) {
			pl[x] = make([]string, len(pl[currentSand.x]))
			for i := 0; i < len(pl[x]); i++ {
				if i < len(pl[x])-1 {
					pl[x][i] = airId
				} else {
					pl[x][i] = stoneId
				}
			}
		}

		if currentSand.x <= minX || currentSand.x >= maxX {
			minX--
			maxX++
			addColumn(minX)
			addColumn(maxX)
		}
	}

	for {
		if !toTopPoint && (currentSand.x <= minX || currentSand.x >= maxX || currentSand.y >= len(pl[maxX])-1) {
			break
		}

		if toTopPoint {
			checkColumns()
			if pl[currentSand.x][currentSand.y+1] == sandId {
				currentSand.allowDiagonal = true
			}
			if pl[startSand.x-1][startSand.y+1] == sandId &&
				pl[startSand.x+1][startSand.y+1] == sandId {
				ctn++
				break
			}
		}

		if pl[currentSand.x][currentSand.y+1] == airId {
			currentSand.y++
			currentSand.allowDiagonal = pl[currentSand.x-1][currentSand.y] == airId ||
				pl[currentSand.x+1][currentSand.y] == airId
		} else {
			if currentSand.allowDiagonal && pl[currentSand.x-1][currentSand.y+1] == airId {
				currentSand.x--
				currentSand.y++
			} else if currentSand.allowDiagonal && pl[currentSand.x+1][currentSand.y+1] == airId {
				currentSand.x++
				currentSand.y++
			} else {
				pl[currentSand.x][currentSand.y] = sandId
				ctn++
				currentSand = startSand
			}
		}
	}

	return ctn
}

func ReadData(filePath string, addY int) place {
	var (
		pl         = make(place)
		dataReader = &common.DataReader{}
	)

	for line := range dataReader.Read(filePath) {
		line = strings.Replace(line, " ", "", -1)
		stones := strings.Split(line, stoneDelimiter)
		addStoneY := func(x, y int) {
			if len(pl[x])-1 < y {
				for i := len(pl[x]); i <= y; i++ {
					pl[x] = append(pl[x], airId)
				}
			}
			pl[x][y] = stoneId
		}

		addStoneX := func(x int) {
			if _, ok := pl[x]; !ok {
				pl[x] = make([]string, 0)
			}
		}

		for i := 0; i < len(stones)-1; i++ {
			firstStone := newStone(stones[i])
			secondStone := newStone(stones[i+1])
			if firstStone.x != secondStone.x {
				for x := common.GetMin(firstStone.x, secondStone.x); x <= common.GetMax(firstStone.x, secondStone.x); x++ {
					addStoneX(x)
					addStoneY(x, firstStone.y)
				}
			} else {
				for y := common.GetMin(firstStone.y, secondStone.y); y <= common.GetMax(firstStone.y, secondStone.y); y++ {
					addStoneX(firstStone.x)
					addStoneY(firstStone.x, y)
				}
			}
		}
	}

	return pl.align(addY)
}
