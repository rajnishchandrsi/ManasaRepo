// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: validator_proto3.proto

package validatortest

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on ValidatorMessage3 with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ValidatorMessage3) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AlphaTrue

	// no validation rules for AlphaFalse

	// no validation rules for Beta

	// no validation rules for Noval

	return nil
}

// ValidatorMessage3ValidationError is the validation error returned by
// ValidatorMessage3.Validate if the designated constraints aren't met.
type ValidatorMessage3ValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ValidatorMessage3ValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ValidatorMessage3ValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ValidatorMessage3ValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ValidatorMessage3ValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ValidatorMessage3ValidationError) ErrorName() string {
	return "ValidatorMessage3ValidationError"
}

// Error satisfies the builtin error interface
func (e ValidatorMessage3ValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sValidatorMessage3.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ValidatorMessage3ValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ValidatorMessage3ValidationError{}

// Validate checks the field values on OuterMessage with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *OuterMessage) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return OuterMessageValidationError{
				field:  "Address",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// OuterMessageValidationError is the validation error returned by
// OuterMessage.Validate if the designated constraints aren't met.
type OuterMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OuterMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OuterMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OuterMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OuterMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OuterMessageValidationError) ErrorName() string { return "OuterMessageValidationError" }

// Error satisfies the builtin error interface
func (e OuterMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOuterMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OuterMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OuterMessageValidationError{}

// Validate checks the field values on InnerMessage with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *InnerMessage) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Inn

	return nil
}

// InnerMessageValidationError is the validation error returned by
// InnerMessage.Validate if the designated constraints aren't met.
type InnerMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InnerMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InnerMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InnerMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InnerMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InnerMessageValidationError) ErrorName() string { return "InnerMessageValidationError" }

// Error satisfies the builtin error interface
func (e InnerMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInnerMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InnerMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InnerMessageValidationError{}
