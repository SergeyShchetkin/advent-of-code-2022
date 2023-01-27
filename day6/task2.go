package main

import (
	"advent_of_code_2022/day6/base"
	"fmt"
)

func main() {
	var startPosition int
	for i := 0; i < len(base.DataStreamBuffer) - (base.PackageLength - 1); i++ {
		startPosition = i
		marker := base.Marker{Data: []byte(base.DataStreamBuffer[i:i+base.PackageLength])}
		if marker.IsDifferentSymbols() {
			startPosition += base.PackageLength
			break
		}
	}

	fmt.Printf("Start package position = %d\n", startPosition)
}
