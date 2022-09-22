package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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
	//originalTextBytes := []byte("BBBAABAABA BBAABAABAB")
	fmt.Println(originalTextBytes)

	hashForDecompression := initFrequencyHashWithFloat64ForValues(os.Args[1])

	// hashForDecompression.Sort(func(a *orderedmap.Pair, b *orderedmap.Pair) bool {
	//	return a.Value().(float64) < b.Value().(float64)
	// })
	hashForDecompression.SortKeys(sort.Strings)

	probabilityKeys := hashForDecompression.Keys()
	probabilityValues := []float64{}
	for _, key := range probabilityKeys {
		value, _ := hashForDecompression.Get(key)
		probabilityValues = append(probabilityValues, value.(float64))
	}

	// probabilityKeys := []string{"A", "B", " "}
	//probabilityValues := []float64{1.0 / 3.0, 1.0 / 3.0, 1.0 / 3.0}

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
	fmt.Printf("\nProbability keys are:  %#V\n", probabilityKeys)
	fmt.Printf("\nProbability values are:  %#V\n", probabilityValues)

	lengthOfMessage := len(originalTextBytes)
	fmt.Println("Length of message is:  ", lengthOfMessage)

	encodedDocument := (currentInterval.upperLimit + currentInterval.lowerLimit) / 2.0
	fmt.Println("\nYour magic number is: ", encodedDocument)

	fmt.Printf("\n len(hashForDecompression):  %#V\n", len(hashForDecompression.Keys()))

	// Write out
	//  the probabilityKey slice
	probabilityKeysMarshalled, err := json.Marshal(probabilityKeys)
	if err != nil {
		panic(err)
	}
	// probabilityKeysMarshalled length
	probabilityKeysMarshalledLength := getBytesForInt(len(probabilityKeysMarshalled))
	//  the probabilityValue slice
	probabilityValuesMarshalled, err := json.Marshal(probabilityValues)
	if err != nil {
		panic(err)
	}
	probabilityValuesMarshalledLength := getBytesForInt(len(probabilityValuesMarshalled))
	//  decimal enocoding of the text
	byteEncodedDecimalCodedDocument := float64ToByte(encodedDocument)

	outputSlice := []byte{}
	outputSlice = append(outputSlice, probabilityKeysMarshalledLength...)
	outputSlice = append(outputSlice, probabilityKeysMarshalled...)
	outputSlice = append(outputSlice, probabilityValuesMarshalledLength...)
	outputSlice = append(outputSlice, probabilityValuesMarshalled...)
	outputSlice = append(outputSlice, byteEncodedDecimalCodedDocument...)

	file, err := os.OpenFile(
		os.Args[1]+".comp",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}

	bytesWritten, err := file.Write(outputSlice)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)

}

func float64ToByte(f float64) []byte {
	var buf [8]byte
	binary.BigEndian.PutUint64(buf[:], math.Float64bits(f))
	return buf[:]
}

func getBytesForInt(length int) []byte {

	b := make([]byte, 8)

	binary.BigEndian.PutUint64(b, uint64(length))

	return b
}

func Float64frombytes(bytes []byte) float64 {
	bits := binary.BigEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}
