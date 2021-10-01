package validatortest

import (
	fmt "fmt"
	"testing"
)

func buildProto3(AlphaTrue, AlphaFalse, Noval string) *ValidatorMessage3 {

	goodProto3 := &ValidatorMessage3{
		AlphaTrue:  AlphaTrue,
		AlphaFalse: AlphaFalse,
		Noval:      Noval,
	}

	return goodProto3
}

func TestStringRegex(t *testing.T) {
	var isTestCasePass = true
	tooLong1Proto3 := buildProto3("hello", "test123^^&*&(*", "test1234")
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
