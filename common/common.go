package common

import (
	"bufio"
	"os"
	"strconv"
)

type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func StrToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func IsFloat64(val interface{}) bool {
	_, ok := val.(float64)
	return ok
}

func InArray[T comparable](a []T, x T) bool {
	for _, v := range a {
		if v == x {
			return true
		}
	}

	return false
}

func OpenFile(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return f
}

func GetMax[T Number](x T, y T) T {
	if x > y {
		return x
	}

	return y
}

func GetMin[T Number](x T, y T) T {
	if x < y {
		return x
	}

	return y
}

func IsEmptyStr(s string) bool {
	return len(s) == 0
}

func ArrayReverse[T any](a []T) []T {
	var reverse []T
	for i := len(a) - 1; i >= 0; i-- {
		reverse = append(reverse, a[i])
	}

	return reverse
}

func SliceSum[T Number](sl []T) T {
	var sum T
	for _, i := range sl {
		sum += i
	}

	return sum
}

// ----

type DataReader struct{}

func (dr *DataReader) Read(filePath string) <-chan string {
	lines := make(chan string)
	f := OpenFile(filePath)
	sc := bufio.NewScanner(f)

	go func() {
		defer close(lines)
		defer f.Close()
		for sc.Scan() {
			lines <- sc.Text()
		}
	}()

	return lines
}
