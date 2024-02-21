package c3

//https://cryptopals.com/sets/1/challenges/3
//Single-byte XOR cipher
// method used:
// https://en.wikipedia.org/wiki/Letter_frequency

//The hex encoded string:
//1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736

//... has been XOR'd against a single character. Find the key, decrypt the message.
//You can do this by hand. But don't: write code to do it for you.
//How? Devise some method for "scoring" a piece of English plaintext. Character frequency is a good metric. Evaluate each output and choose the one with the best score.

import (
	"encoding/hex"
	"fmt"
	"os"
	"unicode/utf8"
)

func CharFrequencyMap(sample string, path string) map[rune]float64 { // create a map [rune] frequency of a letter appearing in the text
	// if no sample is provided for comparison, a baseline will be created from the file "sampleText.txt"
	if sample == "" {
		sampleFile, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}
		sample = string(sampleFile)
	}
	frequency := make(map[rune]float64)
	for _, char := range sample {
		frequency[char]++
	}
	total := utf8.RuneCountInString(sample)
	for char := range frequency {
		frequency[char] = frequency[char] / float64(total)
	}
	return frequency
}

func SingleXOR(value []byte, key byte) []byte {
	xored := make([]byte, len(value))
	for i := range value {
		xored[i] = value[i] ^ key
	}
	return xored
}

func EnglishScore(sampleText string, frequency map[rune]float64) float64 {
	var score float64
	for _, char := range sampleText {
		score = score + frequency[char]
	}
	total := utf8.RuneCountInString(sampleText)
	return score / float64(total)
}


func DecodeSingleXOR(sample string, realTextFrequency map[rune]float64) (string, rune, error) {
	if sample == "" {
		return sample, 0, fmt.Errorf("very funny. Provide an actual string, please")
	}
	hexSample, err := hex.DecodeString(sample)
	if err != nil {
		return sample, 0, err
	}
	topScore := float64(0)
	var probableKey rune
	for run := range realTextFrequency {
		possibleOne := SingleXOR(hexSample, byte(run))
		score := EnglishScore(string(possibleOne), realTextFrequency)
		if score > topScore {
			topScore = score
			probableKey = run
		}
	}
	probableText := string(SingleXOR(hexSample, byte(probableKey))) // I do this at the end so the MMU and not in the loop
	// doesn't need to change the value every loop
	return probableText, probableKey, nil
}
