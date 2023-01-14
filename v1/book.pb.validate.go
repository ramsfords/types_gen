// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: book.proto

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

// Validate checks the field values on BookingResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *BookingResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BookingResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// BookingResponseMultiError, or nil if none found.
func (m *BookingResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *BookingResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetShipmentDetails()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BookingResponseValidationError{
					field:  "ShipmentDetails",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BookingResponseValidationError{
					field:  "ShipmentDetails",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetShipmentDetails()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BookingResponseValidationError{
				field:  "ShipmentDetails",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetCommodities() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BookingResponseValidationError{
						field:  fmt.Sprintf("Commodities[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BookingResponseValidationError{
						field:  fmt.Sprintf("Commodities[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BookingResponseValidationError{
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
				errors = append(errors, BookingResponseValidationError{
					field:  "Pickup",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BookingResponseValidationError{
					field:  "Pickup",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPickup()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BookingResponseValidationError{
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
				errors = append(errors, BookingResponseValidationError{
					field:  "Delivery",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BookingResponseValidationError{
					field:  "Delivery",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDelivery()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BookingResponseValidationError{
				field:  "Delivery",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetBid()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BookingResponseValidationError{
					field:  "Bid",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BookingResponseValidationError{
					field:  "Bid",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBid()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BookingResponseValidationError{
				field:  "Bid",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for BookingSuccess

	if all {
		switch v := interface{}(m.GetDispatchResponse()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BookingResponseValidationError{
					field:  "DispatchResponse",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BookingResponseValidationError{
					field:  "DispatchResponse",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDispatchResponse()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BookingResponseValidationError{
				field:  "DispatchResponse",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetBookingInfo()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BookingResponseValidationError{
					field:  "BookingInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BookingResponseValidationError{
					field:  "BookingInfo",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBookingInfo()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BookingResponseValidationError{
				field:  "BookingInfo",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for BusinessId

	// no validation rules for QuoteId

	// no validation rules for SvgData

	// no validation rules for BolUrl

	if len(errors) > 0 {
		return BookingResponseMultiError(errors)
	}

	return nil
}

// BookingResponseMultiError is an error wrapping multiple validation errors
// returned by BookingResponse.ValidateAll() if the designated constraints
// aren't met.
type BookingResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BookingResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BookingResponseMultiError) AllErrors() []error { return m }

// BookingResponseValidationError is the validation error returned by
// BookingResponse.Validate if the designated constraints aren't met.
type BookingResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BookingResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BookingResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BookingResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BookingResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BookingResponseValidationError) ErrorName() string { return "BookingResponseValidationError" }

// Error satisfies the builtin error interface
func (e BookingResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBookingResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BookingResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BookingResponseValidationError{}

// Validate checks the field values on BookingInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *BookingInfo) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BookingInfo with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BookingInfoMultiError, or
// nil if none found.
func (m *BookingInfo) ValidateAll() error {
	return m.validate(true)
}

func (m *BookingInfo) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for FirstShipperBolNumber

	// no validation rules for FreightTerm

	// no validation rules for SpecialInstructions

	// no validation rules for CarrierName

	// no validation rules for CarrierProNumber

	// no validation rules for CarrierLogoUrl

	// no validation rules for CarrierDispatchContactNumber

	// no validation rules for CarrierDispatchEmail

	// no validation rules for CarrierBolNumber

	// no validation rules for CarrierReference

	if len(errors) > 0 {
		return BookingInfoMultiError(errors)
	}

	return nil
}

// BookingInfoMultiError is an error wrapping multiple validation errors
// returned by BookingInfo.ValidateAll() if the designated constraints aren't met.
type BookingInfoMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BookingInfoMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BookingInfoMultiError) AllErrors() []error { return m }

// BookingInfoValidationError is the validation error returned by
// BookingInfo.Validate if the designated constraints aren't met.
type BookingInfoValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BookingInfoValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BookingInfoValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BookingInfoValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BookingInfoValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BookingInfoValidationError) ErrorName() string { return "BookingInfoValidationError" }

// Error satisfies the builtin error interface
func (e BookingInfoValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBookingInfo.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BookingInfoValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BookingInfoValidationError{}

// Validate checks the field values on Bookings with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Bookings) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Bookings with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BookingsMultiError, or nil
// if none found.
func (m *Bookings) ValidateAll() error {
	return m.validate(true)
}

func (m *Bookings) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetBookingResponses() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BookingsValidationError{
						field:  fmt.Sprintf("BookingResponses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BookingsValidationError{
						field:  fmt.Sprintf("BookingResponses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BookingsValidationError{
					field:  fmt.Sprintf("BookingResponses[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return BookingsMultiError(errors)
	}

	return nil
}

// BookingsMultiError is an error wrapping multiple validation errors returned
// by Bookings.ValidateAll() if the designated constraints aren't met.
type BookingsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BookingsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BookingsMultiError) AllErrors() []error { return m }

// BookingsValidationError is the validation error returned by
// Bookings.Validate if the designated constraints aren't met.
type BookingsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BookingsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BookingsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BookingsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BookingsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BookingsValidationError) ErrorName() string { return "BookingsValidationError" }

// Error satisfies the builtin error interface
func (e BookingsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBookings.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BookingsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BookingsValidationError{}

// Validate checks the field values on DispatchResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *DispatchResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DispatchResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DispatchResponseMultiError, or nil if none found.
func (m *DispatchResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DispatchResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ShipmentID

	// no validation rules for SecurityKey

	// no validation rules for PickupNumber

	// no validation rules for CarrierName

	// no validation rules for CarrierPhone

	// no validation rules for CarrierPRONumber

	// no validation rules for HandlingUnitTotal

	// no validation rules for IsShipmentEdit

	// no validation rules for IsShipmentManual

	// no validation rules for ServiceType

	// no validation rules for IsTrackingEmailSend

	// no validation rules for IsTrackingAPIEnabled

	// no validation rules for CustomerBOLNumber

	// no validation rules for ShipperEmail

	// no validation rules for ConsigneeEmail

	if all {
		switch v := interface{}(m.GetResult()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, DispatchResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, DispatchResponseValidationError{
					field:  "Result",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DispatchResponseValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return DispatchResponseMultiError(errors)
	}

	return nil
}

// DispatchResponseMultiError is an error wrapping multiple validation errors
// returned by DispatchResponse.ValidateAll() if the designated constraints
// aren't met.
type DispatchResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DispatchResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DispatchResponseMultiError) AllErrors() []error { return m }

// DispatchResponseValidationError is the validation error returned by
// DispatchResponse.Validate if the designated constraints aren't met.
type DispatchResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DispatchResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DispatchResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DispatchResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DispatchResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DispatchResponseValidationError) ErrorName() string { return "DispatchResponseValidationError" }

// Error satisfies the builtin error interface
func (e DispatchResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDispatchResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DispatchResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DispatchResponseValidationError{}

// Validate checks the field values on Result with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Result) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Result with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ResultMultiError, or nil if none found.
func (m *Result) ValidateAll() error {
	return m.validate(true)
}

func (m *Result) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for CapacityProviderBolUrl

	// no validation rules for ShipmentIdentifier

	// no validation rules for PickupNote

	// no validation rules for PickupDateTime

	// no validation rules for Type

	if len(errors) > 0 {
		return ResultMultiError(errors)
	}

	return nil
}

// ResultMultiError is an error wrapping multiple validation errors returned by
// Result.ValidateAll() if the designated constraints aren't met.
type ResultMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ResultMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ResultMultiError) AllErrors() []error { return m }

// ResultValidationError is the validation error returned by Result.Validate if
// the designated constraints aren't met.
type ResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResultValidationError) ErrorName() string { return "ResultValidationError" }

// Error satisfies the builtin error interface
func (e ResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResultValidationError{}
