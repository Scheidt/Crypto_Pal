package c1
//https://cryptopals.com/sets/1/challenges/1

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
