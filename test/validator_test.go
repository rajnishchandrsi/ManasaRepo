package validatortest

import (
	fmt "fmt"
	"testing"
)

func buildProto3(AlphaTrue, AlphaFalse, Beta, Noval string, Name []string, Country []string) *ValidatorMessage3 {

	goodProto3 := &ValidatorMessage3{
		AlphaTrue:  AlphaTrue,
		AlphaFalse: AlphaFalse,
		Beta:       Beta,
		Noval:      Noval,
		Name:       Name,
		Country:    Country,
	}

	return goodProto3
}

func TestStringRegex(t *testing.T) {
	var isTestCasePass = true
	var theArray []string
	theArray = append(theArray, "india")
	theArray = append(theArray, "China")
	theArray = append(theArray, "China")
	theArray = append(theArray, "China")
	theArray = append(theArray, "China")

	var theArray1 []string
	theArray1 = append(theArray1, "usa")

	tooLong1Proto3 := buildProto3("hello", "test", "a", "test", theArray, theArray1)
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

func buildProto3NestedInner(Inn string, Name string) *InnerMessage {
	goodProto3 := &InnerMessage{
		Inn: Inn,
		Name : Name,
	}
	return goodProto3
}

func buildProto3NestedOuter(Name string, Address *InnerMessage) *OuterMessage {
	goodProto3 := &OuterMessage{
		Name:    Name,
		Address: Address,
	}
	return goodProto3
}

func TestNested(t *testing.T) {
	var isTestCasePass = true

	inner := buildProto3NestedInner("inner123", "name12")
	outer := buildProto3NestedOuter("name", inner)
	if len(outer.Secvalidator()) > 0 {
		for _, err := range outer.Secvalidator() {
			fmt.Println(err)
		}
		isTestCasePass = false
	}

	if !isTestCasePass {
		t.Fatalf("expected fail in validator, but it didn't happen")
	}

}
