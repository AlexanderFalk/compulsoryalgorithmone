package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const INPUTVALUE = 300

//Random generating a number between 0 and INPUTVALUE and checking for distinction
func collect_coupons(n int) {
	var isCollected = make([]bool, n)
	var count int = 0
	var distinct int = 0

	rand.Seed(time.Now().UTC().UnixNano())
	for distinct < n {
		value := rand.Intn(n)
		count++

		//Every time a value doesn't match anything in the array, then we add it.
		//When the distinct is bigger than n, then we stop the program and print the count
		if !isCollected[value] {
			distinct++
			isCollected[value] = true
		}
	}
	calculateAccurary(count)
	fmt.Println("Total collected coupons: ", count)
}

//For accuracy, the closer to 1, the better, and the closer to prove the hypothesis
//Proving that accuracy between random numbers generated before all values has been found is ~n*Hn
func calculateAccurary(numberGenerationSize int) {

	numberSize := INPUTVALUE
	expectedResults := math.Ceil(getExpectedResultByHypothesis(numberSize))
	accuracy := float64(numberGenerationSize) / expectedResults
	fmt.Printf("%12s : %5d | %17s : %5d | %8s : %5.1f\n",
		"N", numberSize, "Numbers Generated Before All Values", numberGenerationSize,
		"Accuracy", accuracy)

}

func getExpectedResultByHypothesis(numberSize int) float64 {
	return float64(numberSize) * getHarmonicNumber(numberSize)
}

//Collecting and adding the harmonic series: 1 + 1/2 + 1/3 + 1/4 + 1/5 + 1/6 ...
//Returning the total sum
func getHarmonicNumber(number int) float64 {
	var harmonicValues = make([]float64, number+1)
	if len(harmonicValues) == 0 {
		fmt.Println("This is not good. Hold your horses. We are going down! %s")
	}
	index := 1
	for index <= number {
		harmonicValues[index] = (float64(1) / float64(index)) + harmonicValues[index-1]
		index++
	}
	return harmonicValues[number]
}

func main() {
	collect_coupons(INPUTVALUE)
}
