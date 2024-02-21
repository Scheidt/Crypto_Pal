package c4

import (
	"testing"
)

func Test_c4(test *testing.T){
	test.Logf(findMessage("data.txt", "sampleText.txt"))
}