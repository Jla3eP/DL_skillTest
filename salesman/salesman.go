package salesman

import (
	"DL_skillTest/sTime"
	"DL_skillTest/train"
	"errors"
	"math"
	"math/rand"
	"reflect"
	"sort"
)

const (
	startTemperature = 100
	maxFails         = 100
)

type temperature = float64
type weight = float64

type weightCounter = func([]int, train.Matrix) (weight, bool)

var (
	failCounter = 0
	currentWay  []int
)

var (
	CountTime = func(way []int, matrix train.Matrix) (weight, bool) {
		if !validWay(way, matrix) {
			return math.MaxFloat64, false
		}

		res := float64(0)
		for i := range way {
			if i+1 != len(way) {
				currVoyage := matrix[way[i]][way[i+1]]
				res += float64(train.RoadTime(currVoyage))
				if i != 0 {
					res += float64(sTime.Diff(matrix[way[i-1]][way[i]].GetArrTime(), currVoyage.GetDepTime()))
				}
			}
		}

		return res, true
	}
	CountPrice = func(way []int, matrix train.Matrix) (weight, bool) {
		if !validWay(way, matrix) {
			return math.MaxFloat64, false
		}

		res := float64(0)
		for i := range way {
			if i+1 != len(way) {
				res += matrix[way[i]][way[i+1]].GetPrice()
			}
		}
		return res, true
	}
)

func validWay(way []int, matrix train.Matrix) bool {
	for i := range way {
		if i+1 != len(way) { //is not last
			if _, ok := matrix[way[i]][way[i+1]]; !ok {
				return false
			}
		}
	}
	return true
}

func updateTemperature(currentTemperature *temperature, iteration int) {
	*currentTemperature = startTemperature * math.Log(float64(1+iteration)) / float64(1+iteration)
}

func swapRandElements(swapper func(i int, j int)) {
	i := 0
	j := 0
	for i == j {
		i = rand.Int() % len(currentWay)
		j = rand.Int() % len(currentWay)
	}
	swapper(i, j)
}

func findValidWay(matrix train.Matrix, swapper func(i int, j int)) error {
	for !validWay(currentWay, matrix) {
		if failCounter >= maxFails {
			return errors.New("fail counter exceeded\n")
		}
		failCounter++
		swapRandElements(swapper)
	}
	failCounter = 0
	return nil
}

func AnnealingAlgorithm(minTemperature temperature, weightFunc weightCounter, matrix train.Matrix) ([]int, error) {
	currentTemperature := float64(startTemperature)
	iteration := 0
	currentWeight := math.MaxFloat64
	currentWay = make([]int, len(matrix), len(matrix))

	counter := 0
	for i := range matrix {
		currentWay[counter] = i
		counter++
	}

	sort.Ints(currentWay)

	swapper := reflect.Swapper(currentWay)
	err := findValidWay(matrix, swapper)
	if err != nil {
		return nil, err
	}
	currentWeight, _ = weightFunc(currentWay, matrix)

	resetCopy := make([]int, len(matrix), len(matrix))
	copy(resetCopy, currentWay)

	for currentTemperature > minTemperature {
		if iteration != 0 {
			updateTemperature(&currentTemperature, iteration)
		}
		//fmt.Printf("t = %.2f\t w = %.2f\n", currentTemperature, currentWeight)
		copy(resetCopy, currentWay)
		swapRandElements(swapper)
		if !validWay(currentWay, matrix) {
			err := findValidWay(matrix, swapper)
			if err != nil {
				return nil, err
			}
		}
		if newWeight, ok := weightFunc(currentWay, matrix); ok {
			if newWeight < currentWeight {
				currentWeight = newWeight
			} else {
				if math.Pow(math.E, -(newWeight-currentWeight)/currentTemperature) < rand.Float64() {
					copy(currentWay, resetCopy)
				} else {
					currentWeight = newWeight
				}
				iteration++
			}
		}
	}

	return currentWay, nil
}
