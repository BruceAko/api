// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/apis/trainer/v1/trainer.proto

package trainer

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

// Validate checks the field values on GNNRequest with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GNNRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GNNRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GNNRequestMultiError, or
// nil if none found.
func (m *GNNRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GNNRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetDataset()) < 1 {
		err := GNNRequestValidationError{
			field:  "Dataset",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GNNRequestMultiError(errors)
	}

	return nil
}

// GNNRequestMultiError is an error wrapping multiple validation errors
// returned by GNNRequest.ValidateAll() if the designated constraints aren't met.
type GNNRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GNNRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GNNRequestMultiError) AllErrors() []error { return m }

// GNNRequestValidationError is the validation error returned by
// GNNRequest.Validate if the designated constraints aren't met.
type GNNRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GNNRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GNNRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GNNRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GNNRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GNNRequestValidationError) ErrorName() string { return "GNNRequestValidationError" }

// Error satisfies the builtin error interface
func (e GNNRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGNNRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GNNRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GNNRequestValidationError{}

// Validate checks the field values on MLPRequest with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *MLPRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on MLPRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MLPRequestMultiError, or
// nil if none found.
func (m *MLPRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *MLPRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetDataset()) < 1 {
		err := MLPRequestValidationError{
			field:  "Dataset",
			reason: "value length must be at least 1 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return MLPRequestMultiError(errors)
	}

	return nil
}

// MLPRequestMultiError is an error wrapping multiple validation errors
// returned by MLPRequest.ValidateAll() if the designated constraints aren't met.
type MLPRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MLPRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MLPRequestMultiError) AllErrors() []error { return m }

// MLPRequestValidationError is the validation error returned by
// MLPRequest.Validate if the designated constraints aren't met.
type MLPRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MLPRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MLPRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MLPRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MLPRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MLPRequestValidationError) ErrorName() string { return "MLPRequestValidationError" }

// Error satisfies the builtin error interface
func (e MLPRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMLPRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MLPRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MLPRequestValidationError{}

// Validate checks the field values on TrainRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TrainRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TrainRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TrainRequestMultiError, or
// nil if none found.
func (m *TrainRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *TrainRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetHostname()) < 1 {
		err := TrainRequestValidationError{
			field:  "Hostname",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if ip := net.ParseIP(m.GetIp()); ip == nil {
		err := TrainRequestValidationError{
			field:  "Ip",
			reason: "value must be a valid IP address",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetClusterId() < 1 {
		err := TrainRequestValidationError{
			field:  "ClusterId",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	oneofRequestPresent := false
	switch v := m.Request.(type) {
	case *TrainRequest_GnnRequest:
		if v == nil {
			err := TrainRequestValidationError{
				field:  "Request",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofRequestPresent = true

		if all {
			switch v := interface{}(m.GetGnnRequest()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TrainRequestValidationError{
						field:  "GnnRequest",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TrainRequestValidationError{
						field:  "GnnRequest",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetGnnRequest()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TrainRequestValidationError{
					field:  "GnnRequest",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *TrainRequest_MlpRequest:
		if v == nil {
			err := TrainRequestValidationError{
				field:  "Request",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofRequestPresent = true

		if all {
			switch v := interface{}(m.GetMlpRequest()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, TrainRequestValidationError{
						field:  "MlpRequest",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, TrainRequestValidationError{
						field:  "MlpRequest",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetMlpRequest()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return TrainRequestValidationError{
					field:  "MlpRequest",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofRequestPresent {
		err := TrainRequestValidationError{
			field:  "Request",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return TrainRequestMultiError(errors)
	}

	return nil
}

// TrainRequestMultiError is an error wrapping multiple validation errors
// returned by TrainRequest.ValidateAll() if the designated constraints aren't met.
type TrainRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TrainRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TrainRequestMultiError) AllErrors() []error { return m }

// TrainRequestValidationError is the validation error returned by
// TrainRequest.Validate if the designated constraints aren't met.
type TrainRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TrainRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TrainRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TrainRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TrainRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TrainRequestValidationError) ErrorName() string { return "TrainRequestValidationError" }

// Error satisfies the builtin error interface
func (e TrainRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTrainRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TrainRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TrainRequestValidationError{}
