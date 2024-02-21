package c4

import (
	"testing"
	"Crypto_Pals_Questions/c3"
)

func Test_c4(test *testing.T){
	realTextFrequency := c3.CharFrequencyMap("", "sampleText.txt")
	strings := divideStrings("data.txt")
	strings = turnIntoProbables(strings, realTextFrequency)
	topGrade := float64(0)
	topIndex := 0
	for index, string := range strings{
		grade := c3.EnglishScore(string, realTextFrequency)
		if grade > topGrade{
			topGrade = grade
			topIndex = index
		}
	}
	test.Logf("The best text has a index of %v, and it reads %v", topIndex, strings[topIndex])
}