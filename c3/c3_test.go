package c3
/*
To run the code run:
cd c3
go test -v
*/

import (
	"testing"
)

func Test_c3(t *testing.T){
	sampleString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	realTextFrequency := CharFrequencyMap("", "sampleText.txt")
	text, key, err := DecodeSingleXOR(sampleString, realTextFrequency)
	if err != nil{
		t.Fatal(err)
	}
	t.Logf("the probable text is: %v, with the key %c", text, key)
}