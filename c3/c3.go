package c3
//https://cryptopals.com/sets/1/challenges/3

// method used:
	// https://en.wikipedia.org/wiki/Letter_frequency

import (
	"encoding/hex"
	"fmt"
	"os"
	"unicode/utf8"
)

func charFrequencyMap(sample string) map[rune]float64 { // create a map [rune] frequency of a letter appearing in the text
	// if no sample is provided for comparison, a baseline will be created from the file "sampleText.txt"
	if sample == "" {
		sampleFile, err := os.ReadFile("sampleText.txt")
		if err != nil {
			panic(err)
		}
		sample = string(sampleFile)
	}
	sampleText := string(sample)
	frequency := make(map[rune]float64)
	for _, char := range sampleText {
		frequency[char]++
	}
	total := utf8.RuneCountInString(sampleText)
	for char := range frequency {
		frequency[char] = frequency[char] / float64(total)
	}
	return frequency
}

func singleXOR(value []byte, key byte) []byte {
	xored := make([]byte, len(value))
	for i := range value {
		xored[i] = value[i] ^ key
	}
	return xored
}

func englishScore(sampleText string, frequency map[rune]float64) float64 {
	var score float64
	for _, char := range sampleText{
		score = score + frequency[char]
	}
	total := utf8.RuneCountInString(sampleText)
	return score / float64(total)
}

func decodeSingleXOR(sample string) (string, rune, error){
	if sample == ""{
		return sample, 0 , fmt.Errorf("very funny. Provide an actual string, please")
	}
	hexSample, err := hex.DecodeString(sample)
	if err != nil{
		return sample, 0, err
	}
	realTextFrequency := charFrequencyMap("")
	topScore := float64(0)
	var probableKey rune
	for run := range realTextFrequency{
		possibleOne := singleXOR(hexSample, byte(run))
		score := englishScore(string(possibleOne), realTextFrequency)
		if score > topScore{
			topScore = score
			probableKey = run
		}
	}
	probableText := string(singleXOR(hexSample, byte(probableKey))) // I do this at the end so the MMU and not in the loop
																	// doesn't need to change the value every loop
	return probableText, probableKey, nil
}
