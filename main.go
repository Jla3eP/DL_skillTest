package main

import (
	"DL_skillTest/sTime"
	"DL_skillTest/salesman"
	"DL_skillTest/train"
	"fmt"
)

const filePath = "test_task_data.csv"

func main() {
	matrix := train.VoyagesToMatrix(train.GetVoyages(filePath))

	bestPriceWay, _ := salesman.AnnealingAlgorithm(0.1, salesman.CountPrice, matrix)
	bestPrice, _ := salesman.CountPrice(bestPriceWay, matrix)
	fmt.Printf("Best price way: %v\nBest price: %.2f\n\n", bestPriceWay, bestPrice)

	bestTimeWay, _ := salesman.AnnealingAlgorithm(0.1, salesman.CountTime, matrix)
	bestTime, _ := salesman.CountTime(bestTimeWay, matrix)
	fmt.Printf("Best time way: %v\nBest time: %s\n\n", bestTimeWay, sTime.Time(bestTime).ToString())
}
