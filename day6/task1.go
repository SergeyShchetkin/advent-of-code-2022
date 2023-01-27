package main

import (
	"advent_of_code_2022/day6/base"
	"fmt"
)

func main() {
	var startPosition int
	for i := 0; i < len(base.DataStreamBuffer) - (base.MarkerLength - 1); i++ {
		startPosition = i
		marker := base.Marker{Data: []byte(base.DataStreamBuffer[i:i+base.MarkerLength])}
		if marker.IsDifferentSymbols() {
			startPosition += base.MarkerLength
			break
		}
	}

	fmt.Printf("Start marker position = %d\n", startPosition)
}
