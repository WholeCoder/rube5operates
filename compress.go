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
	probabilityValues := []float64{}
	for _, key := range probabilityKeys {
		value, _ := probabilityKeys.Get(key)
		probabilityValues = append(probabilityValues, value)
	}

	currentInterval := Interval{lowerLimit: 0.0, upperLimit: 1.0}
	fmt.Println(currentInterval)

	for idx, letter := range originalTextBytes {
		// break up currentInterval into sub intervals
		loopingLower := currentInterval.lowerLimit
		loopingUpper := currentInterval.upperLimit

		intervalsToTest := []Interval{}

		for jdx := 0; jdx < len(probabilityValues)-1; jdk++ {
			intervalsToTest = append(intervalsToTest, Interval{lowerLimit: probabilityValues[jdx], upperLimit: probabilityValues[jdx-1]})
		}

		// determine which one most closely fits current letter's probability
		var bestFit Interval = nil
		for _, value := range intervalsToTest {
			if bestFit == nil {
				bestFit = value
			} else if 
		}

		// "expand" the the found interval and update the currentInterval with the new infomation
	}

}
