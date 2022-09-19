package main

import (
	"fmt"
	"os"

	"github.com/iancoleman/orderedmap"
)

type Interval struct {
	lowerLimit float64
	upperLimit float64
}

func main() {

	//originalTextBytes, err := ReadInBytesFromFile(os.Args[1])
	//if err != nil {
	//	panic(err)
	//}
	originalTextBytes := []byte("BBBAABAABA BBAABAABAB")
	fmt.Println(originalTextBytes)

	hashForDecompression := initFrequencyHashWithFloat64ForValues(os.Args[1])

	hashForDecompression.Sort(func(a *orderedmap.Pair, b *orderedmap.Pair) bool {
		return a.Value().(float64) < b.Value().(float64)
	})

	//probabilityKeys := hashForDecompression.Keys()
	//probabilityValues := []float64{}
	//for _, key := range probabilityKeys {
	//	value, _ := hashForDecompression.Get(key)
	//	probabilityValues = append(probabilityValues, value.(float64))
	//}

	probabilityKeys := []string{"A", "B", " "}
	probabilityValues := []float64{1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0}

	currentInterval := Interval{lowerLimit: 0.0, upperLimit: 1.0}
	fmt.Println(currentInterval)

	for _, letter := range originalTextBytes {
		// break up currentInterval into sub intervals
		loopingUpper := currentInterval.upperLimit
		loopingLower := currentInterval.lowerLimit

		loopingLength := loopingUpper - loopingLower

		intervalsToTest := []Interval{}

		for jdx := 0; jdx < len(probabilityValues); jdx++ {
			intervalsToTest = append(intervalsToTest, Interval{lowerLimit: loopingLower, upperLimit: loopingLower + probabilityValues[jdx]*loopingLength})
			loopingLower += loopingLength * probabilityValues[jdx]
		}

		// determine which one most closely fits current letter's probability
		indexOfProbability := -1
		fmt.Println("*** start ***")
		for jdx := 0; jdx < len(probabilityKeys); jdx++ {
			fmt.Println("probabilityKeys[jdx] ==", probabilityKeys[jdx], " == ", string(letter))
			if probabilityKeys[jdx] == string(letter) {
				indexOfProbability = jdx
				break
			}
		}
		fmt.Println("*** end ***")

		currentInterval = intervalsToTest[indexOfProbability]
		fmt.Println("currentInterval.lowerLimit ==", currentInterval.lowerLimit)
	}

	fmt.Printf("Your magic interval is:  %#V", currentInterval)
	encodedDocument := (currentInterval.upperLimit + currentInterval.lowerLimit) / 2.0
	fmt.Println("\nYour magic number is: ", encodedDocument)
}
