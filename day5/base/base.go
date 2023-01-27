package base

import (
	"advent_of_code_2022/common"
	"sort"
	"strings"
)

type stack []string

type action struct {
	CountCrate int
	FromStack  int
	ToStack    int
}

var (
	DataStacks = make(map[int]stack, 0)
	Actions    []action
)

func GetVertexDataStack() string {
	var (
		vertex   string
		dataKeys []int
	)

	for k := range DataStacks {
		dataKeys = append(dataKeys, k)
	}

	sort.Ints(dataKeys)
	for _, k := range dataKeys {
		ds := DataStacks[k]
		vertex += ds[len(ds)-1]
	}

	return vertex
}

func buildStacks(stackNumbers string, stacksSrc []string) {
	stackNumbersData := strings.Split(stackNumbers, "")
	for _, sts := range stacksSrc {
		std := strings.Split(sts, "")
		for pos, item := range stackNumbersData {
			if num, err := common.StrToInt(item); err == nil {
				if !common.IsEmptyStr(strings.TrimSpace(std[pos])) {
					DataStacks[num] = append(DataStacks[num], std[pos])
				}
			}
		}
	}
}

func buildActions(a string) {
	actionData := strings.Split(a, " ")
	actionParams := make([]int, 0)
	for _, item := range actionData {
		if intItem, err := common.StrToInt(item); err == nil {
			actionParams = append(actionParams, intItem)
		}
	}
	Actions = append(Actions, action{
		CountCrate: actionParams[0],
		FromStack:  actionParams[1],
		ToStack:    actionParams[2],
	})
}

func isStackNumbers(s string) bool {
	s = strings.ReplaceAll(s, " ", "")
	for _, i := range s {
		if _, err := common.StrToInt(string(i)); err != nil {
			return false
		}
	}
	return true
}

func ReadAndParseData(filePath string) {
	var (
		stacksSrc    []string
		stackNumbers string
		dataReader   = &common.DataReader{}
		isDataStacks = true
	)

	for line := range dataReader.Read(filePath) {
		if common.IsEmptyStr(line) {
			isDataStacks = false
			buildStacks(stackNumbers, stacksSrc)
			continue
		}

		if isDataStacks {
			isStackNumbers(line)
			if !isStackNumbers(line) {
				stacksSrc = append(stacksSrc, line)
			} else {
				stackNumbers = line
			}
		} else {
			buildActions(line)
		}
	}

	for num, st := range DataStacks {
		DataStacks[num] = common.ArrayReverse(st)
	}
}
