// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pkg/apis/common/v2/common.proto

package common

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

// Validate checks the field values on Peer with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Peer) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Peer with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PeerMultiError, or nil if none found.
func (m *Peer) ValidateAll() error {
	return m.validate(true)
}

func (m *Peer) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) < 1 {
		err := PeerValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetPieces()) > 0 {

		if len(m.GetPieces()) < 1 {
			err := PeerValidationError{
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
						errors = append(errors, PeerValidationError{
							field:  fmt.Sprintf("Pieces[%v]", idx),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, PeerValidationError{
							field:  fmt.Sprintf("Pieces[%v]", idx),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return PeerValidationError{
						field:  fmt.Sprintf("Pieces[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}

	}

	if m.GetTask() == nil {
		err := PeerValidationError{
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
				errors = append(errors, PeerValidationError{
					field:  "Task",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PeerValidationError{
					field:  "Task",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PeerValidationError{
				field:  "Task",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetHost() == nil {
		err := PeerValidationError{
			field:  "Host",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetHost()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PeerValidationError{
					field:  "Host",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PeerValidationError{
					field:  "Host",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHost()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PeerValidationError{
				field:  "Host",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if utf8.RuneCountInString(m.GetState()) < 1 {
		err := PeerValidationError{
			field:  "State",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetCreatedAt() == nil {
		err := PeerValidationError{
			field:  "CreatedAt",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetUpdatedAt() == nil {
		err := PeerValidationError{
			field:  "UpdatedAt",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return PeerMultiError(errors)
	}

	return nil
}

// PeerMultiError is an error wrapping multiple validation errors returned by
// Peer.ValidateAll() if the designated constraints aren't met.
type PeerMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PeerMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PeerMultiError) AllErrors() []error { return m }

// PeerValidationError is the validation error returned by Peer.Validate if the
// designated constraints aren't met.
type PeerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PeerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PeerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PeerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PeerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PeerValidationError) ErrorName() string { return "PeerValidationError" }

// Error satisfies the builtin error interface
func (e PeerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPeer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PeerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PeerValidationError{}

// Validate checks the field values on Task with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Task) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Task with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TaskMultiError, or nil if none found.
func (m *Task) ValidateAll() error {
	return m.validate(true)
}

func (m *Task) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) < 1 {
		err := TaskValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := _Task_Type_InLookup[m.GetType()]; !ok {
		err := TaskValidationError{
			field:  "Type",
			reason: "value must be in list [normal super strong weak]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for SizeScope

	if len(m.GetPieces()) > 0 {

		if len(m.GetPieces()) < 1 {
			err := TaskValidationError{
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
						errors = append(errors, TaskValidationError{
							field:  fmt.Sprintf("Pieces[%v]", idx),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, TaskValidationError{
							field:  fmt.Sprintf("Pieces[%v]", idx),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return TaskValidationError{
						field:  fmt.Sprintf("Pieces[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}

	}

	if utf8.RuneCountInString(m.GetState()) < 1 {
		err := TaskValidationError{
			field:  "State",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetMetadata() == nil {
		err := TaskValidationError{
			field:  "Metadata",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetMetadata()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TaskValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TaskValidationError{
					field:  "Metadata",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMetadata()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TaskValidationError{
				field:  "Metadata",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetContentLength() < 1 {
		err := TaskValidationError{
			field:  "ContentLength",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetPeerCount() < 0 {
		err := TaskValidationError{
			field:  "PeerCount",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for HasAvailablePeer

	if m.GetCreatedAt() == nil {
		err := TaskValidationError{
			field:  "CreatedAt",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetUpdatedAt() == nil {
		err := TaskValidationError{
			field:  "UpdatedAt",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return TaskMultiError(errors)
	}

	return nil
}

// TaskMultiError is an error wrapping multiple validation errors returned by
// Task.ValidateAll() if the designated constraints aren't met.
type TaskMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TaskMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TaskMultiError) AllErrors() []error { return m }

// TaskValidationError is the validation error returned by Task.Validate if the
// designated constraints aren't met.
type TaskValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TaskValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TaskValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TaskValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TaskValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TaskValidationError) ErrorName() string { return "TaskValidationError" }

// Error satisfies the builtin error interface
func (e TaskValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTask.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TaskValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TaskValidationError{}

var _Task_Type_InLookup = map[string]struct{}{
	"normal": {},
	"super":  {},
	"strong": {},
	"weak":   {},
}

// Validate checks the field values on Host with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Host) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Host with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in HostMultiError, or nil if none found.
func (m *Host) ValidateAll() error {
	return m.validate(true)
}

func (m *Host) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) < 1 {
		err := HostValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetIpv4() != "" {

		if ip := net.ParseIP(m.GetIpv4()); ip == nil || ip.To4() == nil {
			err := HostValidationError{
				field:  "Ipv4",
				reason: "value must be a valid IPv4 address",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetIpv6() != "" {

		if ip := net.ParseIP(m.GetIpv6()); ip == nil || ip.To4() != nil {
			err := HostValidationError{
				field:  "Ipv6",
				reason: "value must be a valid IPv6 address",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if err := m._validateHostname(m.GetHostname()); err != nil {
		err = HostValidationError{
			field:  "Hostname",
			reason: "value must be a valid hostname",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetPort(); val < 1024 || val >= 65535 {
		err := HostValidationError{
			field:  "Port",
			reason: "value must be inside range [1024, 65535)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if val := m.GetDownloadPort(); val < 1024 || val >= 65535 {
		err := HostValidationError{
			field:  "DownloadPort",
			reason: "value must be inside range [1024, 65535)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetSecurityDomain() != "" {

		if utf8.RuneCountInString(m.GetSecurityDomain()) < 1 {
			err := HostValidationError{
				field:  "SecurityDomain",
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(m.GetLocation()) > 0 {

		if len(m.GetLocation()) < 1 {
			err := HostValidationError{
				field:  "Location",
				reason: "value must contain at least 1 item(s)",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetIdc() != "" {

		if utf8.RuneCountInString(m.GetIdc()) < 1 {
			err := HostValidationError{
				field:  "Idc",
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(m.GetNetTopology()) > 0 {

		if len(m.GetNetTopology()) < 1 {
			err := HostValidationError{
				field:  "NetTopology",
				reason: "value must contain at least 1 item(s)",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if len(errors) > 0 {
		return HostMultiError(errors)
	}

	return nil
}

func (m *Host) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

// HostMultiError is an error wrapping multiple validation errors returned by
// Host.ValidateAll() if the designated constraints aren't met.
type HostMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HostMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HostMultiError) AllErrors() []error { return m }

// HostValidationError is the validation error returned by Host.Validate if the
// designated constraints aren't met.
type HostValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HostValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HostValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HostValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HostValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HostValidationError) ErrorName() string { return "HostValidationError" }

// Error satisfies the builtin error interface
func (e HostValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHost.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HostValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HostValidationError{}

// Validate checks the field values on Range with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Range) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Range with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in RangeMultiError, or nil if none found.
func (m *Range) ValidateAll() error {
	return m.validate(true)
}

func (m *Range) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Begin

	// no validation rules for End

	if len(errors) > 0 {
		return RangeMultiError(errors)
	}

	return nil
}

// RangeMultiError is an error wrapping multiple validation errors returned by
// Range.ValidateAll() if the designated constraints aren't met.
type RangeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RangeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RangeMultiError) AllErrors() []error { return m }

// RangeValidationError is the validation error returned by Range.Validate if
// the designated constraints aren't met.
type RangeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RangeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RangeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RangeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RangeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RangeValidationError) ErrorName() string { return "RangeValidationError" }

// Error satisfies the builtin error interface
func (e RangeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRange.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RangeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RangeValidationError{}

// Validate checks the field values on Metadata with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Metadata) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Metadata with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MetadataMultiError, or nil
// if none found.
func (m *Metadata) ValidateAll() error {
	return m.validate(true)
}

func (m *Metadata) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if uri, err := url.Parse(m.GetUrl()); err != nil {
		err = MetadataValidationError{
			field:  "Url",
			reason: "value must be a valid URI",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	} else if !uri.IsAbs() {
		err := MetadataValidationError{
			field:  "Url",
			reason: "value must be absolute",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetDigest() != "" {

		if !_Metadata_Digest_Pattern.MatchString(m.GetDigest()) {
			err := MetadataValidationError{
				field:  "Digest",
				reason: "value does not match regex pattern \"^(md5)|(sha256):[A-Fa-f0-9]+$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if all {
		switch v := interface{}(m.GetRange()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MetadataValidationError{
					field:  "Range",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MetadataValidationError{
					field:  "Range",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRange()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MetadataValidationError{
				field:  "Range",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if _, ok := TaskType_name[int32(m.GetType())]; !ok {
		err := MetadataValidationError{
			field:  "Type",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Tag

	// no validation rules for Application

	if _, ok := Priority_name[int32(m.GetPriority())]; !ok {
		err := MetadataValidationError{
			field:  "Priority",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Header

	if m.GetPieceSize() < 1 {
		err := MetadataValidationError{
			field:  "PieceSize",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return MetadataMultiError(errors)
	}

	return nil
}

// MetadataMultiError is an error wrapping multiple validation errors returned
// by Metadata.ValidateAll() if the designated constraints aren't met.
type MetadataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetadataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetadataMultiError) AllErrors() []error { return m }

// MetadataValidationError is the validation error returned by
// Metadata.Validate if the designated constraints aren't met.
type MetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetadataValidationError) ErrorName() string { return "MetadataValidationError" }

// Error satisfies the builtin error interface
func (e MetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetadataValidationError{}

var _Metadata_Digest_Pattern = regexp.MustCompile("^(md5)|(sha256):[A-Fa-f0-9]+$")

// Validate checks the field values on Piece with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Piece) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Piece with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PieceMultiError, or nil if none found.
func (m *Piece) ValidateAll() error {
	return m.validate(true)
}

func (m *Piece) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetNumber() < 0 {
		err := PieceValidationError{
			field:  "Number",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetParentId() != "" {

		if utf8.RuneCountInString(m.GetParentId()) < 1 {
			err := PieceValidationError{
				field:  "ParentId",
				reason: "value length must be at least 1 runes",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetOffset() < 0 {
		err := PieceValidationError{
			field:  "Offset",
			reason: "value must be greater than or equal to 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetSize() <= 0 {
		err := PieceValidationError{
			field:  "Size",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetDigest() != "" {

		if !_Piece_Digest_Pattern.MatchString(m.GetDigest()) {
			err := PieceValidationError{
				field:  "Digest",
				reason: "value does not match regex pattern \"^(md5)|(sha256):[A-Fa-f0-9]+$\"",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	// no validation rules for TrafficType

	if m.GetCost() == nil {
		err := PieceValidationError{
			field:  "Cost",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetCreatedAt() == nil {
		err := PieceValidationError{
			field:  "CreatedAt",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return PieceMultiError(errors)
	}

	return nil
}

// PieceMultiError is an error wrapping multiple validation errors returned by
// Piece.ValidateAll() if the designated constraints aren't met.
type PieceMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PieceMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PieceMultiError) AllErrors() []error { return m }

// PieceValidationError is the validation error returned by Piece.Validate if
// the designated constraints aren't met.
type PieceValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PieceValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PieceValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PieceValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PieceValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PieceValidationError) ErrorName() string { return "PieceValidationError" }

// Error satisfies the builtin error interface
func (e PieceValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPiece.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PieceValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PieceValidationError{}

var _Piece_Digest_Pattern = regexp.MustCompile("^(md5)|(sha256):[A-Fa-f0-9]+$")

// Validate checks the field values on ExtendAttribute with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ExtendAttribute) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ExtendAttribute with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ExtendAttributeMultiError, or nil if none found.
func (m *ExtendAttribute) ValidateAll() error {
	return m.validate(true)
}

func (m *ExtendAttribute) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Header

	if val := m.GetStatusCode(); val < 100 || val >= 599 {
		err := ExtendAttributeValidationError{
			field:  "StatusCode",
			reason: "value must be inside range [100, 599)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetStatus()) < 1 {
		err := ExtendAttributeValidationError{
			field:  "Status",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ExtendAttributeMultiError(errors)
	}

	return nil
}

// ExtendAttributeMultiError is an error wrapping multiple validation errors
// returned by ExtendAttribute.ValidateAll() if the designated constraints
// aren't met.
type ExtendAttributeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ExtendAttributeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ExtendAttributeMultiError) AllErrors() []error { return m }

// ExtendAttributeValidationError is the validation error returned by
// ExtendAttribute.Validate if the designated constraints aren't met.
type ExtendAttributeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExtendAttributeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExtendAttributeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExtendAttributeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExtendAttributeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExtendAttributeValidationError) ErrorName() string { return "ExtendAttributeValidationError" }

// Error satisfies the builtin error interface
func (e ExtendAttributeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExtendAttribute.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExtendAttributeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExtendAttributeValidationError{}
