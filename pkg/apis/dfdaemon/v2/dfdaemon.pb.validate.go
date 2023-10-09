// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/apis/dfdaemon/v2/dfdaemon.proto

package dfdaemon

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

// Validate checks the field values on InterestedAllPiecesRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *InterestedAllPiecesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InterestedAllPiecesRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// InterestedAllPiecesRequestMultiError, or nil if none found.
func (m *InterestedAllPiecesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *InterestedAllPiecesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return InterestedAllPiecesRequestMultiError(errors)
	}

	return nil
}

// InterestedAllPiecesRequestMultiError is an error wrapping multiple
// validation errors returned by InterestedAllPiecesRequest.ValidateAll() if
// the designated constraints aren't met.
type InterestedAllPiecesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InterestedAllPiecesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InterestedAllPiecesRequestMultiError) AllErrors() []error { return m }

// InterestedAllPiecesRequestValidationError is the validation error returned
// by InterestedAllPiecesRequest.Validate if the designated constraints aren't met.
type InterestedAllPiecesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InterestedAllPiecesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InterestedAllPiecesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InterestedAllPiecesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InterestedAllPiecesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InterestedAllPiecesRequestValidationError) ErrorName() string {
	return "InterestedAllPiecesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e InterestedAllPiecesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInterestedAllPiecesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InterestedAllPiecesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InterestedAllPiecesRequestValidationError{}

// Validate checks the field values on InterestedPiecesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *InterestedPiecesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InterestedPiecesRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// InterestedPiecesRequestMultiError, or nil if none found.
func (m *InterestedPiecesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *InterestedPiecesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetPieceNumbers()) < 1 {
		err := InterestedPiecesRequestValidationError{
			field:  "PieceNumbers",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return InterestedPiecesRequestMultiError(errors)
	}

	return nil
}

// InterestedPiecesRequestMultiError is an error wrapping multiple validation
// errors returned by InterestedPiecesRequest.ValidateAll() if the designated
// constraints aren't met.
type InterestedPiecesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InterestedPiecesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InterestedPiecesRequestMultiError) AllErrors() []error { return m }

