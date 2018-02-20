package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const INPUT int = 200

//Generates a random number and take n as input. n is the number being randomized
//Appends the numbers to a list and a map
//Checking the map each loop; when it finds an entry in the map, then it knows
//that it is a duplicate
func experiment(n int) int {
	duplicate_frequency := make(map[int]int)
	rand.Seed(time.Now().UTC().UnixNano())
	var list []int
	for {
		value := rand.Intn(n - 1)
		_, exist := duplicate_frequency[value]
		list = append(list, value)
		// check if the item/element exist in the duplicate_frequency map
		if exist {
			fmt.Println("Duplicated value: ", value)
			hypothesis(list)
			break
		} else {

			duplicate_frequency[value] = 1 // else start counting from 1
		}

	}
	return len(duplicate_frequency)
}

//The hypothesis checks, if the accuracy of the total numbers generated before
//a duplicate appears is close to the formular math.Sqrt(math.Pi * 50) / 2
func hypothesis(numberGenerated []int) {
	result := (math.Pi * float64(INPUT)) / 2
	hypoConf := math.Ceil(math.Sqrt(result))
	accuracy := float64(len(numberGenerated)) / hypoConf
	fmt.Printf("%12s : %5d | %17s : %5d | %8s : %5.1f\n",
		"N", float64(INPUT), "Numbers Generated Before Repeat", len(numberGenerated),
		"Accuracy", accuracy)
}

//You can change the value of the experiment, if you want to try different numbers
func main() {
	experiment(INPUT)
}
