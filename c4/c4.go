package c4

//https://cryptopals.com/sets/1/challenges/4

//Detect single-character XOR
//One of the 60-character strings in this file has been encrypted by single-character XOR.

//Find it.

//(Your code from #3 should help.)

import (
	"Crypto_Pals_Questions/c3"
	"fmt"
	"os"
	"strings"
)

//WIP: OPTIMIZE THIS CODE (It iterates twice, it could be done once)

func divideStrings(sorcePath string) []string{
	sampleFile, err := os.ReadFile(sorcePath)
	if err != nil {
		panic(err)
	}
	stringsList := strings.Split(string(sampleFile), "\n")
	/* for index, v := range stringsList{
		fmt.Printf("%v: %v\n", index, v)
	} */
	return stringsList
}

func turnIntoProbables(stringList []string, realTextFrequency map[rune]float64) []string{
	//fmt.Printf("%v", realTextFrequency)
	bestVersions := make([]string, len(stringList))
	for _, v := range stringList{
		//fmt.Printf("%v \n", v)
		meaning, _, err :=  c3.DecodeSingleXOR(v, realTextFrequency)
		fmt.Println(meaning)
		if err != nil{
			panic(err)
		}
		bestVersions = append(bestVersions, meaning)
	}
	return bestVersions
}
	/* var probableTextIndex int
	var probableTextKey rune
	var probableTextGrade float64
	for index, sample := range stringsList{
		meaning, key, err :=
		if err != nil{
			err := fmt.Errorf("%v: error in line 39", err)
			panic(err)
		}
		grade := c3.EnglishScore(meaning, realTextFrequency)
		if grade > probableTextGrade {
			probableTextGrade = grade
			probableTextIndex = index
			probableTextKey = key
		}
	}
	bytedText, err := hex.DecodeString(stringsList[probableTextIndex])
	if err != nil{
		panic(err)
	}
	decodedText := c3.SingleXOR(bytedText, byte(probableTextKey))
	decryptedMessage := hex.EncodeToString(decodedText)
	fmt.Printf("The line %v probably means %v, with the key %c", probableTextIndex+1, decryptedMessage,probableTextKey)
} */
