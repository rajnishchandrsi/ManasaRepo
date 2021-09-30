package validatortest

import (
	fmt "fmt"
	"testing"
)



func buildProto3(AlphaTrue, AlphaFalse string) *ValidatorMessage3 {

	goodProto3 := &ValidatorMessage3{
		AlphaTrue:  AlphaTrue,
		AlphaFalse: AlphaFalse,
	}

	return goodProto3
}

func TestStringRegex(t *testing.T) {
	var isTestCasePass = true
	tooLong1Proto3 := buildProto3("hello", "test123")
	if len(tooLong1Proto3.Secvalidator()) > 0 {
		for _, err := range tooLong1Proto3.Secvalidator() {
			fmt.Println(err)
		}
		isTestCasePass = false
	}

	if !isTestCasePass {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}

}

func TestStringRegex2(t *testing.T) {
	var isTestCasePass = true
	tooLong2Proto3 := buildProto3("hello#$%", "test123 (&*")
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
