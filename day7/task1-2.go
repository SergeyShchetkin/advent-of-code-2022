package main

import (
	"advent_of_code_2022/day7/base"
	"fmt"
	"sort"
)

func main() {
	var (
		directories     []int
		sizes           []int
		task1Size       int
		directoriesSize = make(map[int]int)
	)

	base.ReadAndParseData("input_data.txt")
	for k := range base.Directories {
		directories = append(directories, k)
	}

	sort.Ints(directories)
	for _, dir := range directories {
		dc := base.Directories[dir]
		size := dc.Size()
		directoriesSize[dir] = size
		for dc.Parent > 0 {
			directoriesSize[dc.Parent] += size
			dc = base.Directories[dc.Parent]
		}

	}

	for _, size := range directoriesSize {
		if size <= base.MaxCatalogSize {
			task1Size += size
		}
	}

	fmt.Printf("Task1 result: %d\n", task1Size)

	for _, size := range directoriesSize {
		sizes = append(sizes, size)
	}

	sort.Ints(sizes)
	for _, size := range sizes {
		fmt.Println(size)
		if base.DiskSize+size-directoriesSize[1] >= base.UpdatePackSize {
			fmt.Printf("Task2 result: %d\n", size)
			return
		}
	}
}
