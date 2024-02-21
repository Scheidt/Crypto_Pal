package c1

import (
	"testing"
)

func Test_challange_1(test *testing.T) {
	expectedResult := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	str := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	value, err := hexToB64Converter(str)
	if err != nil {
		test.Fatal(err)
	}
	if value != expectedResult {
		test.Fatal("wrong! Your response is: ", value, "should be: ", expectedResult)
	}
	test.Logf("It worked!")
}
