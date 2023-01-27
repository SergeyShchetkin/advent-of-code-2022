package base

import (
	"advent_of_code_2022/common"
	"fmt"
	"strings"
)

const (
	noopOperation = "noop"
	addOperation  = "addx"
	pixelSymbol   = '#'
	pointSymbol   = '.'
	startIter     = 20
	stepIter      = 40
	maxIter       = 220
	spritePatLen  = 40
)

type signalPower struct {
	x          int
	iterNumber int
	sprites    []byte
	stepValues map[int]int
}

func (sp *signalPower) incIter() {
	sp.iterNumber++
}

func (sp *signalPower) getIter() int {
	return sp.iterNumber
}

func (sp *signalPower) addX(val int) {
	sp.x += val
}

func (sp *signalPower) getX() int {
	return sp.x
}

func (sp *signalPower) calcStep() {
	sp.stepValues[sp.getIter()] = sp.getIter() * sp.getX()
}

func (sp *signalPower) getStepValues() map[int]int {
	return sp.stepValues
}

func (sp *signalPower) addSpite() {
	var (
		s        byte
		position = (sp.getIter() - 1) % spritePatLen
		sprite   = make([]int, 3)
	)

	sprite[0] = sp.getX() - 1
	sprite[1] = sp.getX()
	sprite[2] = sp.getX() + 1

	if common.InArray(sprite, position) {
		s = pixelSymbol
	} else {
		s = pointSymbol
	}

	sp.sprites = append(sp.sprites, s)
}

func (sp *signalPower) renderSprites() {
	fmt.Println()
	for i, sym := range sp.sprites {
		num := i + 1
		fmt.Printf("%s", string(sym))
		if num%spritePatLen == 0 {
			fmt.Println()
		}
	}
}

var (
	dataReader = &common.DataReader{}
	operations = map[string]int{
		noopOperation: 1,
		addOperation:  2,
	}
)

func CalcSumPowerAndPrintSprites(filePath string) int {
	var (
		result int
		sp     = &signalPower{
			iterNumber: 1,
			x:          1,
			stepValues: make(map[int]int),
		}
	)

	for line := range dataReader.Read(filePath) {
		var val int

		parts := strings.Split(line, " ")
		operation := parts[0]
		if operation == addOperation {
			val, _ = common.StrToInt(parts[1])
		}

		opLen := operations[operation]
		for i := 0; i < opLen; i++ {
			sp.addSpite()
			if sp.getIter() == startIter ||
				(sp.getIter() > startIter && (sp.getIter()-startIter)%stepIter == 0) {
				sp.calcStep()
			}

			sp.incIter()
			if i == opLen-1 {
				sp.addX(val)
			}
		}
	}

	for iter, val := range sp.getStepValues() {
		if iter > maxIter {
			continue
		}
		result += val
	}

	sp.renderSprites()
	return result
}
