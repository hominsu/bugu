// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/detect.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on DetectRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DetectRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DetectRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DetectRequestMultiError, or
// nil if none found.
func (m *DetectRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DetectRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Size

	if len(errors) > 0 {
		return DetectRequestMultiError(errors)
	}

	return nil
}

// DetectRequestMultiError is an error wrapping multiple validation errors
// returned by DetectRequest.ValidateAll() if the designated constraints
// aren't met.
type DetectRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DetectRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DetectRequestMultiError) AllErrors() []error { return m }

// DetectRequestValidationError is the validation error returned by
// DetectRequest.Validate if the designated constraints aren't met.
type DetectRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DetectRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DetectRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DetectRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DetectRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DetectRequestValidationError) ErrorName() string { return "DetectRequestValidationError" }

// Error satisfies the builtin error interface
func (e DetectRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDetectRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DetectRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DetectRequestValidationError{}

// Validate checks the field values on DetectReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DetectReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DetectReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DetectReplyMultiError, or
// nil if none found.
func (m *DetectReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DetectReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	if len(errors) > 0 {
		return DetectReplyMultiError(errors)
	}

	return nil
}

// DetectReplyMultiError is an error wrapping multiple validation errors
// returned by DetectReply.ValidateAll() if the designated constraints aren't met.
type DetectReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DetectReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DetectReplyMultiError) AllErrors() []error { return m }

// DetectReplyValidationError is the validation error returned by
// DetectReply.Validate if the designated constraints aren't met.
type DetectReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DetectReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DetectReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DetectReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DetectReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DetectReplyValidationError) ErrorName() string { return "DetectReplyValidationError" }

// Error satisfies the builtin error interface
func (e DetectReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDetectReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DetectReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DetectReplyValidationError{}
