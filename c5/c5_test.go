package c5

/*
To run the code run:
cd c5
go test -v
*/

import (
	"encoding/hex"
	"testing"
)

func Test_c5(test *testing.T) {
	input := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`
	key := "ICE"
	expectedResult := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	result := encryptRepeatingXor([]byte(input), []byte(key))
	if hex.EncodeToString(result)  != expectedResult{
		test.Fatal("result wrong! Your response is: ", result, "should be: ", expectedResult)
	}
	test.Logf("it worked!")
}
