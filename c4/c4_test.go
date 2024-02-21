package c4
/*
To run the code run:
cd c4
go test -v
*/

import (
	"testing"
)

func Test_c4(test *testing.T){
	test.Logf(findMessage("data.txt", "../c3/sampleText.txt"))
}