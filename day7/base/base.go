package base

import (
	"advent_of_code_2022/common"
	"fmt"
	"strings"
)

const (
	MaxCatalogSize         = 100000
	DiskSize               = 70000000
	UpdatePackSize         = 30000000
	StartCommandSymbol     = "$"
	ListCommand            = "ls"
	BackStepCommandPostfix = ".."
	StepCommand            = "cd"
	DirPrefix              = "dir"
)

type file struct {
	name string
	size int
}

type Directory struct {
	Files  []file
	Parent int
}

func (d *Directory) Size() int {
	var size int
	for _, fl := range d.Files {
		size += fl.size
	}

	return size
}

var Directories = make(map[int]*Directory)

func ReadAndParseData(filePath string) {
	var (
		step                int
		currentDirectory    int
		directoryDataOutput bool
		parentDirectory     int
		dataReader          = &common.DataReader{}
	)

	for line := range dataReader.Read(filePath) {
		lineParts := strings.Split(line, " ")
		if isCommandLine(line) {
			step++
			directoryDataOutput = isListCommand(line)
			if isStepCommand(line) {
				parentDirectory = currentDirectory
				dir := lineParts[len(lineParts)-1]
				if dir != BackStepCommandPostfix {
					currentDirectory = step
					if _, ok := Directories[currentDirectory]; !ok {
						Directories[currentDirectory] = &Directory{
							Files:  make([]file, 0),
							Parent: parentDirectory,
						}
					}
				} else {
					currentDirectory = Directories[currentDirectory].Parent
				}
			}
		} else if directoryDataOutput && lineParts[0] != DirPrefix {
			fileSize, _ := common.StrToInt(lineParts[0])
			fl := file{
				name: lineParts[1],
				size: fileSize,
			}
			Directories[currentDirectory].Files = append(Directories[currentDirectory].Files, fl)
		}
	}
}

func getCommandLineParts(str string) []string {
	command := strings.Replace(str, fmt.Sprintf("%s ", StartCommandSymbol), "", 1)
	return strings.Split(command, " ")
}

func isStepCommand(str string) bool {
	commandParts := getCommandLineParts(str)
	return commandParts[0] == StepCommand
}

func isListCommand(str string) bool {
	commandParts := getCommandLineParts(str)
	return commandParts[0] == ListCommand
}

func isCommandLine(str string) bool {
	s := strings.Split(str, " ")
	return s[0] == StartCommandSymbol
}
