package train

import (
	"DL_skillTest/sTime"
	"fmt"
	"sort"
)

type Matrix map[int]map[int]Voyage

func RoadTime(v Voyage) sTime.Time {
	return sTime.Diff(v.departureTime, v.arrivalTime)
}

func VoyagesToMatrix(v Voyages) Matrix {
	res := Matrix{}

	for i := range v {
		priceMap, ok := res[v[i].departureStation]
		if ok {
			currentVoyage, ok := priceMap[v[i].arrivalStation]
			if ok {
				if currentVoyage.price > v[i].price || RoadTime(currentVoyage) > RoadTime(v[i]) {
					priceMap[v[i].arrivalStation] = v[i]
				}
			} else {
				priceMap[v[i].arrivalStation] = v[i]
			}
		} else {
			res[v[i].departureStation] = map[int]Voyage{v[i].arrivalStation: v[i]}
		}
	}
	return res
}

func (m Matrix) Print() {
	keys := make([]int, 0, len(m))
	for i := range m {
		keys = append(keys, i)
	}
	sort.Ints(keys)

	fmt.Printf("\n\t")
	for _, i := range keys {
		fmt.Printf("%v\t", i)
	}
	fmt.Printf("\n\n")

	for _, i := range keys {
		fmt.Printf("%v|\t", i)
		for _, j := range keys {
			res, ok := m[i][j]
			if ok {
				fmt.Printf("%.2f\t", res.price)
			} else {
				fmt.Printf("_____\t")
			}
		}
		fmt.Printf("\n")
	}
}