// InterestedPiecesRequestValidationError is the validation error returned by
// InterestedPiecesRequest.Validate if the designated constraints aren't met.
type InterestedPiecesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InterestedPiecesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InterestedPiecesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InterestedPiecesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InterestedPiecesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InterestedPiecesRequestValidationError) ErrorName() string {
	return "InterestedPiecesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e InterestedPiecesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInterestedPiecesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InterestedPiecesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InterestedPiecesRequestValidationError{}

// Validate checks the field values on SyncPiecesRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SyncPiecesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SyncPiecesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SyncPiecesRequestMultiError, or nil if none found.
func (m *SyncPiecesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SyncPiecesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTaskId()) < 1 {
		err := SyncPiecesRequestValidationError{
			field:  "TaskId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	oneofRequestPresent := false
	switch v := m.Request.(type) {
	case *SyncPiecesRequest_InterestedAllPiecesRequest:
		if v == nil {
			err := SyncPiecesRequestValidationError{
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
			switch v := interface{}(m.GetInterestedAllPiecesRequest()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SyncPiecesRequestValidationError{
						field:  "InterestedAllPiecesRequest",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SyncPiecesRequestValidationError{
						field:  "InterestedAllPiecesRequest",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetInterestedAllPiecesRequest()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SyncPiecesRequestValidationError{
					field:  "InterestedAllPiecesRequest",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *SyncPiecesRequest_InterestedPiecesRequest:
		if v == nil {
			err := SyncPiecesRequestValidationError{
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
			switch v := interface{}(m.GetInterestedPiecesRequest()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SyncPiecesRequestValidationError{
						field:  "InterestedPiecesRequest",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SyncPiecesRequestValidationError{
						field:  "InterestedPiecesRequest",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetInterestedPiecesRequest()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SyncPiecesRequestValidationError{
					field:  "InterestedPiecesRequest",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofRequestPresent {
		err := SyncPiecesRequestValidationError{
			field:  "Request",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SyncPiecesRequestMultiError(errors)
	}

	return nil
}

// SyncPiecesRequestMultiError is an error wrapping multiple validation errors
// returned by SyncPiecesRequest.ValidateAll() if the designated constraints
// aren't met.
type SyncPiecesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SyncPiecesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SyncPiecesRequestMultiError) AllErrors() []error { return m }

// SyncPiecesRequestValidationError is the validation error returned by
// SyncPiecesRequest.Validate if the designated constraints aren't met.
type SyncPiecesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SyncPiecesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SyncPiecesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SyncPiecesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SyncPiecesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SyncPiecesRequestValidationError) ErrorName() string {
	return "SyncPiecesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e SyncPiecesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSyncPiecesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SyncPiecesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SyncPiecesRequestValidationError{}

// Validate checks the field values on InterestedPiecesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *InterestedPiecesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InterestedPiecesResponse with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// InterestedPiecesResponseMultiError, or nil if none found.
func (m *InterestedPiecesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *InterestedPiecesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetPieces()) > 0 {

		if len(m.GetPieces()) < 1 {
			err := InterestedPiecesResponseValidationError{
				field:  "Pieces",
				reason: "value must contain at least 1 item(s)",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		for idx, item := range m.GetPieces() {
			_, _ = idx, item

			if all {
				switch v := interface{}(item).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, InterestedPiecesResponseValidationError{
							field:  fmt.Sprintf("Pieces[%v]", idx),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, InterestedPiecesResponseValidationError{
							field:  fmt.Sprintf("Pieces[%v]", idx),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return InterestedPiecesResponseValidationError{
						field:  fmt.Sprintf("Pieces[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}

	}

	if len(errors) > 0 {
		return InterestedPiecesResponseMultiError(errors)
	}

	return nil
}

// InterestedPiecesResponseMultiError is an error wrapping multiple validation
// errors returned by InterestedPiecesResponse.ValidateAll() if the designated
// constraints aren't met.
type InterestedPiecesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InterestedPiecesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InterestedPiecesResponseMultiError) AllErrors() []error { return m }

// InterestedPiecesResponseValidationError is the validation error returned by
// InterestedPiecesResponse.Validate if the designated constraints aren't met.
type InterestedPiecesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InterestedPiecesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InterestedPiecesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InterestedPiecesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InterestedPiecesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InterestedPiecesResponseValidationError) ErrorName() string {
	return "InterestedPiecesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e InterestedPiecesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInterestedPiecesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InterestedPiecesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InterestedPiecesResponseValidationError{}

// Validate checks the field values on SyncPiecesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SyncPiecesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SyncPiecesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SyncPiecesResponseMultiError, or nil if none found.
func (m *SyncPiecesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SyncPiecesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	oneofResponsePresent := false
	switch v := m.Response.(type) {
	case *SyncPiecesResponse_InterestedPiecesResponse:
		if v == nil {
			err := SyncPiecesResponseValidationError{
				field:  "Response",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofResponsePresent = true

		if all {
			switch v := interface{}(m.GetInterestedPiecesResponse()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SyncPiecesResponseValidationError{
						field:  "InterestedPiecesResponse",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SyncPiecesResponseValidationError{
						field:  "InterestedPiecesResponse",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetInterestedPiecesResponse()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SyncPiecesResponseValidationError{
					field:  "InterestedPiecesResponse",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofResponsePresent {
		err := SyncPiecesResponseValidationError{
			field:  "Response",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SyncPiecesResponseMultiError(errors)
	}

	return nil
}

// SyncPiecesResponseMultiError is an error wrapping multiple validation errors
// returned by SyncPiecesResponse.ValidateAll() if the designated constraints
// aren't met.
type SyncPiecesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SyncPiecesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SyncPiecesResponseMultiError) AllErrors() []error { return m }

// SyncPiecesResponseValidationError is the validation error returned by
// SyncPiecesResponse.Validate if the designated constraints aren't met.
type SyncPiecesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SyncPiecesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SyncPiecesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SyncPiecesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SyncPiecesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SyncPiecesResponseValidationError) ErrorName() string {
	return "SyncPiecesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e SyncPiecesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSyncPiecesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SyncPiecesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SyncPiecesResponseValidationError{}

// Validate checks the field values on DownloadTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DownloadTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DownloadTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DownloadTaskRequestMultiError, or nil if none found.
func (m *DownloadTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DownloadTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetDownload() == nil {
		err := DownloadTaskRequestValidationError{
			field:  "Download",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetDownload()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DownloadTaskRequestValidationError{
					field:  "Download",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DownloadTaskRequestValidationError{
					field:  "Download",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDownload()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DownloadTaskRequestValidationError{
				field:  "Download",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DownloadTaskRequestMultiError(errors)
	}

	return nil
}

// DownloadTaskRequestMultiError is an error wrapping multiple validation
// errors returned by DownloadTaskRequest.ValidateAll() if the designated
// constraints aren't met.
type DownloadTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DownloadTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DownloadTaskRequestMultiError) AllErrors() []error { return m }

// DownloadTaskRequestValidationError is the validation error returned by
// DownloadTaskRequest.Validate if the designated constraints aren't met.
type DownloadTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownloadTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownloadTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownloadTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownloadTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownloadTaskRequestValidationError) ErrorName() string {
	return "DownloadTaskRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DownloadTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownloadTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownloadTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownloadTaskRequestValidationError{}

// Validate checks the field values on UploadTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *UploadTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UploadTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UploadTaskRequestMultiError, or nil if none found.
func (m *UploadTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UploadTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetTask() == nil {
		err := UploadTaskRequestValidationError{
			field:  "Task",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetTask()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UploadTaskRequestValidationError{
					field:  "Task",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UploadTaskRequestValidationError{
					field:  "Task",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UploadTaskRequestValidationError{
				field:  "Task",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UploadTaskRequestMultiError(errors)
	}

	return nil
}

// UploadTaskRequestMultiError is an error wrapping multiple validation errors
// returned by UploadTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type UploadTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UploadTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UploadTaskRequestMultiError) AllErrors() []error { return m }

// UploadTaskRequestValidationError is the validation error returned by
// UploadTaskRequest.Validate if the designated constraints aren't met.
type UploadTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UploadTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UploadTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UploadTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UploadTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UploadTaskRequestValidationError) ErrorName() string {
	return "UploadTaskRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UploadTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUploadTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UploadTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UploadTaskRequestValidationError{}

// Validate checks the field values on StatTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *StatTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StatTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// StatTaskRequestMultiError, or nil if none found.
func (m *StatTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StatTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTaskId()) < 1 {
		err := StatTaskRequestValidationError{
			field:  "TaskId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return StatTaskRequestMultiError(errors)
	}

	return nil
}

// StatTaskRequestMultiError is an error wrapping multiple validation errors
// returned by StatTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type StatTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StatTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StatTaskRequestMultiError) AllErrors() []error { return m }

// StatTaskRequestValidationError is the validation error returned by
// StatTaskRequest.Validate if the designated constraints aren't met.
type StatTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatTaskRequestValidationError) ErrorName() string { return "StatTaskRequestValidationError" }

// Error satisfies the builtin error interface
func (e StatTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStatTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatTaskRequestValidationError{}

// Validate checks the field values on DeleteTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DeleteTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteTaskRequestMultiError, or nil if none found.
func (m *DeleteTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTaskId()) < 1 {
		err := DeleteTaskRequestValidationError{
			field:  "TaskId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return DeleteTaskRequestMultiError(errors)
	}

	return nil
}

// DeleteTaskRequestMultiError is an error wrapping multiple validation errors
// returned by DeleteTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type DeleteTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteTaskRequestMultiError) AllErrors() []error { return m }

// DeleteTaskRequestValidationError is the validation error returned by
// DeleteTaskRequest.Validate if the designated constraints aren't met.
type DeleteTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteTaskRequestValidationError) ErrorName() string {
	return "DeleteTaskRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteTaskRequestValidationError{}
