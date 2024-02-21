package c4

//https://cryptopals.com/sets/1/challenges/4

//Detect single-character XOR
//One of the 60-character strings in this file has been encrypted by single-character XOR.

//Find it.

//(Your code from #3 should help.)

import (
	"Crypto_Pals_Questions/c3"
	"os"
	"strings"
)

func divideStrings(sorcePath string) []string {
	sampleFile, err := os.ReadFile(sorcePath)
	if err != nil {
		panic(err)
	}
	stringsList := strings.Split(string(sampleFile), "\n")
	return stringsList
}

// Yes, the following two functions would be faster as one, so it
// only iterates once. However I decided I would make the code more
// modular so it can be reused later. It was a design choice, not
// stupidity

func turnIntoProbables(stringList []string, realTextFrequency map[rune]float64) []string {
	bestVersions := make([]string, len(stringList))
	for _, v := range stringList {
		meaning, _, err := c3.DecodeSingleXOR(v, realTextFrequency)
		if err != nil {
			panic(err)
		}
		bestVersions = append(bestVersions, meaning)
	}
	return bestVersions
}

func findMessage(textPath string, realTextPath string) string {
	realTextFrequency := c3.CharFrequencyMap("", realTextPath)
	strings := divideStrings(textPath)
	strings = turnIntoProbables(strings, realTextFrequency)
	topGrade := float64(0)
	topIndex := 0
	for index, string := range strings {
		grade := c3.EnglishScore(string, realTextFrequency)
		if grade > topGrade {
			topGrade = grade
			topIndex = index
		}
	}
	return strings[topIndex]
}
