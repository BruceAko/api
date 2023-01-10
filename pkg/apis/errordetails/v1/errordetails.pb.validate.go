// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/apis/errordetails/v1/errordetails.proto

package errordetails

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

// Validate checks the field values on SourceError with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SourceError) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SourceError with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SourceErrorMultiError, or
// nil if none found.
func (m *SourceError) ValidateAll() error {
	return m.validate(true)
}

func (m *SourceError) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Temporary

	if all {
		switch v := interface{}(m.GetMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, SourceErrorValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, SourceErrorValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return SourceErrorValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return SourceErrorMultiError(errors)
	}

	return nil
}

// SourceErrorMultiError is an error wrapping multiple validation errors
// returned by SourceError.ValidateAll() if the designated constraints aren't met.
type SourceErrorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SourceErrorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SourceErrorMultiError) AllErrors() []error { return m }

// SourceErrorValidationError is the validation error returned by
// SourceError.Validate if the designated constraints aren't met.
type SourceErrorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SourceErrorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SourceErrorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SourceErrorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SourceErrorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SourceErrorValidationError) ErrorName() string { return "SourceErrorValidationError" }

// Error satisfies the builtin error interface
func (e SourceErrorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSourceError.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SourceErrorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SourceErrorValidationError{}
