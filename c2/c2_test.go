package c2
/*
To run the code run:
cd c2
go test -v
*/

import (
	"encoding/hex"
	"testing"
)

func Test_c2(test *testing.T) {
	
	expectedResult := "746865206b696420646f6e277420706c6179"


	value1, err := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	if err != nil {
		test.Fatal(err)
	}
	value2, err := hex.DecodeString("686974207468652062756c6c277320657965")
	if err != nil {
		test.Fatal(err)
	}
	

	result := CompareXOR([]byte(value1), []byte(value2))

	resultStringfied := hex.EncodeToString(result)
	if resultStringfied != expectedResult {
		test.Fatal("wrong! Your response is: ", resultStringfied, "should be: ", expectedResult)
	}
	test.Logf("It worked!")
}
