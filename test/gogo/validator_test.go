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
	//^[a-zA-Z[]-/_`'\" \n\r\t&]
	tooLong1Proto3 := buildProto3("hello'\"`\t\n\r 1&", "test123^^&*&(*", "test")
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
