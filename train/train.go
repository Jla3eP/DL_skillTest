package train

import (
	"DL_skillTest/sTime"
	"bufio"
	"bytes"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Voyages []Voyage

type Voyage struct {
	number           int
	departureStation int
	arrivalStation   int
	price            float64
	departureTime    sTime.Time
	arrivalTime      sTime.Time
}

func (t Voyage) GetPrice() float64 {
	return t.price
}

func (t Voyage) GetDepTime() sTime.Time {
	return t.departureTime
}

func (t Voyage) GetArrTime() sTime.Time {
	return t.arrivalTime
}

func UnmarshalBytesToVoyage(data []byte) Voyage {
	t := Voyage{}
	fields := strings.Split(string(data), ";")

	t.number, _ = strconv.Atoi(fields[0])
	t.departureStation, _ = strconv.Atoi(fields[1])
	t.arrivalStation, _ = strconv.Atoi(fields[2])
	t.price, _ = strconv.ParseFloat(fields[3], 32)

	t.departureTime.Set([]byte(fields[4]))
	t.arrivalTime.Set([]byte(fields[5]))

	return t
}

var bytesPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 0, 50))
	},
}

func GetVoyages(filePath string) Voyages {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := Voyages{}
	for scanner.Scan() {
		data, _ := bytesPool.Get().([]byte)
		data = scanner.Bytes()
		result = append(result, UnmarshalBytesToVoyage(data))
		bytesPool.Put(data)
	}
	return result
}
