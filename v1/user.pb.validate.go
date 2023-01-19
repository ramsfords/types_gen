// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: user.proto

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

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *User) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on User with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in UserMultiError, or nil if none found.
func (m *User) ValidateAll() error {
	return m.validate(true)
}

func (m *User) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for UserName

	// no validation rules for Password

	// no validation rules for ConfirmPassword

	// no validation rules for Origin

	// no validation rules for EmailVisibility

	// no validation rules for Type

	// no validation rules for Id

	// no validation rules for Created

	// no validation rules for Updated

	// no validation rules for Verified

	// no validation rules for Avatar

	// no validation rules for LastResetSentAt

	// no validation rules for LastVerificationSentAt

	// no validation rules for PasswordHash

	// no validation rules for TokenKey

	// no validation rules for Token

	// no validation rules for Email

	if len(errors) > 0 {
		return UserMultiError(errors)
	}

	return nil
}

// UserMultiError is an error wrapping multiple validation errors returned by
// User.ValidateAll() if the designated constraints aren't met.
type UserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UserMultiError) AllErrors() []error { return m }

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

// Validate checks the field values on LoginUser with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginUser) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginUser with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginUserMultiError, or nil
// if none found.
func (m *LoginUser) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginUser) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Email

	// no validation rules for Password

	if len(errors) > 0 {
		return LoginUserMultiError(errors)
	}

	return nil
}

// LoginUserMultiError is an error wrapping multiple validation errors returned
// by LoginUser.ValidateAll() if the designated constraints aren't met.
type LoginUserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginUserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginUserMultiError) AllErrors() []error { return m }

// LoginUserValidationError is the validation error returned by
// LoginUser.Validate if the designated constraints aren't met.
type LoginUserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginUserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginUserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginUserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginUserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginUserValidationError) ErrorName() string { return "LoginUserValidationError" }

// Error satisfies the builtin error interface
func (e LoginUserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginUserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginUserValidationError{}

// Validate checks the field values on PocketCollection with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *PocketCollection) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PocketCollection with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PocketCollectionMultiError, or nil if none found.
func (m *PocketCollection) ValidateAll() error {
	return m.validate(true)
}

func (m *PocketCollection) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Type

	if len(errors) > 0 {
		return PocketCollectionMultiError(errors)
	}

	return nil
}

// PocketCollectionMultiError is an error wrapping multiple validation errors
// returned by PocketCollection.ValidateAll() if the designated constraints
// aren't met.
type PocketCollectionMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PocketCollectionMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PocketCollectionMultiError) AllErrors() []error { return m }

// PocketCollectionValidationError is the validation error returned by
// PocketCollection.Validate if the designated constraints aren't met.
type PocketCollectionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PocketCollectionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PocketCollectionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PocketCollectionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PocketCollectionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PocketCollectionValidationError) ErrorName() string { return "PocketCollectionValidationError" }

// Error satisfies the builtin error interface
func (e PocketCollectionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPocketCollection.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PocketCollectionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PocketCollectionValidationError{}

// Validate checks the field values on FrontEndUser with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FrontEndUser) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FrontEndUser with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FrontEndUserMultiError, or
// nil if none found.
func (m *FrontEndUser) ValidateAll() error {
	return m.validate(true)
}

func (m *FrontEndUser) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Email

	// no validation rules for Role

	if len(errors) > 0 {
		return FrontEndUserMultiError(errors)
	}

	return nil
}

// FrontEndUserMultiError is an error wrapping multiple validation errors
// returned by FrontEndUser.ValidateAll() if the designated constraints aren't met.
type FrontEndUserMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FrontEndUserMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FrontEndUserMultiError) AllErrors() []error { return m }

// FrontEndUserValidationError is the validation error returned by
// FrontEndUser.Validate if the designated constraints aren't met.
type FrontEndUserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FrontEndUserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FrontEndUserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FrontEndUserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FrontEndUserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FrontEndUserValidationError) ErrorName() string { return "FrontEndUserValidationError" }

// Error satisfies the builtin error interface
func (e FrontEndUserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFrontEndUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FrontEndUserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FrontEndUserValidationError{}
