package base

import (
	"advent_of_code_2022/common"
	"math"
	"strings"
)

const (
	sensorLinePrefix  = "Sensor at"
	beaconLinePrefix  = "closest beacon is at"
	xLinePrefix       = "x="
	yLinePrefix       = "y="
	dataSeparator     = ":"
	positionSeparator = ","
	sensorSymbol      = "S"
	beaconSymbol      = "B"
	banBeaconSymbol   = "#"
	lineFindBeacon    = 2000000
)

type position struct {
	x int
	y int
}

type positions map[position]position

func calcDistance(sensor, beacon position) int {
	return int(math.Abs(float64(sensor.x-beacon.x)) +
		math.Abs(float64(sensor.y-beacon.y)))
}

func (p positions) add(sensor, beacon position) {
	p[sensor] = beacon
}

func (p positions) render() place {
	pl := make(place)
	checkMapPosition := func(y int) {
		if _, ok := pl[y]; !ok {
			pl[y] = make(map[int]string)
		}
	}

	checkMapPosition(lineFindBeacon)
	for sensor, beacon := range p {
		distance := calcDistance(sensor, beacon)
		checkMapPosition(sensor.y)
		pl[sensor.y][sensor.x] = sensorSymbol
		checkMapPosition(beacon.y)
		pl[beacon.y][beacon.x] = beaconSymbol
		for y := 0; y <= 2*distance; y++ {
			pY := sensor.y - distance + y
			if pY != lineFindBeacon {
				continue
			}
			for x := 0; x <= 2*distance; x++ {
				pX := sensor.x - distance + x
				if int(math.Abs(float64(y-distance))+math.Abs(float64(x-distance))) <= distance &&
					pl[pY][pX] != sensorSymbol && pl[pY][pX] != beaconSymbol {
					pl[pY][pX] = banBeaconSymbol
				}
			}
		}
	}

	return pl
}

func (pl place) CalcLineNotBeaconPositions() int {
	var count int
	for _, val := range pl[lineFindBeacon] {
		if val == banBeaconSymbol {
			count++
		}
	}

	return count
}

func newPosition(data string) position {
	d := strings.Split(data, positionSeparator)
	x, _ := common.StrToInt(d[0])
	y, _ := common.StrToInt(d[1])
	return position{
		x: x,
		y: y,
	}
}

type place map[int]map[int]string

func ReadData(filePath string) place {
	var (
		ps         = make(positions)
		dataReader = &common.DataReader{}
	)

	for line := range dataReader.Read(filePath) {
		line = strings.Replace(
			strings.Replace(
				strings.Replace(strings.Replace(line, yLinePrefix, "", 2),
					xLinePrefix, "", 2),
				beaconLinePrefix, "", 1),
			sensorLinePrefix, "", 1)
		data := strings.Split(strings.Replace(line, " ", "", -1), dataSeparator)
		sensor := newPosition(data[0])
		beacon := newPosition(data[1])
		ps.add(sensor, beacon)
	}

	return ps.render()
}
