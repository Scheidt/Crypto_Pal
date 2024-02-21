package c2
//https://cryptopals.com/sets/1/challenges/2

import (
	"encoding/hex"
	"fmt"
)

func decodeHex(v string) ([]byte, error) {
	bi, err := hex.DecodeString(v)
	if err != nil {
		return nil, err
	}
	return bi, nil
}

func compareXOR(v1, v2 []byte) ([]byte, error) {
	if len(v1) != len(v2) {
		return nil, fmt.Errorf("the two buffers must be of equal lengths")
	}
	xored := make([]byte, len(v1))
	for i := range v1 {
		xored[i] = v1[i] ^ v2[i]
	}
	return xored, nil
}
