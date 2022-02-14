package validatortest

import (
	fmt "fmt"
	"testing"
)

func buildProto3(AlphaTrue, AlphaFalse, Beta, Noval string, Name []string) *ValidatorMessage3 {

	goodProto3 := &ValidatorMessage3{
		AlphaTrue:  AlphaTrue,
		AlphaFalse: AlphaFalse,
		Beta:       Beta,
		Noval:      Noval,
		Name:       Name,
	}

	return goodProto3
}

func TestStringRegex(t *testing.T) {
	var isTestCasePass = true
	var theArray []string
	theArray=append(theArray, "india")
	theArray=append(theArray, "China")


	tooLong1Proto3 := buildProto3("hello", "test", "a", "test", theArray)
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
