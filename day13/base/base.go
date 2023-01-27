package base

import (
	"advent_of_code_2022/common"
	"encoding/json"
	"sort"
	"strings"
	"sync"
)

const (
	startDecoderSignal = "[[2]]"
	endDecoderSignal   = "[[6]]"
)

type pack struct {
	src  string
	data interface{}
}

type pairPacks struct {
	left  *pack
	right *pack
}

type pairs struct {
	packs []pairPacks
	data  []interface{}
}

func (p *pack) parse() {
	var data interface{}
	if err := json.Unmarshal([]byte(p.src), &data); err != nil {
		panic(err)
	}
	p.data = data
}

func comparePacks(p1, p2 interface{}) int {
	var (
		p1Fl, p2Fl float64
		p1IsFl     = common.IsFloat64(p1)
		p2IsFl     = common.IsFloat64(p2)
	)

	if p1IsFl && p2IsFl {
		p1Fl = p1.(float64)
		p2Fl = p2.(float64)
		if p2Fl > p1Fl {
			return 1
		} else if p2Fl < p1Fl {
			return -1
		} else {
			return 0
		}
	}
	if p1IsFl {
		p1 = []interface{}{p1}
	}
	if p2IsFl {
		p2 = []interface{}{p2}
	}

	p1Sl := p1.([]interface{})
	p2Sl := p2.([]interface{})

	for len(p1Sl) > 0 && len(p2Sl) > 0 {
		fp1 := p1Sl[0]
		p1Sl = p1Sl[1:]
		fp2 := p2Sl[0]
		p2Sl = p2Sl[1:]
		res := comparePacks(fp1, fp2)
		if res != 0 {
			return res
		}
	}

	return len(p2Sl) - len(p1Sl)
}

func (pp *pairPacks) isCorrectPair() bool {
	if comparePacks(pp.left.data, pp.right.data) > 0 {
		return true
	}

	return false
}

func (p pairs) GetCorrectOrders() []int {
	var (
		mx      sync.Mutex
		wg      sync.WaitGroup
		correct []int
	)
	for num, pp := range p.packs {
		pp := pp
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			if pp.isCorrectPair() {
				mx.Lock()
				correct = append(correct, n)
				mx.Unlock()
			}
		}(num + 1)
	}

	wg.Wait()
	return correct
}

func (p pairs) GetDistressSignalCode() int {
	var (
		result = 1
		start  interface{}
		end    interface{}
	)
	json.Unmarshal([]byte(startDecoderSignal), &start)
	p.data = append(p.data, start)
	json.Unmarshal([]byte(endDecoderSignal), &end)
	p.data = append(p.data, end)

	sort.Slice(p.data, func(i, j int) bool {
		return comparePacks(p.data[i], p.data[j]) > 0
	})

	for num, item := range p.data {
		js, _ := json.Marshal(item)
		if string(js) == startDecoderSignal || string(js) == endDecoderSignal {
			result *= num + 1
		}
	}

	return result
}

func ReadData(filePath string) pairs {
	var (
		packNumber int
		dataReader = &common.DataReader{}
		pp         = pairPacks{}
		data       pairs
	)
	for line := range dataReader.Read(filePath) {
		line = strings.Replace(line, " ", "", -1)
		p := &pack{src: line}
		if common.IsEmptyStr(line) {
			packNumber = 0
			continue
		}

		p.parse()
		data.data = append(data.data, p.data)
		if packNumber == 0 {
			pp.left = p
		} else {
			pp.right = p
			data.packs = append(data.packs, pp)
			pp = pairPacks{}
		}

		packNumber = 1
	}

	return data
}
