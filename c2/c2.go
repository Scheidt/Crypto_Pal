package c2
//https://cryptopals.com/sets/1/challenges/2
		//"Fixed XOR
//Write a function that takes two equal-length buffers and produces their XOR combination.

//If your function works properly, then when you feed it the string:
//1c0111001f010100061a024b53535009181c

//... after hex decoding, and when XOR'd against:
//686974207468652062756c6c277320657965

//... should produce:
//746865206b696420646f6e277420706c6179"

import (
	"fmt"
)

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
