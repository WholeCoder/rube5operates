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

	originalTextBytes, err := ReadInBytesFromFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(originalTextBytes)

	probDistributionAndMagicNumber := InitNewByteset([]byte{})
	fmt.Println(probDistributionAndMagicNumber)

	hashForDecompression := initFrequencyHashWithFloat64ForValues(os.Args[1])

	hashForDecompression.Sort(func(a *orderedmap.Pair, b *orderedmap.Pair) bool {
		return a.Value().(float64) < b.Value().(float64)
	})

	probabilityKeys := hashForDecompression.Keys()
	probabilityValues := hashForDecompression.Values()
	//for _, key := range probabilityKeys {
	//	value, _ := probabilityKeys.Get(key)
	//	probabilityValues = append(probabilityValues, value)
	//

	currentInterval := Interval{lowerLimit: 0.0, upperLimit: 1.0}
	fmt.Println(currentInterval)

	for idx, letter := range originalTextBytes {
		// break up currentInterval into sub intervals
		loopingLower := currentInterval.lowerLimit
		loopingUpper := currentInterval.upperLimit

		loopingLength := loopingUpper - loopingLower

		intervalsToTest := []Interval{}

		for jdx := 0; jdx < len(probabilityValues)-1; jdk++ {
			intervalsToTest = append(intervalsToTest, Interval{lowerLimit: loopingLower + probabilityValues[jdx]*loopingLength, upperLimit: loopingLower + probabilityValues[jdx]*loopingLength + probabilityValues[jdx-1]*loopingLength})
			loopingLower += probabilityValues[jdx]
		}

		// determine which one most closely fits current letter's probability
		var bestFit Interval = nil
		for _, value := range intervalsToTest {
			if bestFit == nil {
				bestFit = value
			} else if bestFit.lowerLimit < value.lowerLimit && bestFit.upperLimit > value.upperLimit {
				bestFit = value
			}
		}

		currentInterval = bestFit
	}

	fmt.Printf("Your magic interval is:  %@V", currentInterval)
}
