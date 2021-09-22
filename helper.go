// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

package secvalidator

import "strings"

// Validator is a general interface that allows a message to be validated.
type SecValidator interface {
	SecValidate() error
}

func CallValidatorIfExists(candidate interface{}) error {
	if validator, ok := candidate.(SecValidator); ok {
		return validator.SecValidate()
	}
	return nil
}

type fieldError struct {
	fieldStack []string
	nestedErr  error
}

func (f *fieldError) Error() string {
	return "invalid field " + strings.Join(f.fieldStack, ".") + ": " + f.nestedErr.Error()
}

// FieldError wraps a given Validator error providing a message call stack.
func FieldError(fieldName string, err error) error {
	if fErr, ok := err.(*fieldError); ok {
		fErr.fieldStack = append([]string{fieldName}, fErr.fieldStack...)
		return err
	}
	return &fieldError{
		fieldStack: []string{fieldName},
		nestedErr:  err,
	}
}
