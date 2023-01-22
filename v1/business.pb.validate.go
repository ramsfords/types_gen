// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: business.proto

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

// Validate checks the field values on Business with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Business) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Business with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BusinessMultiError, or nil
// if none found.
func (m *Business) ValidateAll() error {
	return m.validate(true)
}

func (m *Business) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for BusinessName

	// no validation rules for BusinessId

	// no validation rules for BusinessEmail

	// no validation rules for AccountingEmail

	// no validation rules for CustomerServiceEmail

	// no validation rules for HighPriorityEmail

	// no validation rules for AvatarURL

	// no validation rules for AdminEmail

	// no validation rules for CreateAt

	// no validation rules for UpdatedUt

	// no validation rules for DeletedAt

	if all {
		switch v := interface{}(m.GetPhoneNumber()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BusinessValidationError{
					field:  "PhoneNumber",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BusinessValidationError{
					field:  "PhoneNumber",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPhoneNumber()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BusinessValidationError{
				field:  "PhoneNumber",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for NeedsAddressUpdate

	// no validation rules for NeedsDefaultPickupAddressUpdate

	// no validation rules for NeedsDefaultDeliveryAddressUpdate

	if all {
		switch v := interface{}(m.GetBillingAddress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BusinessValidationError{
					field:  "BillingAddress",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BusinessValidationError{
					field:  "BillingAddress",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBillingAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BusinessValidationError{
				field:  "BillingAddress",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDefaultPickupAddress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BusinessValidationError{
					field:  "DefaultPickupAddress",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BusinessValidationError{
					field:  "DefaultPickupAddress",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDefaultPickupAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BusinessValidationError{
				field:  "DefaultPickupAddress",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetDefaultDeliveryAddress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BusinessValidationError{
					field:  "DefaultDeliveryAddress",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BusinessValidationError{
					field:  "DefaultDeliveryAddress",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDefaultDeliveryAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BusinessValidationError{
				field:  "DefaultDeliveryAddress",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetAddress()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BusinessValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BusinessValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BusinessValidationError{
				field:  "Address",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return BusinessMultiError(errors)
	}

	return nil
}

// BusinessMultiError is an error wrapping multiple validation errors returned
// by Business.ValidateAll() if the designated constraints aren't met.
type BusinessMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BusinessMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BusinessMultiError) AllErrors() []error { return m }

// BusinessValidationError is the validation error returned by
// Business.Validate if the designated constraints aren't met.
type BusinessValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BusinessValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BusinessValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BusinessValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BusinessValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BusinessValidationError) ErrorName() string { return "BusinessValidationError" }

// Error satisfies the builtin error interface
func (e BusinessValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBusiness.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BusinessValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BusinessValidationError{}

// Validate checks the field values on Businesses with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Businesses) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Businesses with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in BusinessesMultiError, or
// nil if none found.
func (m *Businesses) ValidateAll() error {
	return m.validate(true)
}

func (m *Businesses) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetBusinesses() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, BusinessesValidationError{
						field:  fmt.Sprintf("Businesses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, BusinessesValidationError{
						field:  fmt.Sprintf("Businesses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return BusinessesValidationError{
					field:  fmt.Sprintf("Businesses[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return BusinessesMultiError(errors)
	}

	return nil
}

// BusinessesMultiError is an error wrapping multiple validation errors
// returned by Businesses.ValidateAll() if the designated constraints aren't met.
type BusinessesMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BusinessesMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BusinessesMultiError) AllErrors() []error { return m }

// BusinessesValidationError is the validation error returned by
// Businesses.Validate if the designated constraints aren't met.
type BusinessesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BusinessesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BusinessesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BusinessesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BusinessesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BusinessesValidationError) ErrorName() string { return "BusinessesValidationError" }

// Error satisfies the builtin error interface
func (e BusinessesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBusinesses.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BusinessesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BusinessesValidationError{}

// Validate checks the field values on HomePageResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *HomePageResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HomePageResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HomePageResponseMultiError, or nil if none found.
func (m *HomePageResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *HomePageResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetBusiness()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HomePageResponseValidationError{
					field:  "Business",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HomePageResponseValidationError{
					field:  "Business",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBusiness()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HomePageResponseValidationError{
				field:  "Business",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HomePageResponseValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HomePageResponseValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HomePageResponseValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetBookings() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HomePageResponseValidationError{
						field:  fmt.Sprintf("Bookings[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HomePageResponseValidationError{
						field:  fmt.Sprintf("Bookings[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HomePageResponseValidationError{
					field:  fmt.Sprintf("Bookings[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetQuotes() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HomePageResponseValidationError{
						field:  fmt.Sprintf("Quotes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HomePageResponseValidationError{
						field:  fmt.Sprintf("Quotes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HomePageResponseValidationError{
					field:  fmt.Sprintf("Quotes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return HomePageResponseMultiError(errors)
	}

	return nil
}

// HomePageResponseMultiError is an error wrapping multiple validation errors
// returned by HomePageResponse.ValidateAll() if the designated constraints
// aren't met.
type HomePageResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HomePageResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HomePageResponseMultiError) AllErrors() []error { return m }

// HomePageResponseValidationError is the validation error returned by
// HomePageResponse.Validate if the designated constraints aren't met.
type HomePageResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HomePageResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HomePageResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HomePageResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HomePageResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HomePageResponseValidationError) ErrorName() string { return "HomePageResponseValidationError" }

// Error satisfies the builtin error interface
func (e HomePageResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHomePageResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HomePageResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HomePageResponseValidationError{}
