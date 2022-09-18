package main

import (
	"fmt"
)

type Interval struct {
	lowerLimit float64
	upperLimit float64
}

func main() {
	message := "BBAABA"

	alphabet := []byte{byte('A'), byte('B')}
	pdistribution := []float64{3.0 / 6.0, 3.0 / 6.0}

	fmt.Println(alphabet)
	fmt.Println(pdistribution)

	compressedDocument := 0.785 //0.764005 // 0.47424349188804626
	fmt.Println(compressedDocument)

	currentInterval := Interval{lowerLimit: 0.0, upperLimit: 1.0}
	encoding := ""

	for _, _ = range message {
		loopingUpper := currentInterval.upperLimit
		loopingLower := currentInterval.lowerLimit

		loopingLength := loopingUpper - loopingLower

		intervalsToTest := []Interval{}

		for i := 0; i < len(pdistribution); i++ {
			intervalsToTest = append(intervalsToTest, Interval{lowerLimit: loopingLower, upperLimit: loopingLower + pdistribution[i]*loopingLength})
			loopingLower += loopingLength * pdistribution[i]
		}

		foundInterval := Interval{lowerLimit: -1.0, upperLimit: -1.0}
		for i := 0; i < len(intervalsToTest); i++ {
			if compressedDocument > intervalsToTest[i].lowerLimit && compressedDocument < intervalsToTest[i].upperLimit {
				foundInterval = intervalsToTest[i]
				encoding += string(alphabet[i])
				break
			}
		}
		if foundInterval.lowerLimit == -1.0 {
			fmt.Println("Something has seriously gone wrong!")
			break
		}
		currentInterval = foundInterval
	}

	fmt.Println("\n\nDecoded: ", encoding, "\n")
}
