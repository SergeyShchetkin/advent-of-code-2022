package base

import (
	"advent_of_code_2022/common"
	"sort"
	"strings"
)

var anxiety = 1

const (
	startingItem = "Starting items"
	operation    = "Operation"
	test         = "Test"
	ifTrue       = "If true"
	ifFalse      = "If false"
)

type monkey struct {
	starting    []int
	operation   func(int) int
	test        func(int) bool
	trueMonkey  int
	falseMonkey int
}

type monkeys []monkey

func (ms monkeys) GetBusinessLevel(countRounds int, useAnxiety bool) int {
	var counts = make([]int, len(ms))
	for i := 0; i < countRounds; i++ {
		for j, m := range ms {
			for _, s := range m.starting {
				s = m.operation(s)
				if !useAnxiety {
					s /= 3
				} else {
					s %= anxiety
				}
				counts[j]++
				if m.test(s) {
					ms[m.trueMonkey].starting = append(ms[m.trueMonkey].starting, s)
				} else {
					ms[m.falseMonkey].starting = append(ms[m.falseMonkey].starting, s)
				}
			}
			ms[j].starting = make([]int, 0)
		}
	}

	sort.Ints(counts)
	return counts[len(counts)-2] * counts[len(counts)-1]
}

func ReadData(filePath string) monkeys {
	var (
		dataReader = &common.DataReader{}
		m          = monkey{}
		ms         = monkeys{}
	)

	for line := range dataReader.Read(filePath) {
		line := strings.TrimSpace(line)
		if common.IsEmptyStr(line) {
			ms = append(ms, m)
			m = monkey{}
			continue
		}

		parts := strings.Split(line, ":")
		if parts[0] == startingItem {
			item := strings.TrimSpace(parts[1])
			for _, i := range strings.Split(item, ",") {
				op, _ := common.StrToInt(strings.TrimSpace(i))
				m.starting = append(m.starting, op)
			}
		}

		if parts[0] == operation {
			item := strings.TrimSpace(strings.Replace(parts[1], "new = old", "", 1))
			op := item[0]
			val, err := common.StrToInt(item[2:])
			if err != nil {
				m.operation = func(old int) int {
					return old * old
				}
			} else {
				if op == '*' {
					m.operation = func(old int) int {
						return old * val
					}
				} else {
					m.operation = func(old int) int {
						return old + val
					}
				}
			}
		}

		if parts[0] == test {
			item := strings.TrimSpace(strings.Replace(parts[1], "divisible by", "", 1))
			i, _ := common.StrToInt(item)
			anxiety *= i
			m.test = func(old int) bool {
				return old%i == 0
			}
		}

		if parts[0] == ifTrue || parts[0] == ifFalse {
			item := strings.TrimSpace(strings.Replace(parts[1], "throw to monkey", "", 1))
			i, _ := common.StrToInt(item)
			if parts[0] == ifTrue {
				m.trueMonkey = i
			} else {
				m.falseMonkey = i
			}
		}
	}

	ms = append(ms, m)
	return ms
}
