package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

type Interval struct {
	lowerLimit float64
	upperLimit float64
}

func main() {

	readInBytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	probabilityKeysMarshalledLength := uint64(binary.BigEndian.Uint64(readInBytes[:8]))

	var s2 string = string(readInBytes[8 : probabilityKeysMarshalledLength+8])
	var probabilityKeys = []string{}
	err = json.Unmarshal([]byte(s2), &probabilityKeys)
	if err != nil {
		panic(err)
	}

	probabilityValuesMarshalledLength := uint64(binary.BigEndian.Uint64(readInBytes[probabilityKeysMarshalledLength+8 : probabilityKeysMarshalledLength+8+8]))

	var s3 string = string(readInBytes[probabilityKeysMarshalledLength+8+8 : probabilityKeysMarshalledLength+8+8+probabilityValuesMarshalledLength])
	var probabilityValues = []float64{}
	err = json.Unmarshal([]byte(s3), &probabilityValues)
	if err != nil {
		panic(err)
	}

	alphabet := probabilityKeys        // []byte{byte('A'), byte('B'), byte(' ')}
	pdistribution := probabilityValues // []float64{1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0}

	fmt.Println(alphabet)
	fmt.Println(pdistribution)

	compressedDocument := Float64frombytes(readInBytes[probabilityKeysMarshalledLength+8+8+probabilityValuesMarshalledLength : probabilityKeysMarshalledLength+8+8+probabilityValuesMarshalledLength+8]) // 0.48291785205218946 //0.8933463096618652 //0.78662109375 //0.764005 // 0.47424349188804626
	fmt.Println(compressedDocument)

	messageLength := uint64(binary.BigEndian.Uint64(readInBytes[probabilityKeysMarshalledLength+8+8+probabilityValuesMarshalledLength+8 : probabilityKeysMarshalledLength+8+8+probabilityValuesMarshalledLength+8+8]))

	currentInterval := Interval{lowerLimit: 0.0, upperLimit: 1.0}
	encoding := ""

	for count := 0; count < int(messageLength); count++ {
		loopingUpper := currentInterval.upperLimit
		loopingLower := currentInterval.lowerLimit

		fmt.Println(count, "\tinterval = [", loopingLower, ",", loopingUpper, "]")

		loopingLength := loopingUpper - loopingLower

		intervalsToTest := []Interval{}

		for i := 0; i < len(pdistribution); i++ {
			intervalsToTest = append(intervalsToTest, Interval{lowerLimit: loopingLower, upperLimit: loopingLower + pdistribution[i]*loopingLength})
			loopingLower += loopingLength * pdistribution[i]
		}

		foundInterval := Interval{lowerLimit: -1.0, upperLimit: -1.0}
		for i := 0; i < len(intervalsToTest); i++ {
			if compressedDocument >= intervalsToTest[i].lowerLimit && compressedDocument < intervalsToTest[i].upperLimit {
				foundInterval = intervalsToTest[i]
				encoding += string(alphabet[i])
				fmt.Println("\tdecompressed a: ", alphabet[i])
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

func Float64frombytes(bytes []byte) float64 {
	bits := binary.BigEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}
