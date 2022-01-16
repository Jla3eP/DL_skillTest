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

func (v Voyage) GetPrice() float64 {
	return v.price
}

func (v Voyage) GetDepTime() sTime.Time {
	return v.departureTime
}

func (v Voyage) GetArrTime() sTime.Time {
	return v.arrivalTime
}

func UnmarshalBytesToVoyage(data []byte) Voyage {
	v := Voyage{}
	fields := strings.Split(string(data), ";")

	v.number, _ = strconv.Atoi(fields[0])
	v.departureStation, _ = strconv.Atoi(fields[1])
	v.arrivalStation, _ = strconv.Atoi(fields[2])
	v.price, _ = strconv.ParseFloat(fields[3], 32)

	v.departureTime.Set([]byte(fields[4]))
	v.arrivalTime.Set([]byte(fields[5]))

	return v
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
