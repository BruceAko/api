// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/apis/dfdaemon/v1/dfdaemon.proto

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

	common "d7y.io/api/v2/pkg/apis/common/v1"
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

	_ = common.TaskType(0)
)

// Validate checks the field values on DownRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DownRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DownRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DownRequestMultiError, or
// nil if none found.
func (m *DownRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DownRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uuid

	if uri, err := url.Parse(m.GetUrl()); err != nil {
		err = DownRequestValidationError{
			field:  "Url",
			reason: "value must be a valid URI",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	} else if !uri.IsAbs() {
		err := DownRequestValidationError{
			field:  "Url",
			reason: "value must be absolute",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetOutput()) < 1 {
		err := DownRequestValidationError{
			field:  "Output",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetTimeout() < 0 {
		err := DownRequestValidationError{
			field:  "Timeout",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetLimit() < 0 {
		err := DownRequestValidationError{
			field:  "Limit",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for DisableBackSource

	if m.GetUrlMeta() == nil {
		err := DownRequestValidationError{
			field:  "UrlMeta",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetUrlMeta()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DownRequestValidationError{
					field:  "UrlMeta",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DownRequestValidationError{
					field:  "UrlMeta",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUrlMeta()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DownRequestValidationError{
				field:  "UrlMeta",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Uid

	// no validation rules for Gid

	// no validation rules for KeepOriginalOffset

	// no validation rules for Recursive

	if len(errors) > 0 {
		return DownRequestMultiError(errors)
	}

	return nil
}

// DownRequestMultiError is an error wrapping multiple validation errors
// returned by DownRequest.ValidateAll() if the designated constraints aren't met.
type DownRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DownRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DownRequestMultiError) AllErrors() []error { return m }

// DownRequestValidationError is the validation error returned by
// DownRequest.Validate if the designated constraints aren't met.
type DownRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownRequestValidationError) ErrorName() string { return "DownRequestValidationError" }

// Error satisfies the builtin error interface
func (e DownRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownRequestValidationError{}

// Validate checks the field values on DownResult with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DownResult) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DownResult with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DownResultMultiError, or
// nil if none found.
func (m *DownResult) ValidateAll() error {
	return m.validate(true)
}

func (m *DownResult) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetTaskId()) < 1 {
		err := DownResultValidationError{
			field:  "TaskId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetPeerId()) < 1 {
		err := DownResultValidationError{
			field:  "PeerId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetCompletedLength() < 0 {
		err := DownResultValidationError{
			field:  "CompletedLength",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Done

	// no validation rules for Output

	if len(errors) > 0 {
		return DownResultMultiError(errors)
	}

	return nil
}

// DownResultMultiError is an error wrapping multiple validation errors
// returned by DownResult.ValidateAll() if the designated constraints aren't met.
type DownResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DownResultMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DownResultMultiError) AllErrors() []error { return m }

// DownResultValidationError is the validation error returned by
// DownResult.Validate if the designated constraints aren't met.
type DownResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DownResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DownResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DownResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DownResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DownResultValidationError) ErrorName() string { return "DownResultValidationError" }

// Error satisfies the builtin error interface
func (e DownResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDownResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DownResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DownResultValidationError{}

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

	if utf8.RuneCountInString(m.GetUrl()) < 1 {
		err := StatTaskRequestValidationError{
			field:  "Url",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetUrlMeta() == nil {
		err := StatTaskRequestValidationError{
			field:  "UrlMeta",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetUrlMeta()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StatTaskRequestValidationError{
					field:  "UrlMeta",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StatTaskRequestValidationError{
					field:  "UrlMeta",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUrlMeta()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StatTaskRequestValidationError{
				field:  "UrlMeta",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for LocalOnly

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

// Validate checks the field values on ImportTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ImportTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ImportTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ImportTaskRequestMultiError, or nil if none found.
func (m *ImportTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ImportTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetUrl()) < 1 {
		err := ImportTaskRequestValidationError{
			field:  "Url",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetUrlMeta() == nil {
		err := ImportTaskRequestValidationError{
			field:  "UrlMeta",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetUrlMeta()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ImportTaskRequestValidationError{
					field:  "UrlMeta",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ImportTaskRequestValidationError{
					field:  "UrlMeta",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUrlMeta()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ImportTaskRequestValidationError{
				field:  "UrlMeta",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetPath()) < 1 {
		err := ImportTaskRequestValidationError{
			field:  "Path",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Type

	if len(errors) > 0 {
		return ImportTaskRequestMultiError(errors)
	}

	return nil
}

// ImportTaskRequestMultiError is an error wrapping multiple validation errors
// returned by ImportTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type ImportTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ImportTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ImportTaskRequestMultiError) AllErrors() []error { return m }

// ImportTaskRequestValidationError is the validation error returned by
// ImportTaskRequest.Validate if the designated constraints aren't met.
type ImportTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImportTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImportTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImportTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImportTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImportTaskRequestValidationError) ErrorName() string {
	return "ImportTaskRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ImportTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImportTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImportTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImportTaskRequestValidationError{}

// Validate checks the field values on ExportTaskRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ExportTaskRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExportTaskRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ExportTaskRequestMultiError, or nil if none found.
func (m *ExportTaskRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ExportTaskRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetUrl()) < 1 {
		err := ExportTaskRequestValidationError{
			field:  "Url",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetOutput()) < 1 {
		err := ExportTaskRequestValidationError{
			field:  "Output",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetTimeout() < 0 {
		err := ExportTaskRequestValidationError{
			field:  "Timeout",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetLimit() < 0 {
		err := ExportTaskRequestValidationError{
			field:  "Limit",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetUrlMeta() == nil {
		err := ExportTaskRequestValidationError{
			field:  "UrlMeta",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetUrlMeta()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ExportTaskRequestValidationError{
					field:  "UrlMeta",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ExportTaskRequestValidationError{
					field:  "UrlMeta",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUrlMeta()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ExportTaskRequestValidationError{
				field:  "UrlMeta",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Uid

	// no validation rules for Gid

	// no validation rules for LocalOnly

	if len(errors) > 0 {
		return ExportTaskRequestMultiError(errors)
	}

	return nil
}

// ExportTaskRequestMultiError is an error wrapping multiple validation errors
// returned by ExportTaskRequest.ValidateAll() if the designated constraints
// aren't met.
type ExportTaskRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExportTaskRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExportTaskRequestMultiError) AllErrors() []error { return m }

// ExportTaskRequestValidationError is the validation error returned by
// ExportTaskRequest.Validate if the designated constraints aren't met.
type ExportTaskRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExportTaskRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExportTaskRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExportTaskRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExportTaskRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExportTaskRequestValidationError) ErrorName() string {
	return "ExportTaskRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ExportTaskRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExportTaskRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExportTaskRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExportTaskRequestValidationError{}

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

	if utf8.RuneCountInString(m.GetUrl()) < 1 {
		err := DeleteTaskRequestValidationError{
			field:  "Url",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetUrlMeta() == nil {
		err := DeleteTaskRequestValidationError{
			field:  "UrlMeta",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetUrlMeta()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DeleteTaskRequestValidationError{
					field:  "UrlMeta",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DeleteTaskRequestValidationError{
					field:  "UrlMeta",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUrlMeta()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DeleteTaskRequestValidationError{
				field:  "UrlMeta",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
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

// Validate checks the field values on PeerMetadata with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PeerMetadata) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PeerMetadata with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PeerMetadataMultiError, or
// nil if none found.
func (m *PeerMetadata) ValidateAll() error {
	return m.validate(true)
}

func (m *PeerMetadata) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TaskId

	// no validation rules for PeerId

	// no validation rules for State

	if len(errors) > 0 {
		return PeerMetadataMultiError(errors)
	}

	return nil
}

// PeerMetadataMultiError is an error wrapping multiple validation errors
// returned by PeerMetadata.ValidateAll() if the designated constraints aren't met.
type PeerMetadataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PeerMetadataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PeerMetadataMultiError) AllErrors() []error { return m }

// PeerMetadataValidationError is the validation error returned by
// PeerMetadata.Validate if the designated constraints aren't met.
type PeerMetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PeerMetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PeerMetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PeerMetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PeerMetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PeerMetadataValidationError) ErrorName() string { return "PeerMetadataValidationError" }

// Error satisfies the builtin error interface
func (e PeerMetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPeerMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PeerMetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PeerMetadataValidationError{}

// Validate checks the field values on PeerExchangeData with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PeerExchangeData) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PeerExchangeData with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PeerExchangeDataMultiError, or nil if none found.
func (m *PeerExchangeData) ValidateAll() error {
	return m.validate(true)
}

func (m *PeerExchangeData) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetPeerMetadatas() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PeerExchangeDataValidationError{
						field:  fmt.Sprintf("PeerMetadatas[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PeerExchangeDataValidationError{
						field:  fmt.Sprintf("PeerMetadatas[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PeerExchangeDataValidationError{
					field:  fmt.Sprintf("PeerMetadatas[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PeerExchangeDataMultiError(errors)
	}

	return nil
}

// PeerExchangeDataMultiError is an error wrapping multiple validation errors
// returned by PeerExchangeData.ValidateAll() if the designated constraints
// aren't met.
type PeerExchangeDataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PeerExchangeDataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PeerExchangeDataMultiError) AllErrors() []error { return m }

// PeerExchangeDataValidationError is the validation error returned by
// PeerExchangeData.Validate if the designated constraints aren't met.
type PeerExchangeDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PeerExchangeDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PeerExchangeDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PeerExchangeDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PeerExchangeDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PeerExchangeDataValidationError) ErrorName() string { return "PeerExchangeDataValidationError" }

// Error satisfies the builtin error interface
func (e PeerExchangeDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPeerExchangeData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PeerExchangeDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PeerExchangeDataValidationError{}
