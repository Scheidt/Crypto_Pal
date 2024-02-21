package c3

import (
	"testing"
)

func Test_c3(t *testing.T){
	sampleString := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	text, key, err := decodeSingleXOR(sampleString)
	if err != nil{
		t.Fatal(err)
	}
	t.Logf("the probable text is: %v, with the key %c", text, key)
}