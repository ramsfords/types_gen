// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: quote.proto

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

// Validate checks the field values on BookRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BookRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BookRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BookRequestMultiError, or
// nil if none found.
func (m *BookRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *BookRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetQuoteRequest()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BookRequestValidationError{
					field:  "QuoteRequest",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BookRequestValidationError{
					field:  "QuoteRequest",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetQuoteRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BookRequestValidationError{
				field:  "QuoteRequest",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for BidId

	if len(errors) > 0 {
		return BookRequestMultiError(errors)
	}

	return nil
}

// BookRequestMultiError is an error wrapping multiple validation errors
// returned by BookRequest.ValidateAll() if the designated constraints aren't met.
type BookRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BookRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BookRequestMultiError) AllErrors() []error { return m }

// BookRequestValidationError is the validation error returned by
// BookRequest.Validate if the designated constraints aren't met.
type BookRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BookRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BookRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BookRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BookRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BookRequestValidationError) ErrorName() string { return "BookRequestValidationError" }

// Error satisfies the builtin error interface
func (e BookRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBookRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BookRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BookRequestValidationError{}

// Validate checks the field values on QuoteRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QuoteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuoteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QuoteRequestMultiError, or
// nil if none found.
func (m *QuoteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *QuoteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for QuoteId

	// no validation rules for RequesterId

	// no validation rules for Mode

	// no validation rules for LiablePartyId

	// no validation rules for PickupDate

	// no validation rules for DisplayDate

	// no validation rules for DeliveryDate

	// no validation rules for TotalItems

	// no validation rules for TotalWeight

	// no validation rules for ValidUntil

	// no validation rules for EditMode

	// no validation rules for BusinessId

	// no validation rules for Type

	for idx, item := range m.GetCommodities() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, QuoteRequestValidationError{
						field:  fmt.Sprintf("Commodities[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, QuoteRequestValidationError{
						field:  fmt.Sprintf("Commodities[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QuoteRequestValidationError{
					field:  fmt.Sprintf("Commodities[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetPickup()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, QuoteRequestValidationError{
					field:  "Pickup",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, QuoteRequestValidationError{
					field:  "Pickup",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPickup()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QuoteRequestValidationError{
				field:  "Pickup",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDelivery()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, QuoteRequestValidationError{
					field:  "Delivery",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, QuoteRequestValidationError{
					field:  "Delivery",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDelivery()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QuoteRequestValidationError{
				field:  "Delivery",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for SpecialInstruction

	// no validation rules for ShipperPickupReadyBy

	// no validation rules for ShipperInstructions

	// no validation rules for ReceiverInstructions

	if all {
		switch v := interface{}(m.GetLocationServices()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, QuoteRequestValidationError{
					field:  "LocationServices",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, QuoteRequestValidationError{
					field:  "LocationServices",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLocationServices()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QuoteRequestValidationError{
				field:  "LocationServices",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetBusiness()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, QuoteRequestValidationError{
					field:  "Business",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, QuoteRequestValidationError{
					field:  "Business",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBusiness()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QuoteRequestValidationError{
				field:  "Business",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for RequesterEmail

	if len(errors) > 0 {
		return QuoteRequestMultiError(errors)
	}

	return nil
}

// QuoteRequestMultiError is an error wrapping multiple validation errors
// returned by QuoteRequest.ValidateAll() if the designated constraints aren't met.
type QuoteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuoteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuoteRequestMultiError) AllErrors() []error { return m }

// QuoteRequestValidationError is the validation error returned by
// QuoteRequest.Validate if the designated constraints aren't met.
type QuoteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuoteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuoteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuoteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuoteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuoteRequestValidationError) ErrorName() string { return "QuoteRequestValidationError" }

// Error satisfies the builtin error interface
func (e QuoteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuoteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuoteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuoteRequestValidationError{}

// Validate checks the field values on QuoteResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QuoteResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuoteResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QuoteResponseMultiError, or
// nil if none found.
func (m *QuoteResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *QuoteResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetQuoteRequest()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, QuoteResponseValidationError{
					field:  "QuoteRequest",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, QuoteResponseValidationError{
					field:  "QuoteRequest",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetQuoteRequest()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return QuoteResponseValidationError{
				field:  "QuoteRequest",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetBids() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, QuoteResponseValidationError{
						field:  fmt.Sprintf("Bids[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, QuoteResponseValidationError{
						field:  fmt.Sprintf("Bids[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QuoteResponseValidationError{
					field:  fmt.Sprintf("Bids[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Success

	if len(errors) > 0 {
		return QuoteResponseMultiError(errors)
	}

	return nil
}

// QuoteResponseMultiError is an error wrapping multiple validation errors
// returned by QuoteResponse.ValidateAll() if the designated constraints
// aren't met.
type QuoteResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuoteResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuoteResponseMultiError) AllErrors() []error { return m }

// QuoteResponseValidationError is the validation error returned by
// QuoteResponse.Validate if the designated constraints aren't met.
type QuoteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuoteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuoteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuoteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuoteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuoteResponseValidationError) ErrorName() string { return "QuoteResponseValidationError" }

// Error satisfies the builtin error interface
func (e QuoteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuoteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuoteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuoteResponseValidationError{}

// Validate checks the field values on QuotesRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QuotesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuotesRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QuotesRequestMultiError, or
// nil if none found.
func (m *QuotesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *QuotesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetQuoteRequests() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, QuotesRequestValidationError{
						field:  fmt.Sprintf("QuoteRequests[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, QuotesRequestValidationError{
						field:  fmt.Sprintf("QuoteRequests[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QuotesRequestValidationError{
					field:  fmt.Sprintf("QuoteRequests[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return QuotesRequestMultiError(errors)
	}

	return nil
}

// QuotesRequestMultiError is an error wrapping multiple validation errors
// returned by QuotesRequest.ValidateAll() if the designated constraints
// aren't met.
type QuotesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuotesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuotesRequestMultiError) AllErrors() []error { return m }

// QuotesRequestValidationError is the validation error returned by
// QuotesRequest.Validate if the designated constraints aren't met.
type QuotesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuotesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuotesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuotesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuotesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuotesRequestValidationError) ErrorName() string { return "QuotesRequestValidationError" }

// Error satisfies the builtin error interface
func (e QuotesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuotesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuotesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuotesRequestValidationError{}

// Validate checks the field values on QuotesResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QuotesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QuotesResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QuotesResponseMultiError,
// or nil if none found.
func (m *QuotesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *QuotesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetQuoteResponses() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, QuotesResponseValidationError{
						field:  fmt.Sprintf("QuoteResponses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, QuotesResponseValidationError{
						field:  fmt.Sprintf("QuoteResponses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return QuotesResponseValidationError{
					field:  fmt.Sprintf("QuoteResponses[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return QuotesResponseMultiError(errors)
	}

	return nil
}

// QuotesResponseMultiError is an error wrapping multiple validation errors
// returned by QuotesResponse.ValidateAll() if the designated constraints
// aren't met.
type QuotesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QuotesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QuotesResponseMultiError) AllErrors() []error { return m }

// QuotesResponseValidationError is the validation error returned by
// QuotesResponse.Validate if the designated constraints aren't met.
type QuotesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QuotesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QuotesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QuotesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QuotesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QuotesResponseValidationError) ErrorName() string { return "QuotesResponseValidationError" }

// Error satisfies the builtin error interface
func (e QuotesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQuotesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QuotesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QuotesResponseValidationError{}
