package validatortest

import (
	fmt "fmt"
	"testing"
)

var (
	stableBytes = make([]byte, 12)
)

const (
	uuid4 = "fbe91ff5-fee7-40d3-89a8-f3db6cf210be"
	uuid1 = "66bb25e2-2e0d-11e9-b210-d663bd873d93"
)

func buildProto3(AlphaTrue, AlphaFalse, BetaTrue, BetaFalse string) *ValidatorMessage3 {

	goodProto3 := &ValidatorMessage3{
		AlphaTrue:  AlphaTrue,
		AlphaFalse: AlphaFalse,
		BetaTrue:   BetaTrue,
		BetaFalse:  BetaFalse,
	}

	return goodProto3
}

func TestStringRegex(t *testing.T) {
	var isTestCasePass = true
	tooLong1Proto3 := buildProto3("hello", "test123", "test$%&", " ")
	if len(tooLong1Proto3.Secvalidator()) > 0 {
		for _, err := range tooLong1Proto3.Secvalidator() {
			fmt.Println(err)
		}
		isTestCasePass = false
	}

	fmt.Println("on 2nd condition")
	tooLong2Proto3 := buildProto3("hello#$%", "test123 (&*", "test @#!$#", " +_+")
	if len(tooLong2Proto3.Secvalidator()) > 0 {
		for _, err := range tooLong2Proto3.Secvalidator() {
			fmt.Println(err)
		}
		isTestCasePass = false

	}

	if !isTestCasePass {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}

}
