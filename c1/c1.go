package c1
//https://cryptopals.com/sets/1/challenges/1
		//Convert hex to base64

//The string:
//49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d

//Should produce:
//SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t

//So go ahead and make that happen. You'll need to use this code for the rest of the exercises.

import (
	"encoding/base64"
	"encoding/hex"
)

func hexToB64Converter(hexa string) (string, error) {
	binar, err := hex.DecodeString(hexa)
	if err != nil{
		return "", err
	}
	b64 := base64.StdEncoding.EncodeToString(binar)
	return b64, nil
}
